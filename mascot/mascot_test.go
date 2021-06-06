package mascot_test

import (
	"testing"

	"go.silversonic.net/go-demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot :(")
	}
}
