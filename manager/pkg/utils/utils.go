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

func RemoveFromSlice(from, toRemove []string) []string {
	ret := []string{}
	hostsToRemove := map[string]bool{}
	for _, host := range toRemove {
		hostsToRemove[host] = true
	}
	for _, host := range from {
		if !hostsToRemove[host] {
			ret = append(ret, host)
		}
	}
	return ret
}

func AddToSlice(from, toAdd []string) []string {
	ret := []string{}
	hostsToAdd := map[string]bool{}
	for _, host := range toAdd {
		hostsToAdd[host] = true
		ret = append(ret, host)
	}
	for _, host := range from {
		if !hostsToAdd[host] {
			ret = append(ret, host)
		}
	}
	return ret
}
