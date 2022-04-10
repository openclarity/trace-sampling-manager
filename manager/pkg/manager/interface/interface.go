package _interface

import api "github.com/apiclarity/trace-sampling-manager/api/grpc_gen/trace-sampling-manager-service"

type HostsToTrace struct {
	Hosts       []string
	ComponentID api.ComponentID
}

type Getter interface {
	HostsToTrace() []string
	HostsToTraceByComponentID(api.ComponentID) []string
}

type Setter interface {
	SetHostsToTrace(hostsToTrace *HostsToTrace)
}
