package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/emsl-msc/ovirtapi"
)

func TestTemplate(t *testing.T) {
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
	newTemplate := api.NewTemplate()
	newTemplate.Name = "test-Template"
	allVMs, err := api.GetAllVMs()
	if err != nil {
		t.Error("Error finding a vm to duplicate")
		return
	}
	for _, vm := range allVMs {
		if vm.Status != "up" && vm.Status != "locked" {
			newTemplate.VM = vm
			break
		}
	}
	err = newTemplate.Save()
	if err != nil {
		t.Fatal("Error creating new Template", err)
	}
	retrievedTemplate, err := api.GetTemplate(newTemplate.Id)
	if err != nil {
		t.Fatal("Error retrieving Template", err)
	}
	for retrievedTemplate.Status == "locked" {
		time.Sleep(2 * time.Second)
		err = retrievedTemplate.Update()
		if err != nil {
			t.Fatal("Error retrieving Template", err)
		}
	}
	retrievedTemplate.Description = "about to delete"
	err = retrievedTemplate.Save()
	if err != nil {
		t.Fatal("Error updating Template", err)
	}
	for retrievedTemplate.Status == "locked" {
		time.Sleep(2 * time.Second)
		err = retrievedTemplate.Update()
		if err != nil {
			t.Fatal("Error retrieving Template", err)
		}
	}
	err = retrievedTemplate.Delete()
	if err != nil {
		t.Fatal("Error Deleting Template", err)
	}
}
