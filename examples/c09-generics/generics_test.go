package main

import "testing"

func TestCacheGet(t *testing.T) {
	ports := newCache[string, int]()
	ports.set("http", 80)
	ports.set("ssh", 22)

	tests := []struct {
		name   string
		key    string
		want   int
		wantOk bool
	}{
		{"present", "http", 80, true},
		{"missing", "smtp", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := ports.get(test.key)
			if got != test.want || ok != test.wantOk {
				t.Errorf("get(%q) = %v, %v, want %v, %v", test.key, got, test.want, test.wantOk, ok)
			}
		})
	}
}
