package utils

import (
	"strings"
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func ToBase62(data []byte) string {
	var val uint64
	for _, b := range data {
		val = val*256 + uint64(b) // Convert each byte into the number space
	}

	var result strings.Builder
	for val > 0 {
		result.WriteByte(characterSet[val%base])
		val /= base
	}

	return reverseString(result.String())
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
