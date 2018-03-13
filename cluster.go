// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi

type GlusterClient struct {
	BytesRead    int    `json:"bytes_read,omitempty,string"`
	BytesWritten int    `json:"bytes_written,omitempty,string"`
	ClientPort   int    `json:"client_port,omitempty,string"`
	HostName     string `json:"host_name,omitempty"`
}

type GlusterMemoryPool struct {
	Alloc_count int `json:"alloc_count,omitempty,string"`
	Cold_count  int `json:"cold_count,omitempty,string"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	Hot_count   int    `json:"hot_count,omitempty,string"`
	// A unique identifier.
	Id           string `json:"id,omitempty"`
	Max_alloc    int    `json:"max_alloc,omitempty,string"`
	Max_stdalloc int    `json:"max_stdalloc,omitempty,string"`
	// A human-readable name in plain text.
	Name        string `json:"name,omitempty"`
	Padded_size int    `json:"padded_size,omitempty,string"`
	Pool_misses int    `json:"pool_misses,omitempty,string"`
	Type        string `json:"type,omitempty"`
}

type GlusterBrick struct {
	BrickDir string `json:"brick_dir,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description    string          `json:"description,omitempty"`
	Device         string          `json:"device,omitempty"`
	FSName         string          `json:"fs_name,omitempty"`
	GlusterClients []GlusterClient `json:"gluster_clients,omitempty"`
	// A unique identifier.
	Id          string              `json:"id,omitempty"`
	MemoryPools []GlusterMemoryPool `json:"memory_pools,omitempty"`
	MntOptions  string              `json:"mnt_options,omitempty"`
	// A human-readable name in plain text.
	Name     string `json:"name,omitempty"`
	Pid      int    `json:"pid,omitempty"`
	Port     int    `json:"port,omitempty"`
	ServerId string `json:"server_id,omitempty"`
	Status   string `json:"status,omitempty"`
}

type GlusterVolume struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description   string `json:"description,omitempty"`
	DisperseCount int    `json:"disperse_count,omitempty,string"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name            string         `json:"name,omitempty"`
	Options         []Option       `json:"options,omitempty"`
	RedundancyCount int            `json:"redundancy_count,omitempty,string"`
	ReplicaCount    int            `json:"replica_count,omitempty,string"`
	Status          string         `json:"status,omitempty"`
	StripeCount     int            `json:"stripe_count,omitempty,string"`
	TransportTypes  []string       `json:"transport_types,omitempty"`
	VolumeType      string         `json:"volume_type,omitempty"`
	Bricks          []GlusterBrick `json:"bricks,omitempty"`
	Cluster         Cluster        `json:"cluster,omitempty"`
	// statistics      []Statistic    `json:"statistics,omitempty"`
}

type Property struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type ErrorHandling struct {
	OnError string `json:"on_error,omitempty"`
}

type SkipIfConnectivityBroken struct {
	// If enabled, we will not fence a host in case more than a configurable percentage of hosts in the cluster lost connectivity as well.
	Enabled string `json:"enabled,omitempty"`
	// Threshold for connectivity testing.
	Threshold int `json:"threshold,omitempty,string"`
}

type SkipIfSDActive struct {
	// If enabled, we will skip fencing in case the host maintains its lease in the storage.
	Enabled string `json:"enabled,omitempty"`
}

// FencingPolicy Type representing a cluster fencing policy.
type FencingPolicy struct {
	// Enable or disable fencing on this cluster.
	Enabled string `json:"enabled,omitempty"`
	// If enabled, we will not fence a host in case more than a configurable percentage of hosts in the cluster lost connectivity as well.
	SkipIfConnectivityBroken *SkipIfConnectivityBroken `json:"skip_if_connectivity_broken,omitempty"`
	// A flag indicating if fencing should be skipped if Gluster bricks are up and running in the host being fenced.
	SkipIfGlusterBricksUp string `json:"skip_if_gluster_bricks_up,omitempty"`
	// A flag indicating if fencing should be skipped if Gluster bricks are up and running and Gluster quorum will not be met without those bricks.
	SkipIfGlusterQuorumNotMet string `json:"skip_if_gluster_quorum_not_met,omitempty"`
	// If enabled, we will skip fencing in case the host maintains its lease in the storage.
	SkipIfSdActive *SkipIfSDActive `json:"skip_if_sd_active,omitempty"`
}

type SerialNumber struct {
	Policy string `json:"policy,omitempty"`
	Value  string `json:"value,omitempty"`
}

type Properties struct {
	Property []Property `json:"property,omitempty"`
}

type RequiredRNGSources struct {
	RequiredRNGSource []string `json:"required_rng_source,omitempty"`
}

// Cluster Type representation of a cluster.
type Cluster struct {
	OvirtObject
	BallooningEnabled string `json:"ballooning_enabled,omitempty"`
	Comment           string `json:"comment,omitempty"`
	CPU               *CPU   `json:"cpu,omitempty"`
	// Custom scheduling policy properties of the cluster.
	CustomSchedulingPolicyProperties *Properties    `json:"custom_scheduling_policy_properties,omitempty"`
	ErrorHandling                    *ErrorHandling `json:"ErrorHandling,omitempty"`
	// Custom fencing policy can be defined for a cluster.
	FencingPolicy  *FencingPolicy `json:"fencing_policy,omitempty"`
	GlusterService string         `json:"gluster_service,omitempty"`
	// The name of the https://fedorahosted.
	GlusterTunedProfile       string            `json:"gluster_tuned_profile,omitempty"`
	HAReservation             string            `json:"ha_reservation,omitempty"`
	KSM                       *KSM              `json:"ksm,omitempty"`
	MaintenanceReasonRequired string            `json:"maintenance_reason_required,omitempty"`
	MemoryPolicy              *MemoryPolicy     `json:"memory_policy,omitempty"`
	Migration                 *MigrationOptions `json:"migration,omitempty"`
	OptionalReason            string            `json:"optional_reason,omitempty"`
	// Set of random number generator (RNG) sources required from each host in the cluster.
	RequiredRNGSources *RequiredRNGSources `json:"required_rng_sources,omitempty"`
	SerialNumber       *SerialNumber       `json:"serial_number,omitempty"`
	SupportedVersions  []Version           `json:"supported_versions,omitempty"`
	// Type of switch to be used by all networks in given cluster.
	SwitchType      string `json:"switch_type,omitempty"`
	ThreadsAsCores  string `json:"threads_as_cores,omitempty"`
	TrustedService  string `json:"trusted_service,omitempty"`
	TunnelMigration string `json:"tunnel_migration,omitempty"`
	// The compatibility version of the cluster.
	Version           *Version        `json:"version,omitempty"`
	VirtService       string          `json:"virt_service,omitempty"`
	AffinityGroups    []Link          `json:"affinity_groups,omitempty"`
	CPUProfiles       []Link          `json:"cpu_profiles,omitempty"`
	DataCenter        *DataCenter     `json:"data_center,omitempty"`
	GlusterHooks      []Link          `json:"gluster_hooks,omitempty"`
	GlusterVolumes    []GlusterVolume `json:"gluster_volume,omitempty"`
	MacPool           *Link           `json:"mac_pool,omitempty"`
	ManagementNetwork *Link           `json:"management_network,omitempty"`
	NetworkFilters    *Link           `json:"network_filters,omitempty"`
	Networks          []Link          `json:"networks,omitempty"`
	Permissions       []Link          `json:"permissions,omitempty"`
	SchedulingPolicy  *Link           `json:"scheduling_policy,omitempty"`
}
