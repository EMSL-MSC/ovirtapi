package ovirtapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type Link struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
	ID   string `json:"id,omitempty"`
}

type OvirtObject struct {
	Link
	Con         *Connection `json:"-"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Actions     struct {
		Links []Link `json:"link,omitempty"`
	} `json:"actions,omitempty"`
	Links []Link `json:"link,omitempty"`
}

type Action struct {
	AllowPartialImport string         `json:"allow_partial_import,omitempty"`
	Async              string         `json:"async,omitempty"`
	Bricks             []GlusterBrick `json:"bricks,omitempty"`
	Certificates       []Certificate  `json:"certificates,omitempty"`
	CheckConnectivity  string         `json:"check_connectivity,omitempty"`
	Clone              string         `json:"clone,omitempty"`
	Cluster            *Cluster       `json:"cluster,omitempty"`
	CollapseSnapshots  string         `json:"collapse_snapshots,omitempty"`
	// Free text containing comments about this object.
	Comment             string      `json:"comment,omitempty"`
	ConnectivityTimeout int         `json:"connectivity_timeout,omitempty"`
	DataCenter          *DataCenter `json:"data_center,omitempty"`
	DeployHostedEngine  string      `json:"deploy_hosted_engine,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// TODO: Details          GlusterVolumeProfileDetails `json:"details,omitempty"`
	DiscardSnapshots string `json:"discard_snapshots,omitempty"`
	Disk             *Disk  `json:"disk,omitempty"`
	// TODO: Disks            []Disk                      `json:"disks,omitempty"`
	Exclusive string `json:"exclusive,omitempty"`
	Fault     *Fault `json:"fault,omitempty"`
	FenceType string `json:"fence_type,omitempty"`
	Filter    string `json:"filter,omitempty"`
	FixLayout string `json:"fix_layout,omitempty"`
	Force     string `json:"force,omitempty"`
	// TODO: GracePeriod      GracePeriod                 `json:"grace_period,omitempty"`
	Host *Host `json:"host,omitempty"`
	// A unique identifier.
	ID               string        `json:"_i_d,omitempty"`
	Image            string        `json:"image,omitempty"`
	ImportAsTemplate string        `json:"import_as_template,omitempty"`
	IsAttached       string        `json:"is_attached,omitempty"`
	ISCSI            *ISCSIDetails `json:"iscsi,omitempty"`
	IscsiTargets     []string      `json:"iscsi_targets,omitempty"`
	// TODO: Job                        Job                 `json:"job,omitempty"`
	// TODO: LogicalUnits               []LogicalUnit       `json:"logical_units,omitempty"`
	MaintenanceEnabled string `json:"maintenance_enabled,omitempty"`
	// TODO: ModifiedBonds              []HostNic           `json:"modified_bonds,omitempty"`
	// TODO: ModifiedLabels             []NetworkLabel      `json:"modified_labels,omitempty"`
	// TODO: ModifiedNetworkAttachments []NetworkAttachment `json:"modified_network_attachments,omitempty"`
	// A human-readable name in plain text.
	Name            string           `json:"name,omitempty"`
	Option          *Option          `json:"option,omitempty"`
	Pause           string           `json:"pause,omitempty"`
	PowerManagement *PowerManagement `json:"power_management,omitempty"`
	// TODO: ProxyTicket                    ProxyTicket                          `json:"proxy_ticket,omitempty"`
	Reason                     string `json:"reason,omitempty"`
	ReassignBadMacs            string `json:"reassign_bad_macs,omitempty"`
	RemoteViewerConnectionFile string `json:"remote_viewer_connection_file,omitempty"`
	// TODO: RemovedBonds                   []HostNic                            `json:"removed_bonds,omitempty"`
	// TODO: RemovedLabels                  []NetworkLabel                       `json:"removed_labels,omitempty"`
	// TODO: RemovedNetworkAttachments      []NetworkAttachment                  `json:"removed_network_attachments,omitempty"`
	ResolutionType string `json:"resolution_type,omitempty"`
	RestoreMemory  string `json:"restore_memory,omitempty"`
	RootPassword   string `json:"root_password,omitempty"`
	// TODO: Snapshot                       Snapshot                             `json:"snapshot,omitempty"`
	SSH                *SSH           `json:"ssh,omitempty"`
	Status             string         `json:"status,omitempty"`
	StopGlusterService string         `json:"stop_gluster_service,omitempty"`
	// TODO: StorageDomain      *StorageDomain `json:"storage_domain,omitempty"`
	// TODO: StorageDomains                 []StorageDomain                      `json:"storage_domains,omitempty"`
	Succeeded string `json:"succeeded,omitempty"`
	// TODO: SynchronizedNetworkAttachments []NetworkAttachment                  `json:"synchronized_network_attachments,omitempty"`
	Template *Template `json:"template,omitempty"`
	// TODO: Ticket                         Ticket                               `json:"ticket,omitempty"`
	UnDeployHostedEngine string `json:"undeploy_hosted_engine,omitempty"`
	UseCloudInit         string `json:"use_cloud_init,omitempty"`
	UseSysPrep           string `json:"use_sysprep,omitempty"`
	// TODO: VirtualFunctionsConfiguration  HostNicVirtualFunctionsConfiguration `json:"virtual_functions_configuration,omitempty"`
	VM *VM `json:"_v_m,omitempty"`
	// TODO: VnicProfileMappings            []VnicProfileMapping                 `json:"vnic_profile_mappings,omitempty"`
}

func (ovirtObject *OvirtObject) DoAction(action string, parameters interface{}) (err error) {
	for _, link := range ovirtObject.Actions.Links {
		if link.Rel == action {
			var body []byte
			body, err = json.Marshal(parameters)
			if err != nil {
				return fmt.Errorf("Unable to Marshal parameters")
			}
			_, err = ovirtObject.Con.Request("POST", ovirtObject.Con.ResolveLink(link.Href), body)
			return
		}
	}
	return errors.New("Action not found")
}

func (ovirtObject *OvirtObject) GetLink(rel string) (*url.URL, error) {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			return ovirtObject.Con.ResolveLink(link.Href), nil
		}
	}
	return nil, errors.New("Link not found")
}

func (ovirtObject *OvirtObject) Delete() error {
	_, err := ovirtObject.Con.Request("DELETE", ovirtObject.Con.ResolveLink(ovirtObject.Href), nil)
	return err
}
