package ovirtapi_test

import (
	"os"
	"testing"
	"fmt"
	"github.com/emsl-msc/ovirtapi"
)

func TestNewAPI(t *testing.T) {
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
	if err != nil {
		t.Fatal("Did not create new API", err)
		return
	}
	fmt.Printf("%+v\n", api)
	_, err = ovirtapi.NewAPI(url, "baduser", "badpass")
	if err == nil {
		t.Error("Did not fail when passed bad password", err)
	}
}
