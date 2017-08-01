package ovirtapi

type Cluster struct {
	OvirtObject
	BallooningEnabled string `json:"ballooning_enabled,omitempty"`
	Cpu               struct {
		Architecture string `json:"architecure,omitempty"`
		Type         string `json:"type,omitempty"`
	} `json:"cpu,omitempty"`
	CustomSchedulingPolicyProperties *struct {
		Property []struct {
			Name  string `json:"name,omitempty"`
			Value int    `json:"value,omitempty"`
		} `json:"property,omitempty"`
	} `json:"CustomSchedulingPolicyProperties"`
	ErrorHandling *struct {
		OnError string `json:"on_error,omitempty"`
	} `json:"ErrorHandling,omitempty"`
	FencingPolicy *struct {
		Enabled                  string `json:"enabled,omitempty"`
		SkipIfConnectivitybroken *struct {
			Enabled   string `json:"enabled,omitempty"`
			Threshold int    `json:"threshold,omitempty"`
		} `json:"skip_if_connectivity_broken,omitempty"`
		SkipIfGlusterBricksUp     string `json:"skip_if_gluster_bricks_up,omitempty"`
		SkipIfGlusterQuorumNotMet string `json:"skip_if_gluster_quorum_not_met,omitempty"`
		SkipIfSdActive            *struct {
			Enabled string `json:"enabled,omitempty"`
		} `json:"skip_if_sd_active,omitempty"`
	} `json:"fencing_policy,omitempty"`
	GlusterService string `json:"gluster_service,omitempty"`
	HaReservation  string `json:"ha_reservation,omitempty"`
	Ksm            *struct {
		Enabled          string `json:"enabled,omitempty"`
		MergeAcrossNodes string `json:"merge_across_nodes,omitempty"`
	} `json:"ksm,omitempty"`
	MaintenanceReasonRequired string `json:"maintenance_reason_required,omitempty"`
	MemoryPolicy              *struct {
		OverCommit *struct {
			Percent int `json:"percent,omitempty"`
		} `json:"over_commit,omitempty"`
		TransparentHugepages *struct {
			Enabled string `json:"enabled,omitempty"`
		} `json:"transparent_hugepages,omitempty"`
	} `json:"memory_policy,omitempty"`
	Migration *struct {
		AutoConverge string `json:"auto_converge,omitempty"`
		Bandwidth    *struct {
			AssignmentMethod string `json:"assignment_method,omitempty"`
		} `json:"bandwidth,omitempty"`
		Compressed string `json:"compressed,omitempty"`
	} `json:"migration,omitempty"`
	OptionalReason     string `json:"optional_reason,omitempty"`
	RequiredRngSources *struct {
		RequiredRngSource string `json:"required_rng_source,omitempty"`
	} `json:"required_rng_sources,omitempty"`
	SwitchType      string `json:"switch_type,omitempty"`
	ThreadsAsCores  string `json:"threads_as_cores,omitempty"`
	TrustedService  string `json:"trusted_service,omitempty"`
	TunnelMigration string `json:"tunnel_migration,omitempty"`
	Version         *struct {
		Major int `json:"major,omitempty"`
		Minor int `json:"minor,omitempty"`
	} `json:"version,omitempty"`
	VirtService      string `json:"virt_service,omitempty"`
	DataCenter       *Link  `json:"data_center,omitempty"`
	MacPool          *Link  `json:"mac_pool,omitempty"`
	SchedulingPolicy *Link  `json:"scheduling_policy,omitempty"`
}
