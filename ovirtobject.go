// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi

import (
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
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
	Actions     *Actions    `json:"actions,omitempty"`
	Links       []Link      `json:"link,omitempty"`
}

type Actions struct {
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
	SSH                *SSH   `json:"ssh,omitempty"`
	Status             string `json:"status,omitempty"`
	StopGlusterService string `json:"stop_gluster_service,omitempty"`
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
				return err
			}
			_, err = ovirtObject.Con.Request("POST", ovirtObject.Con.ResolveLink(link.Href), body)
			return
		}
	}
	return errors.New("Action not found")
}

type linkResponse struct {
	DiskAttachment []DiskAttachment `json:"disk_attachment,omitempty"`
	NIC            []NIC            `json:"nic,omitempty"`
}

func (ovirtObject *OvirtObject) GetLink(rel string) (*url.URL, error) {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			return ovirtObject.Con.ResolveLink(link.Href), nil
		}
	}
	return nil, errors.New("Link not found")
}

func (ovirtObject *OvirtObject) getLinkResponse(rel string, addParameters map[string]string) (*linkResponse, error) {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			values := url.Values{}
			for k, v := range addParameters {
				values.Add(k, v)
			}
			href := ovirtObject.Con.ResolveLink(link.Href)
			href.RawQuery = values.Encode()
			body, err := ovirtObject.Con.Request("GET", href, nil)
			if err != nil {
				return nil, err
			}
			linkResp := &linkResponse{}
			err = json.Unmarshal(body, linkResp)
			return linkResp, err
		}
	}
	return nil, errors.New("Link not found")
}

func (ovirtObject *OvirtObject) GetLinkObject(rel string, id string, addParameters map[string]string) (interface{}, error) {
	linkResp, err := ovirtObject.getLinkResponse(rel, addParameters)
	if err != nil {
		return nil, err
	}
	val := reflect.ValueOf(*linkResp)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		for j := 0; j < field.Len(); j++ {
			return field.Index(j).Interface(), nil
		}
	}
	return nil, errors.New("Connection not found")
}

func (ovirtObject *OvirtObject) AddLinkObject(rel string, newObject interface{}, addParameters map[string]string) (string, error) {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			var body []byte
			body, err := json.MarshalIndent(newObject, "", "    ")
			if err != nil {
				return "", err
			}
			values := url.Values{}
			for k, v := range addParameters {
				values.Add(k, v)
			}
			href := ovirtObject.Con.ResolveLink(link.Href)
			href.RawQuery = values.Encode()
			resp, err := ovirtObject.Con.Request("POST", href, body)
			respLink := Link{}
			err = json.Unmarshal(resp, respLink)
			if err != nil {
				return "", err
			}
			return respLink.ID, err
		}
	}
	return "", errors.New("Link not found")
}

func (ovirtObject *OvirtObject) RemoveLinkObject(rel string, id string, addParameters map[string]string) error {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			values := url.Values{}
			for k, v := range addParameters {
				values.Add(k, v)
			}
			href := ovirtObject.Con.ResolveLink(link.Href + "/" + id)
			href.RawQuery = values.Encode()
			_, err := ovirtObject.Con.Request("DELETE", href, nil)
			return err
		}
	}
	return errors.New("Link not found")
}

func (ovirtObject *OvirtObject) Delete() error {
	_, err := ovirtObject.Con.Request("DELETE", ovirtObject.Con.ResolveLink(ovirtObject.Href), nil)
	return err
}
