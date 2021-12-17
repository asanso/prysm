package ssz_static

import (
	"testing"

	"github.com/prysmaticlabs/prysm/testing/spectest/shared/phase0/ssz_generic"
)

func TestMainnet_Phase0_SSZGeneric(t *testing.T) {
	ssz_generic.RunSSZGenericTests(t, "mainnet")
}
