package ovirtapi

type Bios struct {
	BootMenu struct {
		Enabled string `json:"enabled"`
	} `json:"boot_menu"`
}

type Core struct {
	Index  int `json:index,omitempty"`
	Socket int `json:socket,omitempty"`
}

type VCPUPin struct {
	CPUSet string `json:"cpu_set,omitempty"`
	VCPU   int    `json:"vcpu,omitempty"`
}

type CPUTune struct {
	VCPUPins []VCPUPin `json:"vcpu_pins,omitempty"`
}
type CPUTopology struct {
	Cores   string `json:"cores,omitempty"`
	Sockets string `json:"sockets,omitempty"`
	Threads string `json:"threads,omitempty"`
}

type CPU struct {
	Architecture string       `json:"architecture,omitempty"`
	Cores        []Core       `json:"cores,omitempty"`
	CPUTune      *CPUTune     `json:"cpu_tune,omitempty"`
	Level        int          `json:"level,omitempty"`
	CPUMode      string       `json:"cpu_mode,omitempty"`
	Name         string       `json:"name,omitempty"`
	Speed        string       `json:"speed,omitempty"`
	Topology     *CPUTopology `json:"topology,omitempty"`
	Type         string       `json:"type,omitempty"`
}

type Certificate struct {
	Comment      string `json:"comment,omitempty"`
	Content      string `json:"content,omitempty"`
	Description  string `json:"description,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Subject      string `json:"subject,omitempty"`
}

type Display struct {
	Address             string       `json:"address,omitempty"`
	AllowOverride       string       `json:"allow_override,omitempty"`
	Certificate         *Certificate `json:"certificate,omitempty"`
	CopyPasteEnabled    string       `json:"copy_paste_enabled,omitempty"`
	DisconnectAction    string       `json:"disconnect_action,omitempty"`
	FileTransferEnabled string       `json:"file_transfer_enabled,omitempty"`
	Monitors            string       `json:"monitors,omitempty"`
	SecurePort          string       `json:"secure_port,omitempty"`
	SingleQxlPci        string       `json:"single_qxl_pci,omitempty"`
	SmartcardEnabled    string       `json:"smartcard_enabled,omitempty"`
	Type                string       `json:"type,omitempty"`
}

type HighAvailability struct {
	Enabled  string `json:"enabled,omitempty"`
	Priority string `json:"priority,omitempty"`
}

type IO struct {
	Threads string `json:"threads,omitempty"`
}

type MemoryPolicy struct {
	Ballooning string `json:"ballooning,omitempty"`
	Guaranteed int    `json:"guaranteed,omitempty"`
	Max        int    `json:"max,omitempty"`
}

type Migration struct {
	AutoConverge string `json:"auto_converge,omitempty"`
	Compressed   string `json:"compressed,omitempty"`
}

type OperatingSystem struct {
	Boot                  *Boot    `json:"boot,omitempty"`
	Cmdline               string   `json:"cmdline,omitempty"`
	CustomKernelCmdline   string   `json:"custom_kernel_cmdline,omitempty"`
	Initrd                string   `json:"initrd,omitempty"`
	Kernel                string   `json:"kernel,omitempty"`
	ReportedKernelCmdline string   `json:"reported_kernel_cmdline,omitempty"`
	Type                  string   `json:"type,omitempty"`
	Version               *Version `json:"version,omitempty"`
}

type Boot struct {
	Devices []string `json:"devices>device,omitempty"`
}

type Version struct {
	Build       int    `json:"build,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Description string `json:"description,omitempty"`
	FullVersion string `json:"full_version,omitempty"`
	Id          string `json:"id,omitempty"`
	Major       int    `json:"major,omitempty"`
	Minor       int    `json:"minor,omitempty"`
	Name        string `json:"name,omitempty"`
	Revision    int    `json:"revision,omitempty"`
}

type TimeZone struct {
	Name      string `json:"name,omitempty"`
	UTCOffset string `json:"utc_offset,omitempty"`
}

type USB struct {
	Enabled string `json:"enabled,omitempty"`
	Type    string `json:"type,omitempty"`
}

type VMPlacementPolicy struct {
	Affinity string `json:"affinity,omitempty"`
	//Host
}

type VM struct {
	OvirtObject
	Comment                    string             `json:"comment,omitempty"`
	Bios                       *Bios              `json:"bios,omitempty"`
	CPU                        *CPU               `json:"cpu,omitempty"`
	CpuShares                  string             `json:"cpu_shares,omitempty"`
	CreationTime               int                `json:"creation_time,omitempty"`
	Display                    *Display           `json:"display,omitempty"`
	HighAvailability           *HighAvailability  `json:"high_availability,omitempty"`
	Io                         *IO                `json:"io,omitempty"`
	LargeIcon                  *Link              `json:"large_icon,omitempty"`
	Memory                     int                `json:"memory,omitempty"`
	MemoryPolicy               *MemoryPolicy      `json:"memory_policy,omitempty"`
	Migration                  *Migration         `json:"migration,omitempty"`
	MigrationDowntime          string             `json:"migration_downtime,omitempty"`
	Origin                     string             `json:"origin,omitempty"`
	Os                         *OperatingSystem   `json:"os,omitempty"`
	SmallIcon                  *Link              `json:"small_icon,omitempty"`
	StartPaused                string             `json:"start_paused,omitempty"`
	Stateless                  string             `json:"stateless,omitempty"`
	TimeZone                   *TimeZone          `json:"time_zone,omitempty"`
	Type                       string             `json:"type,omitempty"`
	USB                        *USB               `json:"usb,omitempty"`
	Cluster                    *Cluster           `json:"cluster,omitempty"`
	CpuProfile                 *Link              `json:"cpu_profile,omitempty"`
	Quota                      *Link              `json:"quota,omitempty"`
	NextRunConfigurationExists string             `json:"next_run_configuration_exists,omitempty"`
	NumaTuneMode               string             `json:"numa_tune_mode,omitempty"`
	PlacementPolicy            *VMPlacementPolicy `json:"placement_policy,omitempty"`
	Runonce                    string             `json:"run_once,omitempty"`
	Starttime                  int                `json:"start_time,omitempty"`
	StopTime                   int                `json:"stop_time,omitempty"`
	Status                     string             `json:"status,omitempty"`
	Host                       *Link              `json:"host,omitempty"`
	InstanceType               *Link              `json:"instance_type,omitempty"`
	OriginalTemplate           *Link              `json:"original_template,omitempty"`
	Template                   *Template          `json:"template,omitempty"`
}
