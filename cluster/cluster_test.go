package cluster_test

import (
	"os"
	"testing"
	"github.com/emsl-msc/ovirtapi"
	"github.com/emsl-msc/ovirtapi/cluster"
)

func TestCluster(t *testing.T) {
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
	api.Debug = true
	if err != nil {
		t.Error("error creating api connection", err)
		return
	}
	newCluster := cluster.New(api)
	newCluster.Name = "test-cluster"
	newCluster.Cpu.Architecture = "x86_64"
	newCluster.Cpu.Type = "Intel Haswell-noTSX Family"
	newCluster.DataCenter = &ovirtapi.Link{Id: "00000001-0001-0001-0001-000000000311"}
	err = newCluster.Save()
	if err != nil {
		t.Fatal("Error creating new cluster", err)
	}
	retrievedCluster, err := cluster.Get(api, newCluster.Id)
	if err != nil {
		t.Fatal("Error retrieving cluster", err)
	}
	retrievedCluster.Description = "about to delete"
	err = retrievedCluster.Save()
	if err != nil {
		t.Fatal("Error updating cluster", err)
	}
	err = retrievedCluster.Delete()
	if err != nil {
		t.Fatal("Error Deleting cluster", err)
	}
}
