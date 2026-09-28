package hello

import "testing"

func TestHello(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"с именем", "Go", "Привет, Go!"},
		{"пустое имя", "", "Привет, мир!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.in); got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
