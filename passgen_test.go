package passgen

import "testing"

func TestGenerate(t *testing.T) {
	cfg := Config{
		Length: 12,
		UseUpper: true,
		UseLower: true,
		UseNumbers: true,
		UseSymbols: true,
	}

	pass, err := Generate(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pass) != cfg.Length {
		t.Fatalf("unexpected password length %d, got %d", cfg.Length, len(pass))
	}
}
