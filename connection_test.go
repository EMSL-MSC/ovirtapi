// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi_test

import (
	"github.com/EMSL-MSC/ovirtapi"
	"os"
	"testing"
)

func TestNewConnection(t *testing.T) {
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
	_, err := ovirtapi.NewConnection(url, username, password, false)
	if err != nil {
		t.Fatal("Did not create new Connection", err)
		return
	}
	_, err = ovirtapi.NewConnection(url, "baduser", "badpass", false)
	if err == nil {
		t.Error("Did not fail when passed bad password", err)
	}
}
