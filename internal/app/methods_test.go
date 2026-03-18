package app

import (
	"strings"
	"testing"
)

func TestUsageText(t *testing.T) {
	contractABI, err := loadABI()
	if err != nil {
		t.Fatalf("load ABI: %v", err)
	}

	got := usageText(contractABI)

	checks := []string{
		"Pharos staking read CLI",
		"Usage:",
		cliInvocation + " <command> [arguments]",
		"methods     list readable contract methods",
		"help        show top-level help or method help",
		"<method>    call a readable method directly",
		"call        call a readable method explicitly",
		"Use \"" + cliInvocation + " help <method>\" for more information about a method.",
		"Use \"" + cliInvocation + " methods\" to see every readable method.",
		"Readable methods available:",
	}

	for _, check := range checks {
		if !strings.Contains(got, check) {
			t.Fatalf("usage text missing %q\n%s", check, got)
		}
	}

	if strings.Contains(got, "Readable methods with no inputs") {
		t.Fatalf("usage text should not include the full grouped method catalog\n%s", got)
	}
	if strings.Contains(got, cliInvocation+" --help") {
		t.Fatalf("usage text should describe the help flag once, not as a command\n%s", got)
	}
}

func TestMethodHelpText(t *testing.T) {
	contractABI, err := loadABI()
	if err != nil {
		t.Fatalf("load ABI: %v", err)
	}

	method, err := lookupReadMethod(contractABI, "getDelegator")
	if err != nil {
		t.Fatalf("lookup read method: %v", err)
	}

	got := methodHelpText(method)

	checks := []string{
		"Method: getDelegator",
		"Description: Retrieves delegator information for a specific pool.",
		"Signature: getDelegator(_poolId:bytes32, _delegator:address)",
		"Usage: " + cliInvocation + " getDelegator",
		"or: " + cliInvocation + " getDelegator _poolId=",
		"Inputs:",
		"_poolId (bytes32) example=",
		"_delegator (address) example=",
		"Outputs:",
		"Notes:",
		"named args use key=value",
	}

	for _, check := range checks {
		if !strings.Contains(got, check) {
			t.Fatalf("method help missing %q\n%s", check, got)
		}
	}
}

func TestMethodListText(t *testing.T) {
	contractABI, err := loadABI()
	if err != nil {
		t.Fatalf("load ABI: %v", err)
	}

	got := methodListText(readMethods(contractABI))

	checks := []string{
		"Readable contract methods",
		"Arguments shown here are positional. Use \"" + cliInvocation + " help <method>\" for types, named args, and examples.",
		"No-argument methods",
		"Single-argument methods",
		"Two-argument methods",
		"currentEpoch",
		"Returns the current epoch number.",
		"getValidator",
		"Retrieves validator information for a specific pool ID.",
		"[_poolId]",
		"getDelegator",
		"Retrieves delegator information for a specific pool.",
		"[_poolId] [_delegator]",
		"commissionRates",
		"Directly accesses the commission rate mapping data.",
		"[poolId]",
		"Total readable methods:",
	}

	for _, check := range checks {
		if !strings.Contains(got, check) {
			t.Fatalf("method list missing %q\n%s", check, got)
		}
	}

	if strings.Contains(got, "-> (") {
		t.Fatalf("method list should not include full output signatures\n%s", got)
	}
	if strings.Contains(got, "[view]") || strings.Contains(got, "[pure]") {
		t.Fatalf("method list should not include mutability markers\n%s", got)
	}
	if strings.Contains(got, ":bytes32") || strings.Contains(got, ":address") || strings.Contains(got, ":uint256") {
		t.Fatalf("method list should not advertise type signatures inline\n%s", got)
	}
}

func TestUnknownReadMethodErrorIncludesSuggestion(t *testing.T) {
	contractABI, err := loadABI()
	if err != nil {
		t.Fatalf("load ABI: %v", err)
	}

	err = invokeReadMethod(config{}, contractABI, "currentEpoc", nil)
	if err == nil {
		t.Fatal("expected error for unknown method")
	}

	got := err.Error()
	checks := []string{
		"unknown read method \"currentEpoc\"",
		"Did you mean:",
		cliInvocation + " currentEpoch",
		cliInvocation + " help currentEpoch",
		"Try:",
		cliInvocation + " methods",
		cliInvocation + " help <method>",
	}

	for _, check := range checks {
		if !strings.Contains(got, check) {
			t.Fatalf("unknown-method guidance missing %q\n%s", check, got)
		}
	}
}
