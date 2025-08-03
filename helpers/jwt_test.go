package helpers

import (
	"testing"
)

func TestGenerateToken_TidakKosong(t *testing.T) {
	token := GenerateToken("anggakrnwn")
	if token == "" {
		t.Errorf("Diharapkan token berhasil dibuat, tetapi hasilnya string kosong")
	}
}
