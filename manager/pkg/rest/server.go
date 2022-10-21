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

package rest

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"github.com/openclarity/trace-sampling-manager/api/server/restapi"
	"github.com/openclarity/trace-sampling-manager/api/server/restapi/operations"
	manager "github.com/openclarity/trace-sampling-manager/manager/pkg/manager/interface"
)

type Config struct {
	EnableTLS                  bool
	RestServerPort             int
	RestServerTLSPort          int
	GRPCServerPort             int
	HostToTraceSecretName      string
	HostToTraceSecretNamespace string
	HostToTraceSecretOwnerName string
	TLSServerCertFilePath      string
	TLSServerKeyFilePath       string
	RootCertFilePath           string
}

type Server struct {
	server *restapi.Server
	manager.Getter
}

func CreateRESTServer(config *Config, getter manager.Getter) (*Server, error) {
	s := &Server{
		Getter: getter,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to load swagger spec: %v", err)
	}

	api := operations.NewTraceSamplingManagerAPI(swaggerSpec)

	api.GetHostsToTraceHandler = operations.GetHostsToTraceHandlerFunc(func(params operations.GetHostsToTraceParams) middleware.Responder {
		hosts := s.HostsToTrace()
		if hosts == nil {
			hosts = []string{}
		}
		return operations.NewGetHostsToTraceOK().WithPayload(&operations.GetHostsToTraceOKBody{
			Hosts: hosts,
		})
	})

	server := restapi.NewServer(api)

	server.ConfigureFlags()
	server.ConfigureAPI()
	server.Port = config.RestServerPort

	// We want to serve both http and https
	if config.EnableTLS {
		server.EnabledListeners = []string{"https", "http"}
		server.TLSCertificate = flags.Filename(config.TLSServerCertFilePath)
		server.TLSCertificateKey = flags.Filename(config.TLSServerKeyFilePath)
		server.TLSPort = config.RestServerTLSPort
	}

	s.server = server

	return s, nil
}

func (s *Server) Start(errChan chan struct{}) {
	log.Infof("Starting REST server")
	go func() {
		if err := s.server.Serve(); err != nil {
			log.Errorf("Failed to serve REST server: %v", err)
			errChan <- struct{}{}
		}
	}()
}

func (s *Server) Stop() {
	log.Infof("Stopping REST server")
	if s.server != nil {
		if err := s.server.Shutdown(); err != nil {
			log.Errorf("Failed to shutdown REST server: %v", err)
		}
	}
}
