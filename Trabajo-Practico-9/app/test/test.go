package test

import (
	"testing"

	main "github.com/SantiagoMerlo/IS3-Merlo/Trabajo-Practico-9/app/src/main"
)

// TestHello is
func TestHello(t *testing.T) {
	want := "Hello, Word."
	if got := main.Msj(); got != want {
		t.Errorf(" Msj() = %q, want %q", got, want)
	}
}

// TestMsjWithParams is
func TestMsjWithParams(t *testing.T) {
	want := 8
	if got := main.MsjWithParams(2, 4); got != want {
		t.Errorf(" MsjWithParams() = %q, want %q", got, want)
	}
}

// TestWithIf is
func TestWithIf(t *testing.T) {
	want := "Bozz"
	if got := main.WithIf("Escuela"); got != want {
		t.Errorf("WithIf() = %q, want %q", got, want)
	}
}
