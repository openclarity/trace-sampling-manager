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

package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	api "github.com/openclarity/trace-sampling-manager/api/grpc_gen/trace-sampling-manager-service"
	manager "github.com/openclarity/trace-sampling-manager/manager/pkg/manager/interface"
)

const hostnamePortSeparator = ":"

type Server struct {
	listener net.Listener
	server   *grpc.Server
	port     int
	manager.Setter
	manager.Getter
}

func (s *Server) GetHostsToTrace(_ context.Context, request *api.GetHostsToTraceRequest) (*api.GetHostsToTraceResponse, error) {
	log.Debugf("Get hosts to trace was called. request=%+v", request)

	ret := &api.GetHostsToTraceResponse{
		Hosts: createAPIHosts(s.Getter.HostsToTrace()),
	}

	return ret, nil
}

func createAPIHosts(hosts []string) (ret []*api.Host) {
	for _, host := range hosts {
		if apiHost, err := createAPIHost(host); err != nil {
			log.Warnf("failed to create API host (%v) - skipping: %v", host, err)
		} else {
			ret = append(ret, apiHost)
		}
	}

	return ret
}

func createAPIHost(host string) (ret *api.Host, err error) {
	var hostname string
	var port int

	if host == "" {
		// should not happen
		return nil, fmt.Errorf("host is empty - skipping")
	}

	hostnameAndPort := strings.Split(host, hostnamePortSeparator)
	hostname = hostnameAndPort[0]

	if len(hostnameAndPort) > 1 {
		//nolint:gosec
		// TODO figure out safer alternative to use strconv.Atoi
		port, err = strconv.Atoi(hostnameAndPort[1])
		if err != nil {
			// should not happen
			return nil, fmt.Errorf("invalid port number (%v): %v", hostnameAndPort[1], err)
		}
	}

	return &api.Host{
		Hostname: hostname,
		Port:     int32(port),
	}, nil
}

func (s *Server) SetHostsToTrace(_ context.Context, request *api.HostsToTraceRequest) (*api.Empty, error) {
	log.Debugf("Got hosts to trace. request=%+v", request)

	s.Setter.SetHostsToTrace(&manager.HostsByComponentID{
		Hosts: createHostsList(request.Hosts),
	})

	return &api.Empty{}, nil
}

func (s *Server) SetHostsToRemove(_ context.Context, request *api.HostsToRemoveRequest) (*api.Empty, error) {
	log.Debugf("Got hosts to remove. request=%+v", request)

	s.Setter.SetHostsToRemove(&manager.HostsByComponentID{
		Hosts: createHostsList(request.Hosts),
	})

	return &api.Empty{}, nil
}

func createHostsList(hosts []*api.Host) (ret []string) {
	for _, host := range hosts {
		ret = append(ret, createHosts(host)...)
	}

	return ret
}

// createHosts will create hosts in the format of `hostname:port` if port exist, otherwise will return only hostname
// Note: The function will return both `hostname:port` and `hostname` in case port is the default HTTP port (80).
func createHosts(h *api.Host) (ret []string) {
	// TODO: we might need to create multiple hosts from a single api.Host
	// example: hostname=foo, port=8080 ==> host=[foo:8080, foo.namespace:8080, ....]
	if h.Port > 0 {
		ret = append(ret, h.Hostname+hostnamePortSeparator+strconv.Itoa(int(h.Port)))
	}

	if h.Port == 0 || h.Port == 80 {
		ret = append(ret, h.Hostname)
	}

	return ret
}

func NewServer(port int, setter manager.Setter, getter manager.Getter) *Server {
	server := &Server{
		listener: nil,
		server:   grpc.NewServer(),
		port:     port,
		Setter:   setter,
		Getter:   getter,
	}

	api.RegisterTraceSamplingManagerServer(server.server, server)

	return server
}

// Start starts the server run.
func (s *Server) Start(errChan chan struct{}) error {
	log.Infof("Starting GRPC server")

	listenAddr := ":" + strconv.Itoa(s.port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("GRPC server failed to listen: %v", err)
	}

	s.listener = listener

	go func() {
		if err := s.server.Serve(s.listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Errorf("Failed to serve GRPC server: %v", err)
			if errChan != nil {
				errChan <- struct{}{}
			}
		}
	}()

	return nil
}

// Stop gracefully shuts down the server.
func (s *Server) Stop() {
	log.Infof("Stopping GRPC server")
	if s.server != nil {
		s.server.GracefulStop()
	}
	if s.listener != nil {
		_ = s.listener.Close()
	}
}
