package codemeta

import (
	"testing"
)

func TestIsDOI(t *testing.T) {
	identifier := "10.22002/5rbqw-9cc91"
	expected := true
	got := isDOI(identifier)
	if expected != got {
		t.Errorf("for %q, expected %t, got %t", identifier, expected, got)
	}
	identifier = "0000-1111-2222-3333"
	expected = false
	got = isDOI(identifier)
	if expected != got {
		t.Errorf("for %q, expected %t, got %t", identifier, expected, got)
	}


}

func TestCodemeta(t *testing.T) {
	t.Errorf("TestCodemeta() not implemented.")
}

func TestCodemetaShow(t *testing.T) {
	t.Errorf("TestCodemetaShow() not implemented.")
}

func TestCodemetaSet(t *testing.T) {
	t.Errorf("TestCodemetaSet() not implemented.")
}

func TestCodemetaDelete(t *testing.T) {
	t.Errorf("TestCodemetaDelete() not implemented.")
}
