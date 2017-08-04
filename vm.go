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

// Represents a virtual machine.
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

// This operation stops any migration of a virtual machine to another physical host.
func (vm *VM) CancelMigration() (err error) {
	err = vm.DoAction("cancelmigration", struct{}{})
	return
}

func (vm *VM) Clone(async string, newVM *VM) (err error) {
	err = vm.DoAction("clone", struct {
		Async string `json:"async,omitempty"`
		VM    *VM    `json:"vm,omitempty"`
	}{
		async,
		newVM,
	})
	return
}

// Permanently restores the virtual machine to the state of the previewed snapshot.
func (vm *VM) CommitSnapshot(async string) (err error) {
	err = vm.DoAction("commitsnapshot", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// Detaches a virtual machine from a pool.
func (vm *VM) Detach() (err error) {
	err = vm.DoAction("detach", struct{}{})
	return
}

// Exports a virtual machine to an export domain.
// func (vm *VM) Export(async string, discardSnapshots, exclusive string, storageDomain *StorageDomain) (err error) {
// 	err = vm.DoAction("export", struct{
// 		Async string `json:"async,omitempty"`
// 		DiscardSnapshots string `json:"discard_snapshots,omitempty"`
// 		Exclusive string `json:"exclusive,omitempty"`
// 		StorageDomain *StorageDomain `json:"storage_domain,omitempty"`
// 	}{
// 		async,
// 		discardSnapshots,
// 		exclusive,
// 		storageDomain,
// 	})
// 	return
// }

// Freezes virtual machine file systems.
func (vm *VM) FreezeFilesystems(async string) (err error) {
	err = vm.DoAction("freezefilesystems", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// Initiates the automatic user logon to access a virtual machine from an external console.
func (vm *VM) Logon(async string) (err error) {
	err = vm.DoAction("logon", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// Sets the global maintenance mode on the hosted engine virtual machine.
func (vm *VM) Maintenance(async, maintenanceEnabled string) (err error) {
	err = vm.DoAction("maintenance", struct {
		Async              string `json:"async,omitempty"`
		MaintenanceEnabled string `json:"maintenance_enabled,omitempty"`
	}{
		async,
		maintenanceEnabled,
	})
	return
}

// Migrates a virtual machine to another physical host.
func (vm *VM) Migrate(async string, cluster *Cluster, force string, host *Link) (err error) {
	err = vm.DoAction("migrate", struct {
		Async   string   `json:"async,omitempty"`
		Cluster *Cluster `json:"cluster,omitempty"`
		Force   string   `json:"force,omitempty"`
		Host    *Link    `json:"Link,omitempty"`
	}{
		async,
		cluster,
		force,
		host,
	})
	return
}

// Sends a reboot request to a virtual machine.
func (vm *VM) Reboot(async string) (err error) {
	err = vm.DoAction("reboot", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

func (vm *VM) ReorderMACAddresses(async string) (err error) {
	err = vm.DoAction("reordermacaddresses", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// This operation sends a shutdown request to a virtual machine.
func (vm *VM) Shutdown(async string) (err error) {
	err = vm.DoAction("shutdown", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// Starts the virtual machine.
func (vm *VM) Start(async, filter, pause, useCloudInit, useSysprep string, nextBootVM *VM) (err error) {
	err = vm.DoAction("start", struct {
		Async        string `json:"async,omitempty"`
		Filter       string `json:"filter,omitempty"`
		Pause        string `json:"pause,omitempty"`
		UseCloudInit string `json:"use_cloud_init,omitempty"`
		UseSysPrep   string `json:"use_sys_prep,omitempty"`
		VM           *VM    `json:"vm,omitempty"`
	}{
		async,
		filter,
		pause,
		useCloudInit,
		useSysprep,
		nextBootVM,
	})
	return
}

// This operation forces a virtual machine to power-off.
func (vm *VM) Stop(async string) (err error) {
	err = vm.DoAction("stop", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// This operation saves the virtual machine state to disk and stops it.
func (vm *VM) Suspend(async string) (err error) {
	err = vm.DoAction("suspend", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// Thaws virtual machine file systems.
func (vm *VM) ThawFilesystems(async string) (err error) {
	err = vm.DoAction("thawfilesystems", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}

// Restores the virtual machine to the state it had before previewing the snapshot.
func (vm *VM) UndoSnapshot(async string) (err error) {
	err = vm.DoAction("undosnapshot", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
	return
}
