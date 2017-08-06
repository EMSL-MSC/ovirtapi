package ovirtapi

import (
	"encoding/json"
	"fmt"
)

// TransparentHugePages Type representing a transparent huge pages (THP) support
type TransparentHugePages struct {
	Enabled string `json:"enabled,omitempty"`
}

// VMSummary Type containing information related to virtual machines on a particular host.
type VMSummary struct {
	// The number of virtual machines active on the host.
	Active int `json:"active,omitempty"`
	// The number of virtual machines migrating to or from the host.
	Migrating int `json:"migrating,omitempty"`
	// The number of virtual machines present on the host.
	Total int `json:"total,omitempty"`
}

// User Represents a user in the system.
type User struct {
	//Free text containing comments about this object.
	Comment    string `json:"comment,omitempty"`
	Department string `json:"department,omitempty"`
	//A human-readable description in plain text.
	Description   string `json:"description,omitempty"`
	DomainEntryID string `json:"domain_entry_id,omitempty"`
	Email         string `json:"email,omitempty"`
	//A unique identifier.
	ID       string `json:"id,omitempty"`
	LastName string `json:"last_name,omitempty"`
	LoggedIn string `json:"logged_in,omitempty"`
	//A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	//Namespace where the user resides.
	Namespace string `json:"namespace,omitempty"`
	Password  string `json:"password,omitempty"`
	//Similar to user_name.
	Principal string `json:"principal,omitempty"`
	//The user's username.
	UserName string `json:"user_name,omitempty"`
}

type SSH struct {
	AuthenticationMethod string `json:"authentication_method,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	Port int    `json:"port,omitempty"`
	User User   `json:"user,omitempty"`
}

type SPM struct {
	Priority int    `json:"priority,omitempty"`
	Status   string `json:"status,omitempty"`
}

// SELinux Represents SELinux in the system.
type SELinux struct {
	Mode string `json:"mode,omitempty"`
}

type PMProxy struct {
	Type string `json:"type,omitempty"`
}

type Option struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// Agent Type representing a fence agent.
type Agent struct {
	// Fence agent address.
	Address string `json:"address,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// Specifies whether the agent should be used concurrently or sequentially.
	Concurrent string `json:"concurrent,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// Specifies whether the options should be encrypted.
	EncryptOptions string `json:"encrypt_options,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	// Fence agent options (comma-delimited list of key-value pairs).
	Options []Option `json:"options,omitempty"`
	// The order of this agent if used with other agents.
	Order int `json:"order,omitempty"`
	// Fence agent password.
	Password string `json:"password,omitempty"`
	// Fence agent port.
	Port int `json:"port,omitempty"`
	// Fence agent type.
	Type string `json:"type,omitempty"`
	// Fence agent user name.
	Username string `json:"username,omitempty"`
}

type PowerManagement struct {
	// The host name or IP address of the host.
	Address string `json:"address,omitempty"`
	// Specifies fence agent options when multiple fences are used.
	Agents []Agent `json:"agents,omitempty"`
	// Toggles the automated power control of the host in order to save energy.
	AutomaticPMEnabled string `json:"automatic_pm_enabled,omitempty"`
	// Indicates whether power management configuration is enabled or disabled.
	Enabled string `json:"enabled,omitempty"`
	// Toggles whether to determine if kdump is running on the host before it is shut down.
	KdumpDetection string `json:"kdump_detection,omitempty"`
	// Fencing options for the selected type= specified with the option name="" and value="" strings.
	Options []Option `json:"options,omitempty"`
	// A valid, robust password for power management.
	Password string `json:"password,omitempty"`
	// Determines the power management proxy.
	PMProxies []PMProxy `json:"pm_proxies,omitempty"`
	// Determines the power status of the host.
	Status string `json:"status,omitempty"`
	// Fencing device code.
	Type string `json:"type,omitempty"`
	// A valid user name for power management.
	Username string `json:"username,omitempty"`
}

type KSM struct {
	Enabled          string `json:"enabled,omitempty"`
	MergeAcrossNodes string `json:"merge_across_nodes,omitempty"`
}

type HostDevicePassthrough struct {
	Enabled string `json:"enabled,omitempty"`
}

type ISCSIDetails struct {
	Address         string `json:"address,omitempty"`
	DiskID          string `json:"disk_id,omitempty"`
	Initiator       string `json:"initiator,omitempty"`
	LunMapping      int    `json:"lun_mapping,omitempty"`
	Password        string `json:"password,omitempty"`
	Paths           int    `json:"paths,omitempty"`
	Port            int    `json:"port,omitempty"`
	Portal          string `json:"portal,omitempty"`
	ProductID       string `json:"product_id,omitempty"`
	Serial          string `json:"serial,omitempty"`
	Size            int    `json:"size,omitempty"`
	Status          string `json:"status,omitempty"`
	StorageDomainID string `json:"storage_domain_id,omitempty"`
	Target          string `json:"target,omitempty"`
	Username        string `json:"username,omitempty"`
	VendorID        string `json:"vendor_id,omitempty"`
	VolumeGroupID   string `json:"volume_group_id,omitempty"`
}

type HostedEngine struct {
	Active            string `json:"active,omitempty"`
	Configured        string `json:"configured,omitempty"`
	GlobalMaintenance string `json:"global_maintenance,omitempty"`
	LocalMaintenance  string `json:"local_maintenance,omitempty"`
	Score             int    `json:"score,omitempty"`
}

// HardwareInformation Represents hardware information of host.
type HardwareInformation struct {
	// Type of host's CPU.
	Family string `json:"family,omitempty"`
	// Manufacturer of the host's machine and hardware vendor.
	Manufacturer string `json:"manufacturer,omitempty"`
	// Host's product name (for example RHEV Hypervisor).
	ProductName string `json:"product_name,omitempty"`
	// Unique ID for host's chassis.
	SerialNumber string `json:"serial_number,omitempty"`
	// Supported sources of random number generator.
	SupportedRngSources []string `json:"supported_rng_sources,omitempty"`
	// Unique ID for each host.
	Uuid string `json:"uuid,omitempty"`
	// Unique name for each of the manufacturer.
	Version string `json:"version,omitempty"`
}

// Host Type representing a host.
type Host struct {
	OvirtObject
	// The host address (FQDN/IP).
	Address string `json:"address,omitempty"`
	// The host auto non uniform memory access (NUMA) status.
	Auto_numa_status string `json:"auto_numa_status,omitempty"`
	// The host certificate.
	Certificate *Certificate `json:"certificate,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// The CPU type of this host.
	Cpu *CPU `json:"cpu,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// Specifies whether host device passthrough is enabled on this host.
	Device_passthrough *HostDevicePassthrough `json:"device_passthrough,omitempty"`
	// Optionally specify the display address of this host explicitly.
	Display *Display `json:"display,omitempty"`
	// The host external status.
	ExternalStatus string `json:"external_status,omitempty"`
	// The host hardware information.
	HardwareInformation *HardwareInformation `json:"hardware_information,omitempty"`
	// The self-hosted engine status of this host.
	Hosted_engine *HostedEngine `json:"hosted_engine,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// The host iSCSI details.
	ISCSI *ISCSIDetails `json:"iscsi,omitempty"`
	// The host KDUMP status.
	KdumpStatus string `json:"kdump_status,omitempty"`
	// Kernel SamePage Merging (KSM) reduces references to memory pages from multiple identical pages to a single page reference.
	KSM KSM `json:"ksm,omitempty"`
	// The host libvirt version.
	LibvirtVersion *Version `json:"libvirt_version,omitempty"`
	// The max scheduling memory on this host in bytes.
	MaxSchedulingMemory int `json:"max_scheduling_memory,omitempty"`
	// The amount of physical memory on this host in bytes.
	Memory int `json:"memory,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	// Specifies whether non uniform memory access (NUMA) is supported on this host.
	NumaSupported string `json:"numa_supported,omitempty"`
	// The operating system on this host.
	OS *OperatingSystem `json:"os,omitempty"`
	// Specifies whether we should override firewall definitions.
	OverrideIptables string `json:"override_iptables,omitempty"`
	// The host port.
	Port int `json:"port,omitempty"`
	// The host power management definitions.
	PowerManagement *PowerManagement `json:"power_management,omitempty"`
	// The protocol that the engine uses to communicate with the host.
	Protocol string `json:"protocol,omitempty"`
	// When creating a new host, a root password is required if the password authentication method is chosen, but this is not subsequently included in the representation.
	RootPassword string `json:"root_password,omitempty"`
	// The host SElinux status.
	SELinux *SELinux `json:"se_linux,omitempty"`
	// The host storage pool manager (SPM) status and definition.
	SPM *SPM `json:"spm,omitempty"`
	// The SSH definitions.
	SSH *SSH `json:"ssh,omitempty"`
	// The host status.
	Status string `json:"status,omitempty"`
	// The host status details.
	StatusDetail string `json:"status_detail,omitempty"`
	// The virtual machine summary - how many are active, migrating and total.
	Summary *VMSummary `json:"summary,omitempty"`
	// Transparent huge page support expands the size of memory pages beyond the standard 4 KiB limit.
	TransparentHugePages *TransparentHugePages `json:"transparent_huge_pages,omitempty"`
	// Indicates if the host contains a full installation of the operating system or a scaled-down version intended only to host virtual machines.
	Type string `json:"type,omitempty"`
	// Specifies whether there is an oVirt-related update on this host.
	UpdateAvailable string `json:"update_available,omitempty"`
	// The version of VDSM.
	Version *Version `json:"version,omitempty"`
}

// Activate the host for use, such as running virtual machines.
func (host *Host) Activate(async string) error {
	return host.DoAction("activate", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// // Approve a pre-installed Hypervisor host for usage in the virtualization environment.
// // This action also accepts an optional cluster element to define the target cluster for this host.
// func (host *Host) Approve(async string, cluster *Cluster, host *Host) error {
// 	return host.DoAction("Approve", struct {
// 		Async   string `json:"async,omitempty"`
// 		Cluster string `json:"cluster,omitempty"`
// 		Host    string `json:"host,omitempty"`
// 	}{
// 		async,
// 		cluster,
// 		host,
// 	})
// }

// CommitNetConfig Marks the network configuration as good and persists it inside the host.
// An API user commits the network configuration to persist a host network interface attachment or detachment, or persist the creation and deletion of a bonded interface.
func (host *Host) CommitNetConfig(async string) error {
	return host.DoAction("commitnetconfig", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// Deactivate the host to perform maintenance tasks.
func (host *Host) Deactivate(async, reason, stopGlusterService string) error {
	return host.DoAction("deactivate", struct {
		Async              string `json:"async,omitempty"`
		Reason             string `json:"reason,omitempty"`
		StopGlusterService string `json:"stop_gluster_service,omitempty"`
	}{
		async,
		reason,
		stopGlusterService,
	})
}

// EnrollCertificate Enroll certificate of the host. Useful in case you get a warning that it is about to, or already expired.
func (host *Host) EnrollCertificate(async string) error {
	return host.DoAction("enrolcertificate", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// Fence Controls host's power management device.
func (host *Host) Fence(async, fenceType string) error {
	return host.DoAction("fence", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// ForceSelectSPM Manually set a host as the storage pool manager (SPM).
func (host *Host) ForceSelectSPM(async string) error {
	return host.DoAction("fence", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// Install VDSM and related software on the host. The host type defines additional parameters for the action.
func (host *Host) Install(async, deployHostedEngine, undeployHostedEngine, image, rootPassword string, additionalParameters *Host, ssh *SSH) error {
	return host.DoAction("install", struct {
		Async                string `json:"async,omitempty"`
		DeployHostedEngine   string `json:"deploy_hosted_engine,omitempty"`
		UndeployHostedEngine string `json:"undeploy_hosted_engine,omitempty"`
		Image                string `json:"image,omitempty"`
		RootPassword         string `json:"root_password,omitempty"`
		Host                 *Host  `json:"host,omitempty"`
		SSH                  *SSH   `json:"ssh,omitempty"`
	}{
		async,
		deployHostedEngine,
		undeployHostedEngine,
		image,
		rootPassword,
		additionalParameters,
		ssh,
	})
}

// ISCSIDiscover Discover iSCSI targets on the host, using the initiator details.
func (host *Host) ISCSIDiscover(async string, iscsi *ISCSIDetails) error {
	return host.DoAction("iscsidiscover", struct {
		Async string        `json:"async,omitempty"`
		ISCSI *ISCSIDetails `json:"iscsi,omitempty"`
	}{
		async,
		iscsi,
	})
}

// ISCSILogin Login to iSCSI targets on the host, using the target details.
func (host *Host) ISCSILogin(async string, iscsi *ISCSIDetails) error {
	return host.DoAction("iscsilogin", struct {
		Async string        `json:"async,omitempty"`
		ISCSI *ISCSIDetails `json:"iscsi,omitempty"`
	}{
		async,
		iscsi,
	})
}

// Refresh the host devices and capabilities.
func (host *Host) Refresh(async string) error {
	return host.DoAction("refresh", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// // SetupNetwork This method is used to change the configuration of the network interfaces of a host.
// func (host *Host) SetupNetwork(async, checkConnectivity string, connectivityTimeout int, modifiedBonds []HostNic, modifiedLabels []NetworkLabels, modifiedNetworkAttachments []NetworkAttachments, removedBonds []HostNIC, removedLabels []NetworkLabel, removedNetworkAttachments []NetworkAttachment, synchronizedNetworkAttachments []NetworkAttachments) error {
// 	return host.DoAction("refresh", struct {
// 		Async                          string               `json:"asyncomitempty"`
// 		CheckConnectivity              string               `json:"asyncomitempty"`
// 		ConnectivityTimeout            int                  `json:"connectivity_timeoutomitempty"`
// 		ModifiedBonds                  []HostNic            `json:"modified_bondsomitempty"`
// 		ModifiedLabels                 []NetworkLabels      `json:"modified_labelsomitempty"`
// 		ModifiedNetworkAttachments     []NetworkAttachments `json:"modified_attachmentsomitempty"`
// 		RemovedBonds                   []HostNIC            `json:"removed_bondsomitempty"`
// 		RemovedLabels                  []NetworkLabel       `json:"removed_labelsomitempty"`
// 		RemovedNetworkAttachments      []NetworkAttachment  `json:"removed_network_attachmentsomitempty"`
// 		SynchronizedNetworkAttachments []NetworkAttachment  `json:"synchonized_network_attachmentsomitempty"`
// 	}{
// 		async,
// 		checkConnectivity,
// 		connectivityTimeout,
// 		modifiedBonds,
// 		modifiedLabels,
// 		modifiedNetworkAttachments,
// 		removedBonds,
// 		removedLabels,
// 		removedNetworkAttachments,
// 		synchronizedNetworkAttachments,
// 	})
// }

func (host *Host) UnregisteredStorageDomainsDiscover(async string, iscsi *ISCSIDetails) error {
	return host.DoAction("unregisteredstoragedomainsdiscover", struct {
		Async string        `json:"async,omitempty"`
		ISCSI *ISCSIDetails `json:"iscsi,omitempty"`
	}{
		async,
		iscsi,
	})
}

// Upgrade VDSM and selected software on the host.
func (host *Host) Upgrade(async string) error {
	return host.DoAction("upgrade", struct {
		Async string `json:"async,omitempty"`
	}{
		async,
	})
}

// Upgradecheck Check if there are upgrades available for the host. If there are upgrades available an icon will be displayed next to host status icon in the webadmin. Audit log messages are also added to indicate the availability of upgrades. The upgrade can be started from the webadmin or by using the upgrade host action.
func (host *Host) UpgradeCheck() error {
	return host.DoAction("upgradecheck", struct{}{})
}

// GetHost retrieve a host from the server
func (con *Connection) GetHost(id string) (*Host, error) {
	body, err := con.GetLinkBody("hosts", id)
	if err != nil {
		return nil, err
	}
	host := con.NewHost()
	err = json.Unmarshal(body, host)
	if err != nil {
		return nil, err
	}
	return host, err
}

// Update Synchronize the local object with a copy from the server
func (host *Host) Update() error {
	if host.Href == "" {
		return fmt.Errorf("host has not been saved to the server")
	}
	body, err := host.Con.Request("GET", host.Con.ResolveLink(host.Href), nil)
	if err != nil {
		return err
	}
	tempHost := Host{OvirtObject: OvirtObject{Con: host.Con}}
	err = json.Unmarshal(body, &tempHost)
	if err != nil {
		return err
	}
	*host = tempHost
	return nil
}

// GetAllHosts Retrieve all the hosts from the server
func (con *Connection) GetAllHosts() ([]*Host, error) {
	body, err := con.GetLinkBody("hosts", "")
	if err != nil {
		return nil, err
	}
	hosts := []*Host{}
	err = json.Unmarshal(body, &struct {
		Host *[]*Host
	}{&hosts})
	if err != nil {
		return nil, err
	}
	for _, host := range hosts {
		host.Con = con
	}
	return hosts, err
}

// NewHost Create a new host structure
func (con *Connection) NewHost() *Host {
	return &Host{OvirtObject: OvirtObject{Con: con}}
}

// Save Updates the server with the local copy of the host
func (host *Host) Save() error {
	body, err := json.MarshalIndent(host, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved host, we need to update it
	if host.Href != "" {
		body, err = host.Con.Request("PUT", host.Con.ResolveLink(host.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := host.Con.GetLink("hosts")
		if err != nil {
			return err
		}
		body, err = host.Con.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempHost := Host{OvirtObject: OvirtObject{Con: host.Con}}
	err = json.Unmarshal(body, &tempHost)
	if err != nil {
		return err
	}
	*host = tempHost
	return nil
}
