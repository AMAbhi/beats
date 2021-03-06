// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by "stringer -type=LicenseType -linecomment=true"; DO NOT EDIT.

package licenser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OSS-0]
	_ = x[Trial-1]
	_ = x[Basic-2]
	_ = x[Standard-3]
	_ = x[Gold-4]
	_ = x[Platinum-5]
	_ = x[Enterprise-6]
}

const _LicenseType_name = "Open sourceTrialBasicStandardGoldPlatinumEnterprise"

var _LicenseType_index = [...]uint8{0, 11, 16, 21, 29, 33, 41, 51}

func (i LicenseType) String() string {
	if i < 0 || i >= LicenseType(len(_LicenseType_index)-1) {
		return "LicenseType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LicenseType_name[_LicenseType_index[i]:_LicenseType_index[i+1]]
}
