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
	"reflect"
	"sort"
	"testing"

	api "github.com/apiclarity/trace-sampling-manager/api/grpc_gen/trace-sampling-manager-service"
)

func Test_createHost(t *testing.T) {
	type args struct {
		h *api.Host
	}
	tests := []struct {
		name    string
		args    args
		wantRet []string
	}{
		{
			name: "port is missing",
			args: args{
				h: &api.Host{
					Hostname: "foo",
					Port:     0,
				},
			},
			wantRet: []string{"foo"},
		},
		{
			name: "port is not missing",
			args: args{
				h: &api.Host{
					Hostname: "foo",
					Port:     8080,
				},
			},
			wantRet: []string{"foo:8080"},
		},
		{
			name: "default 80 port",
			args: args{
				h: &api.Host{
					Hostname: "foo",
					Port:     80,
				},
			},
			wantRet: []string{"foo:80", "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet := createHosts(tt.args.h)
			sort.Strings(gotRet)
			sort.Strings(tt.wantRet)
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("createHosts() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_createAPIHost(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		args    args
		wantRet *api.Host
		wantErr bool
	}{
		{
			name: "port is missing",
			args: args{
				host: "hostname.namespace",
			},
			wantRet: &api.Host{
				Hostname: "hostname.namespace",
				Port:     0,
			},
			wantErr: false,
		},
		{
			name: "port is not missing",
			args: args{
				host: "hostname.namespace:80",
			},
			wantRet: &api.Host{
				Hostname: "hostname.namespace",
				Port:     80,
			},
			wantErr: false,
		},
		{
			name: "invalid port",
			args: args{
				host: "hostname.namespace:svc",
			},
			wantRet: nil,
			wantErr: true,
		},
		{
			name: "empty host",
			args: args{
				host: "",
			},
			wantRet: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := createAPIHost(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("createAPIHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("createAPIHost() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_createAPIHosts(t *testing.T) {
	type args struct {
		hosts []string
	}
	tests := []struct {
		name    string
		args    args
		wantRet []*api.Host
	}{
		{
			name: "sanity",
			args: args{
				hosts: []string{"hostname1", "hostname2:80"},
			},
			wantRet: []*api.Host{
				{
					Hostname: "hostname1",
					Port:     0,
				},
				{
					Hostname: "hostname2",
					Port:     80,
				},
			},
		},
		{
			name: "1 invalid host",
			args: args{
				hosts: []string{"hostname1", "hostname2:invalid"},
			},
			wantRet: []*api.Host{
				{
					Hostname: "hostname1",
					Port:     0,
				},
			},
		},
		{
			name: "empty list",
			args: args{
				hosts: []string{},
			},
			wantRet: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := createAPIHosts(tt.args.hosts); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("createAPIHosts() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
