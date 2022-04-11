package _interface

type HostsToTrace struct {
	Hosts       []string
}

type Getter interface {
	HostsToTrace() []string
}

type Setter interface {
	SetHostsToTrace(hostsToTrace *HostsToTrace)
}
