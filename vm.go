package ovirtapi

import (
	"encoding/xml"
)

type Vm struct {
	OvirtObject
	XMLName     xml.Name `xml:"vm"`
	Description string   `xml:"description,omitempty"`
	Comment     string   `xml:"comment,omitempty"`
	Bios        *struct {
		BootMenu struct {
			Enabled bool `xml:"enabled"`
		} `xml:"boot_menu"`
	} `xml:"bios,omitempty"`
	Cpu *struct {
		Architecture string `xml:"architecture,omitempty"`
		Topology     struct {
			Cores   int `xml:"cores,omitempty"`
			Sockets int `xml:"sockets,omitempty"`
			Threads int `xml:"threads,omitempty"`
		} `xml:"topology,omitempty"`
	} `xml:"cpu,omitempty"`
	CpuShares    int    `xml:"cpu_shares,omitempty"`
	CreationTime string `xml:"creation_time,omitempty"`
	Display      *struct {
		Address       string `xml:"address,omitempty"`
		AllowOverride bool   `xml:"allow_override,omitempty"`
		Certificate   struct {
			Subject string `xml:"subject,omitempty"`
		} `xml:"certificate,omitempty"`
		CopyPasteEnabled    bool   `xml:"copy_paste_enabled,omitempty"`
		DisconnectAction    string `xml:"disconnect_action,omitempty"`
		FileTransferEnabled bool   `xml:"file_transfer_enabled,omitempty"`
		Monitors            int    `xml:"monitors,omitempty"`
		SecurePort          int    `xml:"secure_port,omitempty"`
		SingleQxlPci        bool   `xml:"single_qxl_pci,omitempty"`
		SmartcardEnabled    bool   `xml:"smartcard_enabled,omitempty"`
		Type                string `xml:"type,omitempty"`
	} `xml:"display,omitempty"`
	HighAvailability *struct {
		Enabled  bool `xml:"enabled,omitempty"`
		Priority int  `xml:"priority,omitempty"`
	} `xml:"high_availability,omitempty"`
	// Io *struct {
	// 	Threads int `xml:"threads,omitempty"`
	// } `xml:"io,omitempty"`
	LargeIcon    *Link `xml:"large_icon,omitempty"`
	Memory       int            `xml:"memory,omitempty"`
	MemoryPolicy *struct {
		Ballooning bool `xml:"ballooning,omitempty"`
		Guaranteed int  `xml:"guaranteed,omitempty"`
		Max        int  `xml:"max,omitempty"`
	} `xml:"memory_policy,omitempty"`
	Migration *struct {
		AutoConverge string `xml:"auto_converge,omitempty"`
		Compressed   string `xml:"compressed,omitempty"`
	} `xml:"migration,omitempty"`
	MigrationDowntime string `xml:"migration_downtime,omitempty"`
	Origin            string `xml:"origin,omitempty"`
	Os                *struct {
		Boot struct {
			Devices []string `xml:"devices>device,omitempty"`
		} `xml:"boot,omitempty"`
		Type string `xml:"type,omitempty"`
	} `xml:"os,omitempty"`
	SmallIcon   *Link `xml:"small_icon,omitempty"`
	StartPaused bool           `xml:"start_paused,omitempty"`
	Stateless   bool           `xml:"stateless,omitempty"`
	TimeZone    *struct {
		Name string `xml:"name,omitempty"`
	} `xml:"time_zone,omitempty"`
	Type string `xml:"type,omitempty"`
	Usb  *struct {
		Enabled bool `xml:"enabled,omitempty"`
	} `xml:"usb,omitempty"`
	Cluster                    *OvirtObject `xml:"cluster,omitempty"`
	CpuProfile                 *Link `xml:"cpu_profile,omitempty"`
	Quota                      *Link `xml:"quota,omitempty"`
	NextRunConfigurationExists bool           `xml:"next_run_configuration_exists,omitempty"`
	NumaTuneMode               string         `xml:"numa_tune_mode,omitempty"`
	PlacementPolicy            *struct {
		Affinity string `xml:"affinity,omitempty"`
	} `xml:"placement_policy,omitempty"`
	RunOnce          bool           `xml:"run_once,omitempty"`
	StartTime        string         `xml:"start_time,omitempty"`
	Status           string         `xml:"status,omitempty"`
	StopTime         string         `xml:"stop_time,omitempty"`
	Host             *Link `xml:"host,omitempty"`
	InstanceType     *Link `xml:"instance_type,omitempty"`
	OriginalTemplate *Link `xml:"original_template,omitempty"`
	Template         *Link `xml:"template,omitempty"`
}
