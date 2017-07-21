package cluster

import (
	"encoding/xml"
	"github.com/emsl-msc/ovirtapi"
)

type Cluster struct {
	ovirtapi.OvirtObject
	XMLName xml.Name `xml:"cluster"`
	BallooningEnabled string `xml:"ballooning_enabled,omitempty"`
	Cpu struct {
		Architecture string `xml:"architecure,omitempty"`
		Type string `xml:"type,omitempty"`
	} `xml:"cpu,omitempty"`
	CustomSchedulingPolicyProperties *struct {
		Property []struct {
			Name string `xml:"name,omitempty"`
			Value int `xml:"value,omitempty"`
		} `xml:"property,omitempty"`
	} `xml:"CustomSchedulingPolicyProperties"`
	ErrorHandling *struct {
		OnError string `xml:"on_error,omitempty"`
	} `xml:"ErrorHandling,omitempty"`
	FencingPolicy *struct {
		Enabled string `xml:"enabled,omitempty"`
		SkipIfConnectivitybroken *struct {
			Enabled string `xml:"enabled,omitempty"`
			Threshold int `xml:"threshold,omitempty"`
		} `xml:"skip_if_connectivity_broken,omitempty"`
		SkipIfGlusterBricksUp string `xml:"skip_if_gluster_bricks_up,omitempty"`
		SkipIfGlusterQuorumNotMet string `xml:"skip_if_gluster_quorum_not_met,omitempty"`
		SkipIfSdActive *struct {
			Enabled string `xml:"enabled,omitempty"`
		} `xml:"skip_if_sd_active,omitempty"`
	} `xml:"fencing_policy,omitempty"`
	GlusterService string `xml:"gluster_service,omitempty"`
	HaReservation string `xml:"ha_reservation,omitempty"`
	Ksm *struct {
		Enabled string `xml:"enabled,omitempty"`
		MergeAcrossNodes string `xml:"merge_across_nodes,omitempty"`
	} `xml:"ksm,omitempty"`
	MaintenanceReasonRequired string `xml:"maintenance_reason_required,omitempty"`
	MemoryPolicy *struct {
		OverCommit *struct {
			Percent int `xml:"percent,omitempty"`
		} `xml:"over_commit,omitempty"`
		TransparentHugepages *struct {
			Enabled string `xml:"enabled,omitempty"`
		} `xml:"transparent_hugepages,omitempty"`
	} `xml:"memory_policy,omitempty"`
	Migration *struct {
		AutoConverge string `xml:"auto_converge,omitempty"`
		Bandwidth *struct {
			AssignmentMethod string `xml:"assignment_method,omitempty"`
		} `xml:"bandwidth,omitempty"`
		Compressed string `xml:"compressed,omitempty"`
	} `xml:"migration,omitempty"`
	OptionalReason string `xml:"optional_reason,omitempty"`
	RequiredRngSources *struct {
		RequiredRngSource string `xml:"required_rng_source,omitempty"`
	} `xml:"required_rng_sources,omitempty"`
	SwitchType string `xml:"switch_type,omitempty"`
	ThreadsAsCores string `xml:"threads_as_cores,omitempty"`
	TrustedService string `xml:"trusted_service,omitempty"`
	TunnelMigration string `xml:"tunnel_migration,omitempty"`
	Version *struct {
		Major int `xml:"major,omitempty"`
		Minor int `xml:"minor,omitempty"`
	} `xml:"version,omitempty"`
	VirtService string `xml:"virt_service,omitempty"`
	DataCenter *ovirtapi.Link `xml:"data_center,omitempty"`
	MacPool *ovirtapi.Link `xml:"mac_pool,omitempty"`
	SchedulingPolicy *ovirtapi.Link `xml:"scheduling_policy,omitempty"`
}

func Get(api *ovirtapi.API, id string) (*Cluster, error) {
	body, err := api.GetLinkBody("clusters", id)
	if err != nil {
		return nil, err
	}
	cluster := New(api)
	err = xml.Unmarshal(body, cluster)
	if err != nil {
		return nil, err
	}
	return cluster, err
}

func GetAll(api *ovirtapi.API) ([]*Cluster, error) {
	body, err := api.GetLinkBody("clusters", "")
	if err != nil {
		return nil, err
	}
	clusters := []*Cluster{}
	err = xml.Unmarshal(body, &struct {Clusters *[]*Cluster `xml:"cluster"`}{&clusters})
	if err != nil {
		return nil, err
	}
	for _, cluster := range clusters {
		cluster.Api = api
	}
	return clusters, err
}

func NewCluster(api *ovirtapi.API) (*Cluster) {
	return &Cluster{OvirtObject: ovirtapi.OvirtObject{Api: api}}
}

func (cluster *Cluster) Save() (error) {
	body, err := xml.MarshalIndent(cluster, "", "    ")
	if err != nil {
		return err
	}
	// If there is a link, it is an already saved object, we need to update it
	if cluster.OvirtObject.Href != "" {
		body, err = cluster.Api.Request("PUT", cluster.Api.ResolveLink(cluster.Href), body)
		if err != nil {
			return err
		}
	} else {
		link, err := cluster.Api.GetLink("clusters")
		if err != nil {
			return err
		}
		body, err = cluster.Api.Request("POST", link, body)
		if err != nil {
			return err
		}
	}
	tempCluster := Cluster{OvirtObject: ovirtapi.OvirtObject{Api: cluster.Api}}
	err = xml.Unmarshal(body, &tempCluster)
	*cluster = tempCluster
	if err != nil {
		return err
	}
	return nil
}

func (cluster *Cluster) Delete() (error) {
	_, err := cluster.Api.Request("DELETE", cluster.Api.ResolveLink(cluster.Href), nil)
	return err
}
