package ovirtapi_test

import (
	"github.com/emsl-msc/ovirtapi"
	"os"
	"testing"
)

func TestNewConnection(t *testing.T) {
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
	_, err := ovirtapi.NewConnection(url, username, password)
	if err != nil {
		t.Fatal("Did not create new Connection", err)
		return
	}
	_, err = ovirtapi.NewConnection(url, "baduser", "badpass")
	if err == nil {
		t.Error("Did not fail when passed bad password", err)
	}
}
