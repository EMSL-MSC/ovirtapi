package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/EMSL-MSC/ovirtapi"
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
	con, err := ovirtapi.NewConnection(url, username, password)
	con.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG_TRANSPORT"))
	if err != nil {
		t.Error("error creating con connection", err)
		return
	}
	newTemplate := con.NewTemplate()
	newTemplate.Name = "test-Template"
	allVMs, err := con.GetAllVMs()
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
	retrievedTemplate, err := con.GetTemplate(newTemplate.ID)
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
