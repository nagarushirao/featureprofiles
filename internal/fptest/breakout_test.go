// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fptest

import "testing"

func TestAristaBreakoutParent(t *testing.T) {
	cases := []struct{ arg, want string }{
		{"Port-Channel1", ""},
		{"Loopback1", ""},
		{"Ethernet3", ""},
		{"Ethernet3/1", "Ethernet3"},
		{"Ethernet3/2", "Ethernet3"},
		{"Ethernet3/3", "Ethernet3"},
		{"Ethernet3/4", "Ethernet3"},
		{"Ethernet3/3/1", "Ethernet3/3"},
		{"Ethernet3/3/3", "Ethernet3/3"},
		{"Ethernet3/3/5", "Ethernet3/3"},
		{"Ethernet3/3/7", "Ethernet3/3"},
	}
	for _, c := range cases {
		got := aristaBreakoutParent(c.arg)
		if got != c.want {
			t.Errorf("Result got %q, want %q", got, c.want)
		}
	}
}

func TestJuniperBreakoutParent(t *testing.T) {
	cases := []struct{ arg, want string }{
		{"ae1", ""},
		{"lo0", ""},
		{"et-0/2", ""},
		{"et-0/2:0", "et-0/2"},
		{"et-0/2:1", "et-0/2"},
		{"et-0/2:2", "et-0/2"},
		{"et-0/2:3", "et-0/2"},
		{"et-0/0/2", ""},
		{"et-0/0/2:0", "et-0/0/2"},
		{"et-0/0/2:1", "et-0/0/2"},
		{"et-0/0/2:2", "et-0/0/2"},
		{"et-0/0/2:3", "et-0/0/2"},
	}
	for _, c := range cases {
		got := juniperBreakoutParent(c.arg)
		if got != c.want {
			t.Errorf("Result got %q, want %q", got, c.want)
		}
	}
}

func TestCiscoBreakoutParent(t *testing.T) {
	cases := []struct{ arg, want string }{
		{"Bundle-Ether1", ""},
		{"Loopback0", ""},
		{"TenGigE0/0/0/35/0", "Optics0/0/0/35"},
		{"TenGigE0/0/0/35/1", "Optics0/0/0/35"},
		{"TenGigE0/0/0/35/2", "Optics0/0/0/35"},
		{"TenGigE0/0/0/35/3", "Optics0/0/0/35"},
		{"HundredGigE0/2/0/35/0", "Optics0/2/0/35"},
		{"HundredGigE0/2/0/35/1", "Optics0/2/0/35"},
		{"HundredGigE0/2/0/35/2", "Optics0/2/0/35"},
		{"HundredGigE0/2/0/35/3", "Optics0/2/0/35"},
	}
	for _, c := range cases {
		got := ciscoBreakoutParent(c.arg)
		if got != c.want {
			t.Errorf("Result got %q, want %q", got, c.want)
		}
	}
}
