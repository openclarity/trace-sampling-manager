package _interface

type HostsToTrace struct {
	Hosts       []string
	ComponentID string
}

type Getter interface {
	HostsToTrace() []string
	HostsToTraceByComponentID(id string) []string
}

type Setter interface {
	SetHostsToTrace(hostsToTrace *HostsToTrace)
}
