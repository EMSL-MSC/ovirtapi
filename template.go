package ovirtapi

type Template struct {
	OvirtObject
	Comment string `json:"comment,omitempty"`
	Bios    *struct {
		BootMenu struct {
			Enabled string `json:"enabled"`
		} `json:"boot_menu"`
	} `json:"bios,omitempty"`
	Cpu *struct {
		Architecture string `json:"architecture,omitempty"`
		Topology     struct {
			Cores   string `json:"cores,omitempty"`
			Sockets string `json:"sockets,omitempty"`
			Threads string `json:"threads,omitempty"`
		} `json:"topology,omitempty"`
	} `json:"cpu,omitempty"`
	CpuShares    string `json:"cpu_shares,omitempty"`
	CreationTime int    `json:"creation_time,omitempty"`
	Display      *struct {
		Address       string `json:"address,omitempty"`
		AllowOverride string `json:"allow_override,omitempty"`
		Certificate   struct {
			Subject string `json:"subject,omitempty"`
		} `json:"certificate,omitempty"`
		CopyPasteEnabled    string `json:"copy_paste_enabled,omitempty"`
		DisconnectAction    string `json:"disconnect_action,omitempty"`
		FileTransferEnabled string `json:"file_transfer_enabled,omitempty"`
		Monitors            string `json:"monitors,omitempty"`
		SecurePort          string `json:"secure_port,omitempty"`
		SingleQxlPci        string `json:"single_qxl_pci,omitempty"`
		SmartcardEnabled    string `json:"smartcard_enabled,omitempty"`
		Type                string `json:"type,omitempty"`
	} `json:"display,omitempty"`
	HighAvailability *struct {
		Enabled  string `json:"enabled,omitempty"`
		Priority string `json:"priority,omitempty"`
	} `json:"high_availability,omitempty"`
	LargeIcon    *Link `json:"large_icon,omitempty"`
	Memory       int   `json:"memory,omitempty"`
	MemoryPolicy *struct {
		Ballooning string `json:"ballooning,omitempty"`
		Guaranteed int    `json:"guaranteed,omitempty"`
		Max        int    `json:"max,omitempty"`
	} `json:"memory_policy,omitempty"`
	Migration *struct {
		AutoConverge string `json:"auto_converge,omitempty"`
		Compressed   string `json:"compressed,omitempty"`
	} `json:"migration,omitempty"`
	MigrationDowntime string `json:"migration_downtime,omitempty"`
	Origin            string `json:"origin,omitempty"`
	Os                *struct {
		Boot struct {
			Devices []string `json:"devices>device,omitempty"`
		} `json:"boot,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"os,omitempty"`
	SmallIcon   *Link  `json:"small_icon,omitempty"`
	StartPaused string `json:"start_paused,omitempty"`
	Stateless   string `json:"stateless,omitempty"`
	TimeZone    *struct {
		Name string `json:"name,omitempty"`
	} `json:"time_zone,omitempty"`
	Type string `json:"type,omitempty"`
	Usb  *struct {
		Enabled string `json:"enabled,omitempty"`
	} `json:"usb,omitempty"`
	Cluster                    *Link  `json:"cluster,omitempty"`
	CpuProfile                 *Link  `json:"cpu_profile,omitempty"`
	Quota                      *Link  `json:"quota,omitempty"`
	NextRunConfigurationExists string `json:"next_run_configuration_exists,omitempty"`
	NumaTuneMode               string `json:"numa_tune_mode,omitempty"`
	PlacementPolicy            *struct {
		Affinity string `json:"affinity,omitempty"`
	} `json:"placement_policy,omitempty"`
	RunOnce          string `json:"run_once,omitempty"`
	StartTime        int    `json:"start_time,omitempty"`
	StopTime         int    `json:"stop_time,omitempty"`
	Status           string `json:"status,omitempty"`
	Host             *Link  `json:"host,omitempty"`
	InstanceType     *Link  `json:"instance_type,omitempty"`
	OriginalTemplate *Link  `json:"original_template,omitempty"`
	Template         *Link  `json:"template,omitempty"`
	Version          *struct {
		VersionName   string `json:"version_name,omitempty"`
		VersionNumber string `json:"version_number,omitempty"`
		BaseTemplate  *Link  `json:"base_template,omitempty"`
	} `json:"version,omitempty"`
	VM *Vm `json:"vm,omitempty"`
}
