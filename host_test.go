package ovirtapi_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/EMSL-MSC/ovirtapi"
)

func TestHost(t *testing.T) {
	t.Skip()
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
	newHost := con.NewHost()
	newHost.Name = "test-host"
	newHost.Address = "test-host"
	newHost.RootPassword = "test-pass"
	err = newHost.Save()
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Fatal("Error creating new host", err)
	}
	retrievedHost, err := con.GetHost(newHost.ID)
	if err != nil {
		t.Fatal("Error retrieving host", err)
	}
	retrievedHost.Description = "about to delete"
	err = retrievedHost.Save()
	if err != nil {
		t.Fatal("Error updating host", err)
	}
	err = retrievedHost.Delete()
	if err != nil {
		t.Fatal("Error Deleting host", err)
	}
}
