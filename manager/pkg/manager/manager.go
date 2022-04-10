package manager

import (
	"encoding/json"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_secret "github.com/Portshift/go-utils/k8s/secret"

	api "github.com/apiclarity/trace-sampling-manager/api/grpc_gen/trace-sampling-manager-service"
	"github.com/apiclarity/trace-sampling-manager/manager/pkg/grpc"
	_interface "github.com/apiclarity/trace-sampling-manager/manager/pkg/manager/interface"
	"github.com/apiclarity/trace-sampling-manager/manager/pkg/rest"
)

const (
	dataFieldName              = "hosts-per-component-map"
	hostToTraceSecretName      = "host-to-trace"
	hostToTraceSecretNamespace = "portshift"
)

var hostToTraceSecretMeta = metav1.ObjectMeta{
	Name:      hostToTraceSecretName,
	Namespace: hostToTraceSecretNamespace,
	Labels:    map[string]string{"owner": "portshift"},
}

type Config struct {
	RestServerPort int
	GRPCServerPort int
}

type Manager struct {
	_secret.Handler
	restServer *rest.Server
	grpcServer *grpc.Server

	hostsToTrace       []string
	componentIDToHosts map[api.ComponentID][]string

	sync.RWMutex
}

func Create(clientset kubernetes.Interface, conf *Config) (*Manager, error) {
	var err error
	m := &Manager{
		Handler: _secret.NewHandler(clientset),
	}

	if err := m.initHostToTrace(); err != nil {
		log.Warnf("failed to init hosts to trace list - initializing with empty list: %v", err)
	}

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

func (m *Manager) HostsToTraceByComponentID(id api.ComponentID) []string {
	m.RLock()
	defer m.RUnlock()

	return m.componentIDToHosts[id]
}

func (m *Manager) SetHostsToTrace(hostsToTrace *_interface.HostsToTrace) {
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

// initHostToTrace will fetch hosts per component map from secret and set manager initial state
func (m *Manager) initHostToTrace() error {
	componentIDToHosts, err := m.getComponentIDToHosts()
	if err != nil {
		return fmt.Errorf("failed to get component ID to hosts: %v", err)
	}

	if componentIDToHosts == nil {
		componentIDToHosts = make(map[api.ComponentID][]string)
	}

	m.componentIDToHosts = componentIDToHosts
	m.hostsToTrace = createHostsToTrace(m.componentIDToHosts)

	log.Infof("Successfully initialized host to trace state. "+
		"hostsToTrace=%+v, componentIDToHosts=%+v", m.hostsToTrace, m.componentIDToHosts)

	return nil
}

// getComponentIDToHosts will fetch hosts per component map from secret
func (m *Manager) getComponentIDToHosts() (map[api.ComponentID][]string, error) {
	var s *corev1.Secret
	var err error

	if s, err = m.Get(hostToTraceSecretMeta); err != nil {
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

	componentIDToHosts := make(map[api.ComponentID][]string)
	if err := json.Unmarshal(dataB, &componentIDToHosts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret data: %v", err)
	}

	return componentIDToHosts, nil
}

// createHostsToTrace will create a union of hosts need to be traced from all components
func createHostsToTrace(componentIDToHosts map[api.ComponentID][]string) (ret []string) {
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

// saveComponentIDToHosts will save hosts per component map to secret
func (m *Manager) saveComponentIDToHosts() error {
	var s *corev1.Secret
	var err error

	if s, err = createSecret(m.componentIDToHosts); err != nil {
		return fmt.Errorf("failed to create secret object: %v", err)
	}

	if _, err := m.CreateOrUpdate(s); err != nil {
		return fmt.Errorf("failed to create or update secret: %v", err)
	}

	return nil
}

func createSecret(hosts map[api.ComponentID][]string) (*corev1.Secret, error) {
	dataB, err := json.Marshal(hosts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal component to hosts map: %v", err)
	}

	return &corev1.Secret{
		ObjectMeta: hostToTraceSecretMeta,
		Data: map[string][]byte{
			dataFieldName: dataB,
		},
	}, nil
}
