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

package utils

import (
	"reflect"
	"testing"
)

func TestRemoveFromSlice(t *testing.T) {
	type args struct {
		from     []string
		toRemove []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "sanity",
			args: args{
				from: []string{
					"host1", "host2", "host3",
				},
				toRemove: []string{
					"host2",
				},
			},
			want: []string{
				"host1", "host3",
			},
		},
		{
			name: "nothing to remove",
			args: args{
				from: []string{
					"host1", "host2", "host3",
				},
				toRemove: []string{},
			},
			want: []string{
				"host1", "host2", "host3",
			},
		},
		{
			name: "remove all",
			args: args{
				from: []string{
					"host1", "host2", "host3",
				},
				toRemove: []string{
					"host1", "host2", "host3",
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFromSlice(tt.args.from, tt.args.toRemove); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
