package util

import (
	"encoding/hex"
	"strings"
)

// HexToUppercase convert 0x${string} to 0x${uppercase string}
func HexToUppercase(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return "0x" + strings.ToUpper(hex[2:])
	}
	return hex
}

// HexToLowercase convert 0x${string} to 0x${lowercase string}
func HexToLowercase(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return "0x" + strings.ToLower(hex[2:])
	}
	return hex
}

// HexIsUppercase check if the hex string is uppercase
func HexIsUppercase(hex string) bool {
	if strings.HasPrefix(hex, "0x") {
		return strings.ToUpper(hex[2:]) == hex[2:]
	}
	return true
}

// HexIsLowercase check if the hex string is lowercase
func HexIsLowercase(hex string) bool {
	if strings.HasPrefix(hex, "0x") {
		return strings.ToLower(hex[2:]) == hex[2:]
	}
	return true
}

// ByteArrayToHex convert byte array to hex string
func ByteArrayToHex(b []byte) string {
	hexString := hex.EncodeToString(b)
	return "0x" + strings.ToUpper(hexString)
}
