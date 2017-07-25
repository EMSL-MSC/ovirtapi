package ovirtapi

import (
	"encoding/xml"
)

type DataCenter struct {
	OvirtObject
	XMLName xml.Name `xml:"data_center"`
	Local string `xml:"local,omitempty"`
	QuotaMode string `xml:"quota_mode,omitempty"`
	Status string `xml:"status,omitempty"`
	StorageFormat string `xml:"storage_format,omitempty"`
	SupportedVersions []struct {
		Version struct {
			Major int `xml:"major,omitempty"`
			Minor int `xml:"minor,omitempty"`
		} `xml:"version,omitempty"`
	} `xml:"supported_versions,omitempty"`
}
