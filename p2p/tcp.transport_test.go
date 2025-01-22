package p2p

import (
	"testing"
	// "github.com/stretcher/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	addr := ":4000"
	tr := InitTransport(addr)

	if addr != tr.address {
		t.Errorf("got %q, wanted %q", tr.address, addr)
	}
	// assert.Equal(t, tr.address, addr)
}