package ovirtapi_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/EMSL-MSC/ovirtapi"
)

func TestDataCenter(t *testing.T) {
	t.Parallel()
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
	con, err := ovirtapi.NewConnection(url, username, password, false)
	con.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG_TRANSPORT"))
	if err != nil {
		t.Error("error creating connection", err)
		return
	}
	newDataCenter := con.NewDataCenter()
	newDataCenter.Name = "test-data-center"
	newDataCenter.Local = "true"
	err = newDataCenter.Save()
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Fatal("Error creating new data center", err)
	}
	retrievedDataCenter, err := con.GetDataCenter(newDataCenter.ID)
	if err != nil {
		t.Fatal("Error retrieving data center", err)
	}
	retrievedDataCenter.Description = "about to delete"
	err = retrievedDataCenter.Save()
	if err != nil {
		t.Fatal("Error updating data center", err)
	}
	err = retrievedDataCenter.Delete()
	if err != nil {
		t.Fatal("Error Deleting data center", err)
	}
}
