package ovirtapi

type DataCenter struct {
	OvirtObject
	Local             string `json:"local,omitempty"`
	QuotaMode         string `json:"quota_mode,omitempty"`
	Status            string `json:"status,omitempty"`
	StorageFormat     string `json:"storage_format,omitempty"`
	SupportedVersions *struct {
		Version []struct {
			Major string `json:"major,omitempty"`
			Minor string `json:"minor,omitempty"`
		} `json:"version,omitempty"`
	} `json:"supported_versions,omitempty"`
	Version *struct {
		Major string `json:"major,omitempty"`
		Minor string `json:"minor,omitempty"`
	} `json:"version,omitempty"`
}
