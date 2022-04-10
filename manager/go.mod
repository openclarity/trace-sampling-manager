module github.com/apiclarity/trace-sampling-manager/manager

go 1.16

require (
	github.com/Portshift/go-utils v0.0.0-20220410101458-977521fb3634
	github.com/apiclarity/trace-sampling-manager/api v0.0.0
	github.com/go-openapi/loads v0.20.3
	github.com/go-openapi/runtime v0.20.0
	github.com/golang/mock v1.6.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/urfave/cli v1.22.5
	google.golang.org/grpc v1.42.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
)

replace github.com/apiclarity/trace-sampling-manager/api v0.0.0 => ./../api
