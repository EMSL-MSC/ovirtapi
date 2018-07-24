// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi

type TemplateVersion struct {
	VersionName   string `json:"version_name,omitempty"`
	VersionNumber string `json:"version_number,omitempty"`
}

type Template struct {
	OvirtObject
	Comment                    string             `json:"comment,omitempty"`
	Bios                       *Bios              `json:"bios,omitempty"`
	CPU                        *CPU               `json:"cpu,omitempty"`
	CPUShares                  string             `json:"cpu_shares,omitempty"`
	CreationTime               int                `json:"creation_time,omitempty"`
	Display                    *Display           `json:"display,omitempty"`
	HighAvailability           *HighAvailability  `json:"high_availability,omitempty"`
	LargeIcon                  *Link              `json:"large_icon,omitempty"`
	Memory                     int                `json:"memory,string,omitempty"`
	MemoryPolicy               *MemoryPolicy      `json:"memory_policy,omitempty"`
	Migration                  *MigrationOptions  `json:"migration,omitempty"`
	MigrationDowntime          string             `json:"migration_downtime,omitempty"`
	Origin                     string             `json:"origin,omitempty"`
	Os                         *OperatingSystem   `json:"os,omitempty"`
	SmallIcon                  *Link              `json:"small_icon,omitempty"`
	StartPaused                string             `json:"start_paused,omitempty"`
	Stateless                  string             `json:"stateless,omitempty"`
	TimeZone                   *TimeZone          `json:"time_zone,omitempty"`
	Type                       string             `json:"type,omitempty"`
	USB                        *USB               `json:"usb,omitempty"`
	Cluster                    *Link              `json:"cluster,omitempty"`
	CPUProfile                 *Link              `json:"cpu_profile,omitempty"`
	Quota                      *Link              `json:"quota,omitempty"`
	NextRunConfigurationExists string             `json:"next_run_configuration_exists,omitempty"`
	NumaTuneMode               string             `json:"numa_tune_mode,omitempty"`
	PlacementPolicy            *VMPlacementPolicy `json:"placement_policy,omitempty"`
	RunOnce                    string             `json:"run_once,omitempty"`
	StartTime                  int                `json:"start_time,omitempty"`
	StopTime                   int                `json:"stop_time,omitempty"`
	Status                     string             `json:"status,omitempty"`
	Host                       *Link              `json:"host,omitempty"`
	InstanceType               *Link              `json:"instance_type,omitempty"`
	OriginalTemplate           *Link              `json:"original_template,omitempty"`
	Template                   *Link              `json:"template,omitempty"`
	Version                    *TemplateVersion   `json:"version,omitempty"`
	VM                         *VM                `json:"vm,omitempty"`
}
