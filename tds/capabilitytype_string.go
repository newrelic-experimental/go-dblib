// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=CapabilityType"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CapabilityRequest-1]
	_ = x[CapabilityResponse-2]
	_ = x[CapabilitySecurity-3]
}

const _CapabilityType_name = "CapabilityRequestCapabilityResponseCapabilitySecurity"

var _CapabilityType_index = [...]uint8{0, 17, 35, 53}

func (i CapabilityType) String() string {
	i -= 1
	if i >= CapabilityType(len(_CapabilityType_index)-1) {
		return "CapabilityType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CapabilityType_name[_CapabilityType_index[i]:_CapabilityType_index[i+1]]
}
