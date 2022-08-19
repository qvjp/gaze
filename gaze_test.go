package main

import "testing"

func TestIBytes(t *testing.T) {
	tests := []struct {
		num  string
		want string
	}{
		{"0x0", "0 B"},
		{"0x8", "8 B"},
		{"0x10", "16 B"},
		{"10", "10 B"},
		{"0x1000", "4 KiB"},
		{"0xC0000000", "3 GiB"},
		{"0xFFFFFFFF", "3 GiB 1023 MiB 1023 KiB 1023 B"},
		{"0x100000000", "4 GiB"},
		{"0xFFFFFFFFFFFFFFFF", "15 EiB 1023 PiB 1023 TiB 1023 GiB 1023 MiB 1023 KiB 1023 B"},
		{"0x10000000000000000", "16 EiB"},
		{"0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", "1152921504606846976 YiB"},
		{"0x100000000000000000000000000000000000", "too big"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IBytes(tt.num); got != tt.want {
				t.Errorf("IBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
