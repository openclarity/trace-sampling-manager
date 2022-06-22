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

package manager

import (
	"reflect"
	"sort"
	"testing"

	"github.com/Portshift/go-utils/k8s/secret"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"

	_interface "github.com/openclarity/trace-sampling-manager/manager/pkg/manager/interface"
)

const (
	SpecReconstructor = "SPEC_RECONSTRUCTOR"
	TraceAnalyzer     = "TRACE_ANALYZER"
)

func Test_createHostsToTrace(t *testing.T) {
	type args struct {
		componentIDToHosts map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []string
	}{
		{
			name: "simple union",
			args: args{
				componentIDToHosts: map[string][]string{
					SpecReconstructor: {"host:port"},
					TraceAnalyzer:     {"host2:port2"},
				},
			},
			wantRet: []string{"host:port", "host2:port2"},
		},
		{
			name: "same host on both",
			args: args{
				componentIDToHosts: map[string][]string{
					SpecReconstructor: {"host:port"},
					TraceAnalyzer:     {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
		{
			name: "empty list in ComponentID_SPEC_RECONSTRUCTOR",
			args: args{
				componentIDToHosts: map[string][]string{
					SpecReconstructor: nil,
					TraceAnalyzer:     {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
		{
			name: "empty list in ComponentID_TRACE_ANALYZER",
			args: args{
				componentIDToHosts: map[string][]string{
					TraceAnalyzer:     nil,
					SpecReconstructor: {"host:port"},
				},
			},
			wantRet: []string{"host:port"},
		},
		{
			name: "empty list in all",
			args: args{
				componentIDToHosts: map[string][]string{
					TraceAnalyzer:     nil,
					SpecReconstructor: nil,
				},
			},
			wantRet: nil,
		},
		{
			name: "missing component ID",
			args: args{
				componentIDToHosts: map[string][]string{
					TraceAnalyzer: {"host:port"},
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

	testComponentIDToHosts := map[string][]string{
		TraceAnalyzer:     {"host:80", "host"},
		SpecReconstructor: {"host:8080"},
	}

	testSecret, _ := createSecret(testComponentIDToHosts)

	type fields struct {
		expectSecretHandler func(handler *secret.MockHandler)
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string][]string
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

	testComponentIDToHosts := map[string][]string{
		TraceAnalyzer:     {"host:80", "host"},
		SpecReconstructor: {"host:8080"},
	}

	testSecret, _ := createSecret(testComponentIDToHosts)
	testSecretNilMap, _ := createSecret(nil)
	testSecretEmptyMap, _ := createSecret(map[string][]string{})

	type fields struct {
		expectSecretHandler func(handler *secret.MockHandler)
	}
	tests := []struct {
		name                       string
		fields                     fields
		expectedComponentIDToHosts map[string][]string
		expectedHostsToTrace       []string
	}{
		{
			name: "failed to get component ID",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(nil, errors.NewBadRequest(""))
				},
			},
			expectedComponentIDToHosts: map[string][]string{},
			expectedHostsToTrace:       nil,
		},
		{
			name: "Successfully initialized host to trace state",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(testSecret, nil)
				},
			},
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
			expectedComponentIDToHosts: make(map[string][]string),
			expectedHostsToTrace:       nil,
		},
		{
			name: "Empty map in secret initialized host to trace state",
			fields: fields{
				expectSecretHandler: func(handler *secret.MockHandler) {
					handler.EXPECT().Get(hostToTraceSecretMeta).Return(testSecretEmptyMap, nil)
				},
			},
			expectedComponentIDToHosts: make(map[string][]string),
			expectedHostsToTrace:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Handler: secretMockHandler,
			}
			tt.fields.expectSecretHandler(secretMockHandler)
			m.initHostToTrace()
			sort.Strings(m.hostsToTrace)
			sort.Strings(tt.expectedHostsToTrace)
			if !reflect.DeepEqual(m.hostsToTrace, tt.expectedHostsToTrace) {
				t.Errorf("initHostToTrace() hostsToTrace missmatch. got = %+v, expected = %+v", m.hostsToTrace, tt.expectedHostsToTrace)
			}
			if !reflect.DeepEqual(m.componentIDToHosts, tt.expectedComponentIDToHosts) {
				t.Errorf("initHostToTrace() componentIDToHosts missmatch. got = %+v, expected = %+v", m.componentIDToHosts, tt.expectedComponentIDToHosts)
			}
			// assignment validation that we do not crash on assignment
			m.componentIDToHosts[SpecReconstructor] = []string{"test"}
		})
	}
}

func TestManager_SetHostsToTrace(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	secretMockHandler := secret.NewMockHandler(mock)

	testComponentIDToHostsBefore := map[string][]string{
		SpecReconstructor: {"host:8080"},
	}
	testHostToTraceBefore := []string{"host:8080"}

	testComponentIDToHostsBeforeWithTraceAnalyzer := map[string][]string{
		TraceAnalyzer:     {"blalala"},
		SpecReconstructor: {"host:8080"},
	}
	testHostToTraceBeforeWithTraceAnalyzer := []string{"host:8080", "blalala"}

	hostsToTraceInput := &_interface.HostsByComponentID{
		Hosts:       []string{"host:80", "host"},
		ComponentID: TraceAnalyzer,
	}

	testComponentIDToHostsAfter := map[string][]string{
		TraceAnalyzer:     {"host:80", "host"},
		SpecReconstructor: {"host:8080"},
	}
	testHostToTraceAfter := []string{"host:80", "host", "host:8080"}

	testSecret, _ := createSecret(testComponentIDToHostsAfter)

	type fields struct {
		expectSecretHandler func(handler *secret.MockHandler)
		hostsToTrace        []string
		componentIDToHosts  map[string][]string
	}
	type args struct {
		hostsToTrace *_interface.HostsByComponentID
	}
	tests := []struct {
		name                       string
		fields                     fields
		args                       args
		expectedComponentIDToHosts map[string][]string
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
