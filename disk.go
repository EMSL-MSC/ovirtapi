package ovirtapi

import (
	"encoding/json"
	"fmt"
)

type VolumeGroup struct {
	id            string        `json:"id,omitempty"`
	logical_units []LogicalUnit `json:"logical_units,omitempty"`
	name          string        `json:"name,omitempty"`
}

type LogicalUnit struct {
	Address        string `json:"address,omitempty"`
	DiscardMaxSize int    `json:"discard_max_size,omitempty"`
	// The maximum number of bytes that can be discarded by the logical unit's underlying storage in a single operation.
	DiscardZeroesData string `json:"discard_zeroes_data,omitempty"`
	// True, if previously discarded blocks in the logical unit's underlying storage are read back as zeros.
	DiskID          string `json:"disk_id,omitempty"`
	ID              string `json:"id,omitempty"`
	LUNMapping      int    `json:"lun_mapping,omitempty,string"`
	Password        string `json:"password,omitempty"`
	Paths           int    `json:"paths,omitempty,string"`
	Port            int    `json:"port,omitempty,string"`
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

type HostStorage struct {
	Address string `json:"address,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID           string        `json:"id,omitempty"`
	LogicalUnits []LogicalUnit `json:"logical_units,omitempty"`
	MountOptions string        `json:"mount_options,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	// The number of times to retry a request before attempting further recovery actions.
	NfsRetrans int `json:"nfs_retrans,omitempty,string"`
	// The time in tenths of a second to wait for a response before retrying NFS requests.
	NfsTimeo     int          `json:"nfs_timeo,omitempty,string"`
	NfsVersion   string       `json:"nfs_version,omitempty"`
	OverrideLUNS string       `json:"override_luns,omitempty"`
	Password     string       `json:"password,omitempty"`
	Path         string       `json:"path,omitempty"`
	Port         int          `json:"port,omitempty,string"`
	Portal       string       `json:"portal,omitempty"`
	Target       string       `json:"target,omitempty"`
	Type         string       `json:"type,omitempty"`
	Username     string       `json:"username,omitempty"`
	VfsType      string       `json:"vfs_type,omitempty"`
	VolumeGroup  *VolumeGroup `json:"volume_group,omitempty"`
}

type StorageDomains struct {
	//TODO make StorageDomain
	StorageDomain []Link `json:"storage_domain,omitempty"`
}

type Disk struct {
	OvirtObject
	//  Indicates if the disk is visible to the virtual machine.
	Active string `json:"active,omitempty"`
	//  The actual size of the disk, in bytes.
	ActualSize int    `json:"actual_size,omitempty"`
	Alias      string `json:"alias,omitempty"`
	//  Indicates if the disk is marked as bootable.
	Bootable string `json:"bootable,omitempty"`
	//  Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	//  The underlying storage format.
	Format  string `json:"format,omitempty"`
	ImageID string `json:"image_id,omitempty"`
	//  The initial size of a sparse image disk created on block storage, in bytes.
	InitialSize int `json:"initial_size,omitempty"`
	//  The type of interface driver used to connect the disk device to the virtual machine.
	Interface   string       `json:"interface,omitempty"`
	LogicalName string       `json:"logical_name,omitempty"`
	LunStorage  *HostStorage `json:"lun_storage,omitempty"`
	//  Indicates if disk errors should cause virtual machine to be paused or if disk errors should be propagated to the the guest operating system instead.
	PropagateErrors string `json:"propagate_errors,omitempty"`
	//  The virtual size of the disk, in bytes.
	ProvisionedSize int `json:"provisioned_size,omitempty"`
	//  The underlying QCOW version of a QCOW volume.
	QcowVersion string `json:"qcow_version,omitempty"`
	//  Indicates if the disk is in read-only mode.
	ReadOnly string `json:"read_only,omitempty"`
	SGIO     string `json:"sgio,omitempty"`
	//  Indicates if the disk can be attached to multiple virtual machines.
	Shareable string `json:"shareable,omitempty"`
	//  Indicates if the physical storage for the disk should not be preallocated.
	Sparse string `json:"sparse,omitempty"`
	//  The status of the disk device.
	Status              string `json:"status,omitempty"`
	StorageType         string `json:"storage_type,omitempty"`
	UsesSCSIReservation string `json:"uses_scsi_reservation,omitempty"`
	// Indicates if the disk's blocks will be read back as zeros after it is deleted:
	//
	// - On block storage, the disk will be zeroed and only then deleted.
	WipeAfterDelete string `json:"wipe_after_delete,omitempty"`
	// TODO Make DiskProfile
	DiskProfile *Link `json:"disk_profile,omitempty"`
	// Optionally references to an instance type the device is used by.
	// TODO Make InstanceType
	InstanceType *Link `json:"instance_type,omitempty"`
	// TODO Make OpenStackVolumeType
	OpenstackVolumeType *Link `json:"openstack_volume_type,omitempty"`
	// TODO Make Permission
	Permissions []Link `json:"permissions,omitempty"`
	// TODO Make Quota
	Quota *Link `json:"quota,omitempty"`
	// TODO Make Snapshot
	Snapshot *Link `json:"snapshot,omitempty"`
	// Statistics exposed by the disk.
	// TODO Make Statistic
	Statistics []Link `json:"statistics,omitempty"`
	// The storage domains associated with this disk.
	// TODO Make StorageDomain
	StorageDomains *StorageDomains `json:"storage_domains,omitempty"`
	// Optionally references to a template the device is used by.
	Template *Template `json:"template,omitempty"`
	// References to the virtual machines that are using this device.
	VMs []VM `json:"vms,omitempty"`
}

func (con *Connection) GetDisk(id string) (*Disk, error) {
	body, err := con.GetLinkBody("disks", id)
	if err != nil {
		return nil, err
	}
	disk := con.NewDisk()
	err = json.Unmarshal(body, disk)
	if err != nil {
		return nil, err
	}
	return disk, err
}

// Update Synchronize the local Disk with a copy from the server
func (disk *Disk) Update() error {
	if disk.Href == "" {
		return fmt.Errorf("Disk has not been saved to the server")
	}
	newDisk, err := disk.Con.GetDisk(disk.ID)
	if err != nil {
		return err
	}
	*disk = *newDisk
	return nil
}

func (con *Connection) GetAllDisks() ([]*Disk, error) {
	body, err := con.GetLinkBody("disks", "")
	if err != nil {
		return nil, err
	}
	disks := []*Disk{}
	err = json.Unmarshal(body, &struct {
		Disk *[]*Disk
	}{&disks})
	if err != nil {
		return nil, err
	}
	for _, disk := range disks {
		disk.Con = con
	}
	return disks, err
}

func (con *Connection) NewDisk() *Disk {
	return &Disk{OvirtObject: OvirtObject{Con: con}}
}

func (disk *Disk) Save() error {
	body, err := json.MarshalIndent(disk, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved disk, we need to update it
	if disk.OvirtObject.Href != "" {
		body, err = disk.Con.Request("PUT", disk.Con.ResolveLink(disk.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := disk.Con.GetLink("disks")
		if err != nil {
			return err
		}
		body, err = disk.Con.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempDisk := Disk{OvirtObject: OvirtObject{Con: disk.Con}}
	err = json.Unmarshal(body, &tempDisk)
	if err != nil {
		return err
	}
	*disk = tempDisk
	return nil
}

// Copy This operation copies a disk to the specified storage domain.
func (vm *VM) Copy(async string, disk *Disk, filter string, storageDomain *StorageDomain) error {
	return vm.DoAction("copy", Action{
		Async:         aync,
		Disk:          disk,
		Filter:        filter,
		StorageDomain: StorageDomain,
	})
}

func (vm *VM) Export(async, filter string, storageDomain *StorageDomain) error {
	return vm.DoAction("export", Action{
		Async:         aync,
		Disk:          disk,
		Filter:        filter,
		StorageDomain: StorageDomain,
	})
}

// Move a disk to another storage domain.
func (vm *VM) Move(async, filter string, storageDomain *StorageDomain) error {
	return vm.DoAction("move", Action{
		Async:         aync,
		Disk:          disk,
		Filter:        filter,
		StorageDomain: StorageDomain,
	})
}

// Sparsify the disk.
func (vm *VM) Sparsify() error {
	return vm.DoAction("move", Action{})
}
