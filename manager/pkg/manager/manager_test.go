package manager

import (
	"reflect"
	"sort"
	"testing"

	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Portshift/go-utils/k8s/secret"
	api "github.com/apiclarity/trace-sampling-manager/api/grpc_gen/trace-sampling-manager-service"
	_interface "github.com/apiclarity/trace-sampling-manager/manager/pkg/manager/interface"
)

func Test_createHostsToTrace(t *testing.T) {
	type args struct {
		componentIDToHosts map[api.ComponentID][]string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []string
	}{
		{
			name: "simple union",
			args: args{
				componentIDToHosts: map[api.ComponentID][]string{
					api.ComponentID_SPEC_RECONSTRUCTOR: {"host:port"},
					api.ComponentID_TRACE_ANALYZER:     {"host2:port2"},
				},
			},
			wantRet: []string{"host:port", "host2:port2"},
		},
		{
			name: "same host on both",
			args: args{
				componentIDToHosts: map[api.ComponentID][]string{
					api.ComponentID_SPEC_RECONSTRUCTOR: {"host:port"},
					api.ComponentID_TRACE_ANALYZER:     {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
		{
			name: "empty list in ComponentID_SPEC_RECONSTRUCTOR",
			args: args{
				componentIDToHosts: map[api.ComponentID][]string{
					api.ComponentID_SPEC_RECONSTRUCTOR: nil,
					api.ComponentID_TRACE_ANALYZER:     {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
		{
			name: "empty list in ComponentID_TRACE_ANALYZER",
			args: args{
				componentIDToHosts: map[api.ComponentID][]string{
					api.ComponentID_TRACE_ANALYZER:     nil,
					api.ComponentID_SPEC_RECONSTRUCTOR: {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
		{
			name: "empty list in all",
			args: args{
				componentIDToHosts: map[api.ComponentID][]string{
					api.ComponentID_TRACE_ANALYZER:     nil,
					api.ComponentID_SPEC_RECONSTRUCTOR: nil,
				},
			},
			wantRet: nil,
		},
		{
			name: "missing component ID",
			args: args{
				componentIDToHosts: map[api.ComponentID][]string{
					api.ComponentID_TRACE_ANALYZER: {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet := createHostsToTrace(tt.args.componentIDToHosts)
			sort.Strings(gotRet)
			sort.Strings(tt.wantRet)
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("createHostsToTrace() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestManager_getComponentIDToHosts(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	secretMockHandler := secret.NewMockHandler(mock)

	secretGroupResource := schema.GroupResource{
		Group:    "v1",
		Resource: "secret",
	}

	testComponentIDToHosts := map[api.ComponentID][]string{
		api.ComponentID_TRACE_ANALYZER:     {"host:80", "host"},
		api.ComponentID_SPEC_RECONSTRUCTOR: {"host:8080"},
	}

	testSecret, _ := createSecret(testComponentIDToHosts)

	type fields struct {
		expectSecretHandler func(handler *secret.MockHandler)
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[api.ComponentID][]string
		wantErr bool
	}{
		{
			name: "secret not found - expected empty result",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(nil, errors.NewNotFound(secretGroupResource, ""))
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "failed to get secret - expected error",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(nil, errors.NewBadRequest(""))
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "data in secret is empty",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(&corev1.Secret{
						ObjectMeta: hostToTraceSecretMeta,
						Data:       map[string][]byte{},
					}, nil)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to unmarshal secret data",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(&corev1.Secret{
						ObjectMeta: hostToTraceSecretMeta,
						Data: map[string][]byte{
							dataFieldName: []byte("not a map"),
						},
					}, nil)
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(testSecret, nil)
				},
			},
			want:    testComponentIDToHosts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Handler: secretMockHandler,
			}
			tt.fields.expectSecretHandler(secretMockHandler)
			got, err := m.getComponentIDToHosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("getComponentIDToHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getComponentIDToHosts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_initHostToTrace(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	secretMockHandler := secret.NewMockHandler(mock)

	testComponentIDToHosts := map[api.ComponentID][]string{
		api.ComponentID_TRACE_ANALYZER:     {"host:80", "host"},
		api.ComponentID_SPEC_RECONSTRUCTOR: {"host:8080"},
	}

	testSecret, _ := createSecret(testComponentIDToHosts)
	testSecretNilMap, _ := createSecret(nil)
	testSecretEmptyMap, _ := createSecret(map[api.ComponentID][]string{})

	type fields struct {
		expectSecretHandler func(handler *secret.MockHandler)
	}
	tests := []struct {
		name                       string
		fields                     fields
		wantErr                    bool
		expectedComponentIDToHosts map[api.ComponentID][]string
		expectedHostsToTrace       []string
	}{
		{
			name: "failed to get component ID",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(nil, errors.NewBadRequest(""))
				},
			},
			wantErr:                    true,
			expectedComponentIDToHosts: nil,
			expectedHostsToTrace:       nil,
		},
		{
			name: "Successfully initialized host to trace state",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(testSecret, nil)
				},
			},
			wantErr:                    false,
			expectedComponentIDToHosts: testComponentIDToHosts,
			expectedHostsToTrace:       []string{"host:80", "host", "host:8080"},
		},
		{
			name: "Nil map in secret initialized host to trace state",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(testSecretNilMap, nil)
				},
			},
			wantErr:                    false,
			expectedComponentIDToHosts: make(map[api.ComponentID][]string),
			expectedHostsToTrace:       nil,
		},
		{
			name: "Empty map in secret initialized host to trace state",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(testSecretEmptyMap, nil)
				},
			},
			wantErr:                    false,
			expectedComponentIDToHosts: make(map[api.ComponentID][]string),
			expectedHostsToTrace:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Handler: secretMockHandler,
			}
			tt.fields.expectSecretHandler(secretMockHandler)
			if err := m.initHostToTrace(); (err != nil) != tt.wantErr {
				t.Errorf("initHostToTrace() error = %v, wantErr %v", err, tt.wantErr)
			}
			sort.Strings(m.hostsToTrace)
			sort.Strings(tt.expectedHostsToTrace)
			if !reflect.DeepEqual(m.hostsToTrace, tt.expectedHostsToTrace) {
				t.Errorf("initHostToTrace() hostsToTrace missmatch. got = %+v, expected = %+v", m.hostsToTrace, tt.expectedHostsToTrace)
			}
			if !reflect.DeepEqual(m.componentIDToHosts, tt.expectedComponentIDToHosts) {
				t.Errorf("initHostToTrace() componentIDToHosts missmatch. got = %+v, expected = %+v", m.componentIDToHosts, tt.expectedComponentIDToHosts)
			}
			if !tt.wantErr {
				// assignment validation that we do not crash on assignment
				m.componentIDToHosts[api.ComponentID_SPEC_RECONSTRUCTOR] = []string{"test"}
			}
		})
	}
}

func TestManager_SetHostsToTrace(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	secretMockHandler := secret.NewMockHandler(mock)

	testComponentIDToHostsBefore := map[api.ComponentID][]string{
		api.ComponentID_SPEC_RECONSTRUCTOR: {"host:8080"},
	}
	testHostToTraceBefore := []string{"host:8080"}

	testComponentIDToHostsBeforeWithTraceAnalyzer := map[api.ComponentID][]string{
		api.ComponentID_TRACE_ANALYZER:     {"blalala"},
		api.ComponentID_SPEC_RECONSTRUCTOR: {"host:8080"},
	}
	testHostToTraceBeforeWithTraceAnalyzer := []string{"host:8080", "blalala"}

	hostsToTraceInput := &_interface.HostsToTrace{
		Hosts:       []string{"host:80", "host"},
		ComponentID: api.ComponentID_TRACE_ANALYZER,
	}

	testComponentIDToHostsAfter := map[api.ComponentID][]string{
		api.ComponentID_TRACE_ANALYZER:     {"host:80", "host"},
		api.ComponentID_SPEC_RECONSTRUCTOR: {"host:8080"},
	}
	testHostToTraceAfter := []string{"host:80", "host", "host:8080"}

	testSecret, _ := createSecret(testComponentIDToHostsAfter)

	type fields struct {
		expectSecretHandler func(handler *secret.MockHandler)
		hostsToTrace        []string
		componentIDToHosts  map[api.ComponentID][]string
	}
	type args struct {
		hostsToTrace *_interface.HostsToTrace
	}
	tests := []struct {
		name                       string
		fields                     fields
		args                       args
		expectedComponentIDToHosts map[api.ComponentID][]string
		expectedHostsToTrace       []string
	}{
		{
			name: "empty ComponentID_TRACE_ANALYZER data",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().CreateOrUpdate(testSecret).Return(nil, nil)
				},
				hostsToTrace:       testHostToTraceBefore,
				componentIDToHosts: testComponentIDToHostsBefore,
			},
			args: args{
				hostsToTrace: hostsToTraceInput,
			},
			expectedComponentIDToHosts: testComponentIDToHostsAfter,
			expectedHostsToTrace:       testHostToTraceAfter,
		},
		{
			name: "verify that component id data is overridden",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().CreateOrUpdate(testSecret).Return(nil, nil)
				},
				hostsToTrace:       testHostToTraceBeforeWithTraceAnalyzer,
				componentIDToHosts: testComponentIDToHostsBeforeWithTraceAnalyzer,
			},
			args: args{
				hostsToTrace: hostsToTraceInput,
			},
			expectedComponentIDToHosts: testComponentIDToHostsAfter,
			expectedHostsToTrace:       testHostToTraceAfter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Handler:            secretMockHandler,
				hostsToTrace:       tt.fields.hostsToTrace,
				componentIDToHosts: tt.fields.componentIDToHosts,
			}
			tt.fields.expectSecretHandler(secretMockHandler)

			m.SetHostsToTrace(tt.args.hostsToTrace)

			sort.Strings(m.hostsToTrace)
			sort.Strings(tt.expectedHostsToTrace)
			if !reflect.DeepEqual(m.hostsToTrace, tt.expectedHostsToTrace) {
				t.Errorf("initHostToTrace() hostsToTrace missmatch. got = %+v, expected = %+v", m.hostsToTrace, tt.expectedHostsToTrace)
			}
			if !reflect.DeepEqual(m.componentIDToHosts, tt.expectedComponentIDToHosts) {
				t.Errorf("initHostToTrace() componentIDToHosts missmatch. got = %+v, expected = %+v", m.componentIDToHosts, tt.expectedComponentIDToHosts)
			}
		})
	}
}
