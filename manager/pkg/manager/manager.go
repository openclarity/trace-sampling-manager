// Copyright © 2021 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manager

import (
	"encoding/json"
	"fmt"
	"sync"

	_secret "github.com/Portshift/go-utils/k8s/secret"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/openclarity/trace-sampling-manager/manager/pkg/grpc"
	_interface "github.com/openclarity/trace-sampling-manager/manager/pkg/manager/interface"
	"github.com/openclarity/trace-sampling-manager/manager/pkg/rest"
	"github.com/openclarity/trace-sampling-manager/manager/pkg/utils"
)

const (
	dataFieldName = "hosts-per-component-map"
)

func createHostToTraceSecretMeta(secretName string, secretNamespace string, secretOwner string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      secretName,
		Namespace: secretNamespace,
		Labels:    map[string]string{"owner": secretOwner},
	}
}

type Config struct {
	RestServerPort             int
	GRPCServerPort             int
	HostToTraceSecretName      string
	HostToTraceSecretNamespace string
	HostToTraceSecretOwnerName string
}

type Manager struct {
	_secret.Handler
	restServer *rest.Server
	grpcServer *grpc.Server

	hostsToTrace          []string
	componentIDToHosts    map[string][]string
	hostToTraceSecretMeta metav1.ObjectMeta

	sync.RWMutex
}

func Create(clientset kubernetes.Interface, conf *Config) (*Manager, error) {
	var err error
	m := &Manager{
		Handler:               _secret.NewHandler(clientset),
		hostToTraceSecretMeta: createHostToTraceSecretMeta(conf.HostToTraceSecretName, conf.HostToTraceSecretNamespace, conf.HostToTraceSecretOwnerName),
	}

	m.initHostToTrace()

	m.restServer, err = rest.CreateRESTServer(conf.RestServerPort, m)
	if err != nil {
		return nil, fmt.Errorf("failed to start rest server: %v", err)
	}

	m.grpcServer = grpc.NewServer(conf.GRPCServerPort, m, m)

	return m, nil
}

func (m *Manager) Start(errChan chan struct{}) error {
	if err := m.grpcServer.Start(errChan); err != nil {
		return fmt.Errorf("failed to start GRPC server: %v", err)
	}
	m.restServer.Start(errChan)

	return nil
}

func (m *Manager) Stop() {
	if m.grpcServer != nil {
		m.grpcServer.Stop()
	}
	if m.restServer != nil {
		m.restServer.Stop()
	}
}

func (m *Manager) HostsToTrace() []string {
	m.RLock()
	defer m.RUnlock()

	return m.hostsToTrace
}

func (m *Manager) HostsToTraceByComponentID(id string) []string {
	m.RLock()
	defer m.RUnlock()

	return m.componentIDToHosts[id]
}

func (m *Manager) SetHostsToTrace(hostsToTrace *_interface.HostsByComponentID) {
	m.Lock()
	defer m.Unlock()

	// reset to the most up-to-date hosts list for the given component
	m.componentIDToHosts[hostsToTrace.ComponentID] = hostsToTrace.Hosts
	m.hostsToTrace = createHostsToTrace(m.componentIDToHosts)
	if err := m.saveComponentIDToHosts(); err != nil {
		// TODO: consider retrying
		log.Errorf("failed to save component ID to hosts: %v", err)
	}
}

func (m *Manager) RemoveHostsToTrace(hostsToRemove *_interface.HostsByComponentID) {
	m.Lock()
	defer m.Unlock()

	m.componentIDToHosts[hostsToRemove.ComponentID] = utils.RemoveFromSlice(m.componentIDToHosts[hostsToRemove.ComponentID], hostsToRemove.Hosts)
	m.hostsToTrace = createHostsToTrace(m.componentIDToHosts)
	if err := m.saveComponentIDToHosts(); err != nil {
		// TODO: consider retrying
		log.Errorf("failed to save component ID to hosts: %v", err)
	}
}

// initHostToTrace will fetch hosts per component map from secret and set manager initial state.
func (m *Manager) initHostToTrace() {
	componentIDToHosts, err := m.getComponentIDToHosts()
	if err != nil {
		log.Warnf("Failed to get component ID to hosts: %v", err)
	}

	if componentIDToHosts == nil {
		componentIDToHosts = make(map[string][]string)
	}

	m.componentIDToHosts = componentIDToHosts
	m.hostsToTrace = createHostsToTrace(m.componentIDToHosts)

	log.Infof("Successfully initialized host to trace state. "+
		"hostsToTrace=%+v, componentIDToHosts=%+v", m.hostsToTrace, m.componentIDToHosts)
}

// getComponentIDToHosts will fetch hosts per component map from secret.
func (m *Manager) getComponentIDToHosts() (map[string][]string, error) {
	var s *corev1.Secret
	var err error

	if s, err = m.Get(m.hostToTraceSecretMeta); err != nil {
		if errors.IsNotFound(err) {
			log.Infof("Secret not found: %v", err)
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get secret: %v", err)
	}

	dataB := s.Data[dataFieldName]
	if len(dataB) == 0 {
		return nil, fmt.Errorf("data in secret is empty: %v", err)
	}

	componentIDToHosts := make(map[string][]string)
	if err := json.Unmarshal(dataB, &componentIDToHosts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret data: %v", err)
	}

	return componentIDToHosts, nil
}

// createHostsToTrace will create a union of hosts need to be traced from all components.
func createHostsToTrace(componentIDToHosts map[string][]string) (ret []string) {
	hostsToLearnUnion := make(map[string]struct{})
	for _, hosts := range componentIDToHosts {
		for _, host := range hosts {
			hostsToLearnUnion[host] = struct{}{}
		}
	}

	for host := range hostsToLearnUnion {
		ret = append(ret, host)
	}

	return ret
}

// saveComponentIDToHosts will save hosts per component map to secret.
func (m *Manager) saveComponentIDToHosts() error {
	var s *corev1.Secret
	var err error

	if s, err = createSecret(m.componentIDToHosts, m.hostToTraceSecretMeta); err != nil {
		return fmt.Errorf("failed to create secret object: %v", err)
	}

	if _, err := m.CreateOrUpdate(s); err != nil {
		return fmt.Errorf("failed to create or update secret: %v", err)
	}

	return nil
}

func createSecret(hosts map[string][]string, meta metav1.ObjectMeta) (*corev1.Secret, error) {
	dataB, err := json.Marshal(hosts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal component to hosts map: %v", err)
	}

	return &corev1.Secret{
		ObjectMeta: meta,
		Data: map[string][]byte{
			dataFieldName: dataB,
		},
	}, nil
}
