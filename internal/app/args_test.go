package app

import "testing"

func TestParseMethodArgsSupportsPositionalInput(t *testing.T) {
	contractABI, err := loadABI()
	if err != nil {
		t.Fatalf("load ABI: %v", err)
	}

	method, err := lookupReadMethod(contractABI, "commissionRates")
	if err != nil {
		t.Fatalf("lookup read method: %v", err)
	}

	args, err := parseMethodArgs(method, []string{"0x1111111111111111111111111111111111111111111111111111111111111111"})
	if err != nil {
		t.Fatalf("parse positional args: %v", err)
	}
	if len(args) != 1 {
		t.Fatalf("expected 1 parsed arg, got %d", len(args))
	}
}
