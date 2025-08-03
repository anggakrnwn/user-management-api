package config

import (
	"os"
	"testing"
)

func TestGetEnv_MengembalikanNilaiYangTelahDisetel(t *testing.T) {
	os.Setenv("APP_PORT", "9999")
	defer os.Unsetenv("APP_PORT") // Membersihkan setelah test

	hasil := GetEnv("APP_PORT", "3000")
	if hasil != "9999" {
		t.Errorf("Seharusnya mendapatkan '9999', tetapi malah mendapat '%s'", hasil)
	}
}

func TestGetEnv_MengembalikanNilaiDefault(t *testing.T) {
	hasil := GetEnv("KUNCI_TIDAK_ADA", "fallback")
	if hasil != "fallback" {
		t.Errorf("Seharusnya mendapatkan 'fallback', tetapi malah mendapat '%s'", hasil)
	}
}
