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

import (
	"strings"

	"github.com/openconfig/ondatra"
)

// aristaBreakoutParent returns the parent port name of the breakout.
//
// Arista ports are named EthernetY/Z (port Y on fixed form
// factor) or EthernetX/Y/Z (linecard X, port Y).  Z is the
// breakout number.  Non-breakout ports have Z always 1.  To get
// the parent, we just strip /Z at the end.
func aristaBreakoutParent(name string) string {
	if !strings.HasPrefix(name, "Ethernet") {
		return "" // Only "Ethernet" can be broken out.
	}
	if i := strings.LastIndexByte(name, '/'); i >= 0 {
		return name[:i]
	}
	return "" // Not a breakout.
}

// juniperBreakoutParent returns the parent port name of the breakout.
//
// Juniper ports are named et-W/X/Y (fpc W, pic X, port Y), and
// the breakout ports are named et-W/X/Y:Z.  To get the parent, we
// just strip :Z at the end.
func juniperBreakoutParent(name string) string {
	if !strings.HasPrefix(name, "et-") {
		return "" // Only "et" can be broken out.
	}
	if i := strings.LastIndexByte(name, ':'); i >= 0 {
		return name[:i]
	}
	return "" // Not a breakout.
}

// ciscoBreakoutParent returns the parent optics name of the breakout.
//
// Cisco controller optics are named Optics0/X/0/Y (slot X, port Y),
// and the breakout interfaces are named
// {TenGigE,HundredGigE}0/X/0/Y/Z.  To get the parent, we strip the /Z
// and replace any prefix ending in GigE with "Optics".
func ciscoBreakoutParent(name string) string {
	i := strings.Index(name, "GigE")
	if i < 0 {
		return "" // GigE not found.
	}
	j := strings.LastIndexByte(name, '/')
	if j < 0 {
		return "" // Not an interface name.
	}
	return "Optics" + name[i+4:j]
}

// BreakoutParent returns the parent port name of the breakout
// according to vendor convention.
func BreakoutParent(port *ondatra.Port) string {
	switch port.Device().Vendor() {
	case ondatra.ARISTA:
		return aristaBreakoutParent(port.Name())
	case ondatra.JUNIPER:
		return juniperBreakoutParent(port.Name())
	case ondatra.CISCO:
		return ciscoBreakoutParent(port.Name())
	}
	return "" // Vendor is not supported yet.
}

// CollateBreakoutPorts sorts the ports into a breakout map mapping
// from the physical port name to the logical ports.  If a port is not
// part of the breakout, the key is the empty string.
func CollateBreakoutPorts(ports []*ondatra.Port) map[string][]*ondatra.Port {
	m := make(map[string][]*ondatra.Port)
	for _, dp := range ports {
		phy := BreakoutParent(dp)
		m[phy] = append(m[phy], dp)
	}
	return m
}
