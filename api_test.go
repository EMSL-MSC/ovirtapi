package ovirtapi

import (
	"testing"
)

func TestNewAPI(t *testing.T) {
	_, err := NewAPI("https://c0.emsl.pnl.gov:443/ovirt-engine/api/", "admin@internal", "boog7Urea")
	if err == nil {
		t.Error("Did not fail when passed bad password")
	}
	_, err = NewAPI("https://c0.emsl.pnl.gov:443/ovirt-engine/api/", "admin@internal", "boog7Ure")
	if err != nil {
		t.Error("Did not create new API")
	}
}
