// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

import (
	"bytes"
	"fmt"

	"github.com/snapcore/snapd/interfaces"
)

const plasmaConnectedPlugAppArmor = `
# Description: Can query UPower for power devices, history and statistics.

#include <abstractions/dbus-strict>

# Find all devices monitored by UPower
dbus (send)
    bus=system
    path=/org/freedesktop/UPower
    interface=org.freedesktop.UPower
    member=EnumerateDevices
    peer=(label=###SLOT_SECURITY_TAGS###),
`

// PlasmaInterface is the hello interface for a tutorial.
type PlasmaInterface struct{}

// String returns the same value as Name().
func (iface *PlasmaInterface) Name() string {
	return "plasma"
}

// SanitizeSlot checks and possibly modifies a slot.
func (iface *PlasmaInterface) SanitizeSlot(slot *interfaces.Slot) error {
	if iface.Name() != slot.Interface {
		panic(fmt.Sprintf("slot is not of interface %q", iface))
	}
	// NOTE: currently we don't check anything on the slot side.
	return nil
}

// SanitizePlug checks and possibly modifies a plug.
func (iface *PlasmaInterface) SanitizePlug(plug *interfaces.Plug) error {
	if iface.Name() != plug.Interface {
		panic(fmt.Sprintf("plug is not of interface %q", iface))
	}
	// NOTE: currently we don't check anything on the plug side.
	return nil
}

// ConnectedSlotSnippet returns security snippet specific to a given connection between the hello slot and some plug.
func (iface *PlasmaInterface) ConnectedSlotSnippet(plug *interfaces.Plug, slot *interfaces.Slot, securitySystem interfaces.SecuritySystem) ([]byte, error) {
	switch securitySystem {
	case interfaces.SecurityAppArmor:
		return nil, nil
	case interfaces.SecuritySecComp:
		return nil, nil
	}
	return nil, nil
}

// PermanentSlotSnippet returns security snippet permanently granted to hello slots.
func (iface *PlasmaInterface) PermanentSlotSnippet(slot *interfaces.Slot, securitySystem interfaces.SecuritySystem) ([]byte, error) {
	switch securitySystem {
	case interfaces.SecurityAppArmor:
		return nil, nil
	case interfaces.SecuritySecComp:
		return nil, nil
	}
	return nil, nil
}

// ConnectedPlugSnippet returns security snippet specific to a given connection between the hello plug and some slot.
func (iface *PlasmaInterface) ConnectedPlugSnippet(plug *interfaces.Plug, slot *interfaces.Slot, securitySystem interfaces.SecuritySystem) ([]byte, error) {
	fmt.Println("plasma got ConnectedPlugSnippet")
	switch securitySystem {
	case interfaces.SecurityAppArmor:
		old := []byte("###SLOT_SECURITY_TAGS###")
		new := slotAppLabelExpr(slot)
		snippet := bytes.Replace([]byte(plasmaConnectedPlugAppArmor), old, new, -1)
		return snippet, nil
	}
	return nil, nil
}

// PermanentPlugSnippet returns the configuration snippet required to use a hello interface.
func (iface *PlasmaInterface) PermanentPlugSnippet(plug *interfaces.Plug, securitySystem interfaces.SecuritySystem) ([]byte, error) {
	switch securitySystem {
	case interfaces.SecurityAppArmor:
		return nil, nil
	case interfaces.SecuritySecComp:
		return nil, nil
	}
	return nil, nil
}

// AutoConnect returns true if plugs and slots should be implicitly
// auto-connected when an unambiguous connection candidate is available.
//
// This interface does not auto-connect.
func (iface *PlasmaInterface) AutoConnect(*interfaces.Plug, *interfaces.Slot) bool {
	return false
}

func init() {
	allInterfaces = append(allInterfaces, &PlasmaInterface{})
}
