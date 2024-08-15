package mascot_test

import (
	"testing"

	"example.com/go-demo/mascot"
)

func TestMascot(t *testing.T){
	if mascot.BestMascot() != "Hello world"{
		t.Fatal("Wrong mascot :(")
	}
}