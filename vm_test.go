package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/emsl-msc/ovirtapi"
)

func TestVM(t *testing.T) {
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
	newVM := api.NewVM()
	newVM.Name = "test-vm"
	allTemplates, err := api.GetAllTemplates()
	if err != nil {
		t.Error("Error finding a Template to assign to a test vm")
		return
	}
	newVM.Template = allTemplates[0]
	allClusters, err := api.GetAllClusters()
	if err != nil {
		t.Error("Error finding a Cluster to assign to a test vm")
		return
	}
	newVM.Cluster = allClusters[0]
	err = newVM.Save()
	if err != nil {
		t.Fatal("Error creating new vm", err)
	}
	retrievedVM, err := api.GetVM(newVM.Id)
	if err != nil {
		t.Fatal("Error retrieving vm", err)
	}
	retrievedVM.Description = "about to delete"
	err = retrievedVM.Save()
	if err != nil {
		t.Fatal("Error updating vm", err)
	}
	err = retrievedVM.Delete()
	if err != nil {
		t.Fatal("Error Deleting vm", err)
	}
}
