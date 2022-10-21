module github.com/openclarity/trace-sampling-manager/manager

go 1.16

require (
	github.com/Portshift/go-utils v0.0.0-20220421083203-89265d8a6487
	github.com/go-openapi/loads v0.21.1
	github.com/go-openapi/runtime v0.24.2
	github.com/golang/mock v1.6.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/openclarity/trace-sampling-manager/api v0.0.0
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.42.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
)

replace github.com/openclarity/trace-sampling-manager/api v0.0.0 => ./../api
