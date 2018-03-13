// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi

type DataCenter struct {
	OvirtObject
	Local             string `json:"local,omitempty"`
	QuotaMode         string `json:"quota_mode,omitempty"`
	Status            string `json:"status,omitempty"`
	StorageFormat     string `json:"storage_format,omitempty"`
	SupportedVersions *struct {
		Version []struct {
			Major string `json:"major,omitempty"`
			Minor string `json:"minor,omitempty"`
		} `json:"version,omitempty"`
	} `json:"supported_versions,omitempty"`
	Version *struct {
		Major string `json:"major,omitempty"`
		Minor string `json:"minor,omitempty"`
	} `json:"version,omitempty"`
}
