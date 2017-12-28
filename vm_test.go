package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/EMSL-MSC/ovirtapi"
)

func TestVM(t *testing.T) {
	t.Parallel()
	username := os.Getenv("OVIRT_USERNAME")
	if username == "" {
		t.Error("OVIRT_USERNAME is not set")
		return
	}
	password := os.Getenv("OVIRT_PASSWORD")
	if password == "" {
		t.Error("OVIRT_PASSWORD is not set")
		return
	}
	url := os.Getenv("OVIRT_URL")
	if url == "" {
		t.Error("OVIRT_URL is not set")
		return
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_TRANSPORT"))
	con, err := ovirtapi.NewConnection(url, username, password, debug)
	if err != nil {
		t.Error("error creating connection", err)
		return
	}
	newVM := con.NewVM()
	newVM.Name = "test-vm"
	display := ovirtapi.Display{}
	display.Type = "vnc"
	newVM.Display = &display
	allTemplates, err := con.GetAllTemplates()
	if err != nil {
		t.Error("Error finding a Template to assign to a test vm", err)
		return
	}
	newVM.Template = allTemplates[0]
	allClusters, err := con.GetAllClusters()
	if err != nil {
		t.Error("Error finding a Cluster to assign to a test vm", err)
		return
	}
	newVM.Cluster = allClusters[0]
	err = newVM.Save()
	if err != nil {
		t.Error("Error creating new vm", err)
		return
	}
	newDisk := con.NewDisk()
	newDisk.ProvisionedSize = 102400
	newDisk.Format = "cow"
	newDisk.Name = "attach-disk"
	storageDomains := ovirtapi.StorageDomains{}
	storageDomains.StorageDomain = append(storageDomains.StorageDomain, ovirtapi.Link{
		ID: "dfe8e7be-e495-49a7-be2d-71aba891ceb4",
	})
	newDisk.StorageDomains = &storageDomains
	err = newDisk.Save()
	if err != nil {
		t.Fatal("Error creating a disk to attach to the vm", err)
	}
	for count := 0; (newDisk.Status != "ok" || newVM.Status != "down") && count < 30; count++ {
		time.Sleep(2 * time.Second)
		err = newDisk.Update()
		if err != nil {
			t.Error("Error updating Disk", err)
			return
		}
		err = newVM.Update()
		if err != nil {
			t.Error("Error updating VM", err)
			return
		}
	}
	err = newVM.AddLink("diskattachments", ovirtapi.DiskAttachment{
		Active:      "true",
		Bootable:    "true",
		Disk:        newDisk,
		Interface:   "virtio_scsi",
		LogicalName: "/dev/vdb",
	}, nil)
	if err != nil {
		t.Fatal("Error attaching disk to the vm", err)
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
	retrievedVM.RemoveLink("diskattachments", newDisk.ID, nil)
	err = retrievedVM.Delete()
	if err != nil {
		t.Error("Error Deleting vm", err)
		return
	}
}
