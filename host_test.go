// Copyright (C) 2017 Battelle Memorial Institute
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD-2 license.  See the LICENSE file for details.

package ovirtapi_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/EMSL-MSC/ovirtapi"
)

func TestHost(t *testing.T) {
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
	retrievedHosts, err := con.GetAllHosts()
	if err != nil {
		t.Fatal("Error retrieving host", err)
	}
	savedComment := retrievedHosts[0].Comment
	retrievedHosts[0].Comment = "testing Description change"
	err = retrievedHosts[0].Save()
	if err != nil {
		t.Fatal("Error updating host", err)
	}
	retrievedHosts[0].Comment = savedComment
	err = retrievedHosts[0].Save()
	if err != nil {
		t.Fatal("Error reverting description on host", err)
	}
}
