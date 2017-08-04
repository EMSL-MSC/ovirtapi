package ovirtapi_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/EMSL-MSC/ovirtapi"
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
	con, err := ovirtapi.NewConnection(url, username, password)
	con.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG_TRANSPORT"))
	if err != nil {
		t.Error("error creating con connection", err)
		return
	}
	newVM := con.NewVM()
	newVM.Name = "test-vm"
	display := ovirtapi.Display{}
	display.Type = "vnc"
	newVM.Display = &display
	allTemplates, err := con.GetAllTemplates()
	if err != nil {
		t.Error("Error finding a Template to assign to a test vm")
		return
	}
	newVM.Template = allTemplates[0]
	allClusters, err := con.GetAllClusters()
	if err != nil {
		t.Error("Error finding a Cluster to assign to a test vm")
		return
	}
	newVM.Cluster = allClusters[0]
	err = newVM.Save()
	if err != nil {
		t.Error("Error creating new vm", err)
		return
	}
	retrievedVM, err := con.GetVM(newVM.ID)
	if err != nil {
		t.Error("Error retrieving vm", err)
		return
	}
	for count := 0; retrievedVM.Status != "down" && count < 30; count++ {
		time.Sleep(2 * time.Second)
		err = retrievedVM.Update()
		if err != nil {
			t.Error("Error updating VM", err)
			return
		}
	}
	err = retrievedVM.Start("", "", "", "", "", nil)
	if err != nil {
		t.Error("Error starting vm", err)
		return
	}
	err = retrievedVM.Stop("false")
	if err != nil {
		t.Error("Error stopping vm", err)
		return
	}
	retrievedVM.Description = "about to delete"
	err = retrievedVM.Save()
	if err != nil {
		fmt.Printf("%s\n", retrievedVM.Status)
		t.Error("Error updating vm", err)
		return
	}
	for count := 0; retrievedVM.Status != "down" && count < 30; count++ {
		time.Sleep(2 * time.Second)
		err = retrievedVM.Update()
		if err != nil {
			t.Error("Error updating VM", err)
			return
		}
	}
	err = retrievedVM.Delete()
	if err != nil {
		t.Error("Error Deleting vm", err)
		return
	}
}
