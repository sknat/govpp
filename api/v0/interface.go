// Copyright (C) 2021 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"go/types"
	"net"
)

type RxMode uint32

const (
	// TODO: RxModeUnknown, RxModePolling, RxModeInterrupt, RxModeAdaptive, RxModeDefault

	UnknownRxMode RxMode = 0
	Polling       RxMode = 1
	Interrupt     RxMode = 2
	Adaptative    RxMode = 3
	DefaultRxMode RxMode = 4

	AllQueues = ^uint32(0)
)

type TapFlags uint32

const (
	TapFlagNone        TapFlags = 0
	TapFlagGSO         TapFlags = 1
	TapFlagCsumOffload TapFlags = 2
	TapFlagPersist     TapFlags = 4
	TapFlagAttach      TapFlags = 8
	TapFlagTun         TapFlags = 16
	TapGROCoalesce     TapFlags = 32
)

const (
	MaxMtu           = 9216
	DefaultQueueSize = 1024
)

type Interface struct {
	// pre-required for calls towards Set and Get
	SwIfIndex uint32
	// TODO: Set & Get calls set appropriate values here

	Gso       bool
	PromiscOn bool

	Name              string /* Desired name in VPP */
	HostInterfaceName string /* Name of the host interface */
	Tag               string /* Can be configured */

	// TODO: one of the available types that CreateInterface() takes
	// CreateInterface() will wrap for example createAVF etc.
	// registerAPIHandler in every generated file that will register
	// type during init() go method
	// store will contain the registered types - CreateInterface()
	// will take opaque data with specific interface extra parameters
	// and use the type to lookup in the store
	// Type should be string
	Type string

	NumRxQueues int
	NumTxQueues int
	TxQueueSize int
	RxQueueSize int
	// meta ?? interface some sort of data
	Mtu int

	HardwareAddr *net.HardwareAddr
}

type TapInterface struct {
	Interface
	HostNamespace  string
	Tag            string
	HostMacAddress net.HardwareAddr
	Flags          TapFlags
	HostMtu        int
}

type InterfaceDetails struct {
	SwIfIndex uint32
	IsUp      bool
	Name      string
	Tag       string
	Type      string
}

type VppInterface interface {
	CreateLoopback(hwAddr *net.HardwareAddr) (swIfIndex uint32, err error)
	DeleteLoopback(iface *types.Interface) (err error)
	CreateTapV2(tap *TapInterface) (swIfIndex uint32, err error)
	CreateOrAttachTapV2(tap *TapInterface) (swIfIndex uint32, err error)
	DelTap(iface *types.Interface) error
	SearchInterfaceWithName(name string) (err error, swIfIndex uint32)
	SearchInterfaceWithTag(tag string) (uint32, error)
	SearchInterfacesWithTagPrefix(tag string) (map[string]uint32, error)
	SetInterfaceMtu(iface *types.Interface, mtu int) error
	SetInterfaceRxMode(iface *types.Interface, queueID uint32, mode RxMode) error
	SetInterfaceMacAddress(iface *types.Interface, mac *net.HardwareAddr) error
	SetInterfaceVRF(iface *types.Interface, vrfIndex uint32, isIP6 bool) error
	SetInterfaceTxPlacement(iface *types.Interface, queue int, worker int) error
	SetInterfaceRxPlacement(iface *types.Interface, queue int, worker int, main bool) error
	GetInterfaceDetails(iface *types.Interface) (i *InterfaceDetails, err error)
	GetInterfaceNeighbors(iface *types.Interface, isIPv6 bool) (err error, neighbors []Neighbor)
	EnableInterfaceIP46(iface *types.Interface) (err error)
	DisableInterfaceIP46(iface *types.Interface) (err error)
	EnableInterfaceIP4(iface *types.Interface) error
	DisableInterfaceIP4(iface *types.Interface) error
	EnableInterfaceIP6(iface *types.Interface) error
	DisableInterfaceIP6(iface *types.Interface) error
	GetInterfaceAddressesIP4(iface *types.Interface) ([]IfAddress, error)
	GetInterfaceAddressesIP6(iface *types.Interface) ([]IfAddress, error)
	SetInterfaceVrfIP4(iface *types.Interface, vrfIndex uint32) error
	SetInterfaceVrfIP6(iface *types.Interface, vrfIndex uint32) error
	AddInterfaceAddress(iface *types.Interface, addr *net.IPNet) error
	DelInterfaceAddress(iface *types.Interface, addr *net.IPNet) error
	SetInterfaceTag(iface *types.Interface, tag string) error
	UnsetInterfaceTag(iface *types.Interface, tag string) error
	EnableGSOFeature(iface *types.Interface) error
	DisableGSOFeature(iface *types.Interface) error
	SetPromiscOn(iface *types.Interface) error
	SetPromiscOff(iface *types.Interface) error
	InterfaceAdminUp(iface *types.Interface) error
	InterfaceAdminDown(iface *types.Interface) error
	InterfaceSetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error
	InterfaceUnsetUnnumbered(unnumberedSwIfIndex uint32, swIfIndex uint32) error
}

func (i *Interface) String() string {
	return fmt.Sprintf("[%d]", i.SwIfIndex)
}

func UnformatRxMode(str string) RxMode {
	switch str {
	case "interrupt":
		return Interrupt
	case "polling":
		return Polling
	case "adaptive":
		return Adaptative
	default:
		return UnknownRxMode
	}
}

func (a RxMode) String() string {
	switch a {
	case Interrupt:
		return "interrupt"
	case Polling:
		return "polling"
	case Adaptative:
		return "adaptive"
	default:
		return "default"
	}
}
