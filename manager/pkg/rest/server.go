package rest

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/apiclarity/trace-sampling-manager/api/server/restapi"
	"github.com/apiclarity/trace-sampling-manager/api/server/restapi/operations"
	manager "github.com/apiclarity/trace-sampling-manager/manager/pkg/manager/interface"
)

type Server struct {
	server *restapi.Server
	manager.Getter
}

func CreateRESTServer(port int, getter manager.Getter) (*Server, error) {
	s := &Server{
		Getter: getter,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to load swagger spec: %v", err)
	}

	api := operations.NewTraceSamplingManagerAPI(swaggerSpec)

	api.GetHostsToTraceHandler = operations.GetHostsToTraceHandlerFunc(func(params operations.GetHostsToTraceParams) middleware.Responder {
		return operations.NewGetHostsToTraceOK().WithPayload(&operations.GetHostsToTraceOKBody{
			Hosts: s.HostsToTrace(),
		})
	})

	server := restapi.NewServer(api)

	server.ConfigureFlags()
	server.ConfigureAPI()
	server.Port = port

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
