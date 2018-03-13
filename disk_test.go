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
	"time"

	"github.com/EMSL-MSC/ovirtapi"
)

func TestDisk(t *testing.T) {
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
	newDisk := con.NewDisk()
	newDisk.ProvisionedSize = 1024
	newDisk.Format = "cow"
	newDisk.Name = "test-disk"
	storageDomains := ovirtapi.StorageDomains{}
	storageDomains.StorageDomain = append(storageDomains.StorageDomain, ovirtapi.Link{
		ID: "dfe8e7be-e495-49a7-be2d-71aba891ceb4",
	})
	newDisk.StorageDomains = &storageDomains
	err = newDisk.Save()
	if err != nil {
		t.Error("Error creating new disk", err)
		return
	}
	retrievedDisk, err := con.GetDisk(newDisk.ID)
	if err != nil {
		t.Error("Error retrieving disk", err)
		return
	}
	for count := 0; retrievedDisk.Status != "ok" && count < 30; count++ {
		time.Sleep(2 * time.Second)
		err = retrievedDisk.Update()
		if err != nil {
			t.Error("Error updating Disk", err)
			return
		}
	}
	err = retrievedDisk.Delete()
	if err != nil {
		t.Error("Error Deleting disk", err)
		return
	}
}
