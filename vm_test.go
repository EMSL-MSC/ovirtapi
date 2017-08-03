package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/emsl-msc/ovirtapi"
)

func TestVm(t *testing.T) {
	username := os.Getenv("OVIRT_USERNAME")
	if username == "" {
		t.Error("OVIRT_USERNAME is not set")
	}
	password := os.Getenv("OVIRT_PASSWORD")
	if password == "" {
		t.Error("OVIRT_PASSWORD is not set")
	}
	url := os.Getenv("OVIRT_URL")
	if url == "" {
		t.Error("OVIRT_URL is not set")
	}
	api, err := ovirtapi.NewAPI(url, username, password)
	api.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG_TRANSPORT"))
	if err != nil {
		t.Error("error creating api connection", err)
		return
	}
	newVm := api.NewVm()
	newVm.Name = "test-vm"
	newVm.Template = &ovirtapi.Link{
		Href: "/ovirt-engine/api/templates/00000000-0000-0000-0000-000000000000",
		Id:   "00000000-0000-0000-0000-000000000000",
	}
	newVm.Cluster = &ovirtapi.Link{
		Href: "/ovirt-engine/api/clusters/00000002-0002-0002-0002-00000000017a",
		Id:   "00000002-0002-0002-0002-00000000017a",
	}
	err = newVm.Save()
	if err != nil {
		t.Fatal("Error creating new vm", err)
	}
	retrievedVm, err := api.GetVm(newVm.Id)
	if err != nil {
		t.Fatal("Error retrieving vm", err)
	}
	retrievedVm.Description = "about to delete"
	err = retrievedVm.Save()
	if err != nil {
		t.Fatal("Error updating vm", err)
	}
	err = retrievedVm.Delete()
	if err != nil {
		t.Fatal("Error Deleting vm", err)
	}
}
