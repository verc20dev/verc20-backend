package web3

import (
	"bytes"
	"testing"
)

func TestIsValidTx(t *testing.T) {
	x := []byte{0xd9, 0xd9, 0xf7}
	y := []byte{0xd9, 0xd9, 0xf7, 0x00}
	z := []byte{0xd9, 0xd9, 0xf1, 0x00, 0x00}

	if !isValidCBOR(x) {
		t.Error("x should be valid")
	}

	if !isValidCBOR(y) {
		t.Error("y should be valid")
	}

	if isValidCBOR(z) {
		t.Error("z should not be valid")
	}
}



func isValidCBOR(data []byte) bool {
	if len(data) < 3 {
		return false
	}
	cborTag := []byte{0xd9, 0xd9, 0xf7}
	return bytes.Equal(data[0:3], cborTag)
}
