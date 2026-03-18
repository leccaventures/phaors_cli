package app

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const cliInvocation = "go run ./cmd/pharoscli"

func printUsage(contractABI abi.ABI) {
	fmt.Print(usageText(contractABI))
}

func printMethods(contractABI abi.ABI) {
	fmt.Print(methodListText(readMethods(contractABI)))
}

func printMethodHelp(contractABI abi.ABI, methodName string) error {
	method, err := lookupReadMethod(contractABI, methodName)
	if err != nil {
		return err
	}

	fmt.Print(methodHelpText(method))
	return nil
}

func usageText(contractABI abi.ABI) string {
	methods := readMethods(contractABI)
	var b strings.Builder
	b.WriteString("Pharos staking read CLI\n\n")
	b.WriteString("Usage:\n\n")
	b.WriteString("        " + cliInvocation + " <command> [arguments]\n\n")
	b.WriteString("The commands are:\n\n")
	b.WriteString("        methods     list readable contract methods\n")
	b.WriteString("        help        show top-level help or method help\n")
	b.WriteString("        <method>    call a readable method directly\n")
	b.WriteString("        call        call a readable method explicitly\n\n")
	b.WriteString("Global flags:\n\n")
	b.WriteString(fmt.Sprintf("        --rpc URL            RPC endpoint (default %s)\n", defaultRPCURL))
	b.WriteString(fmt.Sprintf("        --contract ADDRESS   contract address (default %s)\n", defaultContractHex))
	b.WriteString(fmt.Sprintf("        --timeout DURATION   request timeout (default %s)\n", defaultTimeout))
	b.WriteString("        -h, -help, --help    show top-level help\n\n")
	b.WriteString("Use \"" + cliInvocation + " help <method>\" for more information about a method.\n")
	b.WriteString("Use \"" + cliInvocation + " methods\" to see every readable method.\n\n")
	b.WriteString("Examples:\n\n")
	b.WriteString("        " + cliInvocation + " methods\n")
	b.WriteString("        " + cliInvocation + " help getValidator\n")
	b.WriteString("        " + cliInvocation + " currentEpoch\n")
	b.WriteString("        " + cliInvocation + " getDelegator _poolId=0x1234... _delegator=0xabc...\n\n")
	b.WriteString(fmt.Sprintf("Readable methods available: %d\n", len(methods)))
	return b.String()
}

func methodHelpText(method abi.Method) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Method: %s\n", method.Name))
	if description := methodDescription(method.Name); description != "" {
		b.WriteString(fmt.Sprintf("Description: %s\n", description))
	}
	b.WriteString(fmt.Sprintf("Signature: %s\n", describeMethod(method)))
	b.WriteString(fmt.Sprintf("Usage: %s %s%s\n", cliInvocation, method.Name, renderPositionalSuffix(method)))
	b.WriteString(fmt.Sprintf("   or: %s %s%s\n", cliInvocation, method.Name, renderNamedSuffix(method)))
	b.WriteString("\n")
	b.WriteString("Inputs:\n")
	if len(method.Inputs) == 0 {
		b.WriteString("  none\n")
	} else {
		for idx, input := range method.Inputs {
			b.WriteString(fmt.Sprintf("  %s (%s) example=%s\n", inputName(input, idx), input.Type.String(), exampleValue(input.Type.String(), idx)))
		}
	}
	b.WriteString("\n")
	b.WriteString("Outputs:\n")
	if len(method.Outputs) == 0 {
		b.WriteString("  none\n")
	} else {
		for idx, output := range method.Outputs {
			b.WriteString(fmt.Sprintf("  %s (%s)\n", outputName(output, idx), output.Type.String()))
		}
	}
	b.WriteString("\n")
	b.WriteString("Notes:\n")
	b.WriteString("  positional args follow ABI order\n")
	b.WriteString("  named args use key=value\n")
	b.WriteString("  arrays accept JSON or comma-separated values\n")
	return b.String()
}

func methodListText(methods []abi.Method) string {
	var b strings.Builder
	b.WriteString("Readable contract methods\n\n")
	b.WriteString("Arguments shown here are positional. Use \"" + cliInvocation + " help <method>\" for types, named args, and examples.\n\n")

	sections := []struct {
		title string
		keep  func(abi.Method) bool
	}{
		{title: "No-argument methods", keep: func(m abi.Method) bool { return len(m.Inputs) == 0 }},
		{title: "Single-argument methods", keep: func(m abi.Method) bool { return len(m.Inputs) == 1 }},
		{title: "Two-argument methods", keep: func(m abi.Method) bool { return len(m.Inputs) == 2 }},
		{title: "Three-or-more-argument methods", keep: func(m abi.Method) bool { return len(m.Inputs) >= 3 }},
	}

	for _, section := range sections {
		count := 0
		width := 0
		argWidth := 0
		for _, method := range methods {
			if !section.keep(method) {
				continue
			}
			count++
			if len(method.Name) > width {
				width = len(method.Name)
			}
			if len(compactPositionalList(method)) > argWidth {
				argWidth = len(compactPositionalList(method))
			}
		}
		if count == 0 {
			continue
		}

		b.WriteString(section.title)
		b.WriteString("\n")
		for _, method := range methods {
			if !section.keep(method) {
				continue
			}
			description := methodDescription(method.Name)
			if description == "" {
				b.WriteString(fmt.Sprintf("  %-*s  %s\n", width, method.Name, compactPositionalList(method)))
				continue
			}
			b.WriteString(fmt.Sprintf("  %-*s  %-*s  %s\n", width, method.Name, argWidth, compactPositionalList(method), description))
		}
		b.WriteString("\n")
	}

	b.WriteString(fmt.Sprintf("Total readable methods: %d\n", len(methods)))
	return b.String()
}

func lookupReadMethod(contractABI abi.ABI, methodName string) (abi.Method, error) {
	method, ok := contractABI.Methods[methodName]
	if !ok {
		return abi.Method{}, unknownReadMethodError(contractABI, methodName)
	}
	if method.StateMutability != "view" && method.StateMutability != "pure" {
		return abi.Method{}, fmt.Errorf("method %q exists but is not readable\n\nTry:\n  %s methods", methodName, cliInvocation)
	}
	return method, nil
}

func unknownReadMethodError(contractABI abi.ABI, methodName string) error {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("unknown read method %q\n", methodName))
	if suggestion, ok := suggestedReadMethod(readMethods(contractABI), methodName); ok {
		b.WriteString("\nDid you mean:\n")
		b.WriteString(fmt.Sprintf("  %s %s\n", cliInvocation, suggestion))
		b.WriteString(fmt.Sprintf("  %s help %s\n", cliInvocation, suggestion))
	}
	b.WriteString("\nTry:\n")
	b.WriteString("  " + cliInvocation + " methods\n")
	b.WriteString("  " + cliInvocation + " help <method>")
	return fmt.Errorf("%s", b.String())
}

func suggestedReadMethod(methods []abi.Method, input string) (string, bool) {
	query := strings.ToLower(strings.TrimSpace(input))
	if query == "" {
		return "", false
	}

	bestName := ""
	bestScore := 1 << 30
	for _, method := range methods {
		name := method.Name
		lowerName := strings.ToLower(name)
		score := levenshteinDistance(query, lowerName)
		if strings.HasPrefix(lowerName, query) || strings.Contains(lowerName, query) {
			score = 0
		}
		if score < bestScore || (score == bestScore && name < bestName) {
			bestName = name
			bestScore = score
		}
	}

	threshold := 3
	if len(query) >= 8 {
		threshold = 4
	}
	if bestName == "" || bestScore > threshold {
		return "", false
	}
	return bestName, true
}

func levenshteinDistance(a, b string) int {
	if a == b {
		return 0
	}
	if a == "" {
		return len(b)
	}
	if b == "" {
		return len(a)
	}

	prev := make([]int, len(b)+1)
	for j := range prev {
		prev[j] = j
	}

	for i := 1; i <= len(a); i++ {
		curr := make([]int, len(b)+1)
		curr[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			curr[j] = minInt(
				prev[j]+1,
				curr[j-1]+1,
				prev[j-1]+cost,
			)
		}
		prev = curr
	}

	return prev[len(b)]
}

func minInt(values ...int) int {
	best := values[0]
	for _, value := range values[1:] {
		if value < best {
			best = value
		}
	}
	return best
}

func readMethods(contractABI abi.ABI) []abi.Method {
	methods := make([]abi.Method, 0)
	for _, method := range contractABI.Methods {
		if method.StateMutability == "view" || method.StateMutability == "pure" {
			methods = append(methods, method)
		}
	}
	sort.Slice(methods, func(i, j int) bool {
		return methods[i].Name < methods[j].Name
	})
	return methods
}

func describeMethod(method abi.Method) string {
	inputs := make([]string, 0, len(method.Inputs))
	for idx, input := range method.Inputs {
		inputs = append(inputs, fmt.Sprintf("%s:%s", inputName(input, idx), input.Type.String()))
	}
	outputs := make([]string, 0, len(method.Outputs))
	for idx, output := range method.Outputs {
		outputs = append(outputs, fmt.Sprintf("%s:%s", outputName(output, idx), output.Type.String()))
	}
	return fmt.Sprintf("%s(%s) -> (%s) [%s]", method.Name, strings.Join(inputs, ", "), strings.Join(outputs, ", "), method.StateMutability)
}

func compactPositionalList(method abi.Method) string {
	if len(method.Inputs) == 0 {
		return "()"
	}
	parts := make([]string, 0, len(method.Inputs))
	for idx, input := range method.Inputs {
		name := inputName(input, idx)
		parts = append(parts, fmt.Sprintf("[%s]", name))
	}
	return strings.Join(parts, " ")
}

func inputName(input abi.Argument, idx int) string {
	if input.Name != "" {
		return input.Name
	}
	return fmt.Sprintf("arg%d", idx)
}

func outputName(output abi.Argument, idx int) string {
	if output.Name != "" {
		return output.Name
	}
	return fmt.Sprintf("ret%d", idx)
}

func renderPositionalSuffix(method abi.Method) string {
	if len(method.Inputs) == 0 {
		return ""
	}
	parts := make([]string, 0, len(method.Inputs))
	for idx, input := range method.Inputs {
		parts = append(parts, exampleValue(input.Type.String(), idx))
	}
	return " " + strings.Join(parts, " ")
}

func renderNamedSuffix(method abi.Method) string {
	if len(method.Inputs) == 0 {
		return ""
	}
	parts := make([]string, 0, len(method.Inputs))
	for idx, input := range method.Inputs {
		parts = append(parts, fmt.Sprintf("%s=%s", inputName(input, idx), exampleValue(input.Type.String(), idx)))
	}
	return " " + strings.Join(parts, " ")
}

func exampleValue(typeName string, idx int) string {
	switch typeName {
	case "address":
		return "0x1111111111111111111111111111111111111111"
	case "bytes32":
		return "0x1111111111111111111111111111111111111111111111111111111111111111"
	case "bytes4":
		return "0x01ffc9a7"
	case "bool":
		return "true"
	case "bytes32[]":
		return fmt.Sprintf("'[\"%s\",\"%s\"]'", exampleValue("bytes32", idx), strings.Replace(exampleValue("bytes32", idx), "1", "2", 1))
	case "address[]":
		return "'[\"0x1111111111111111111111111111111111111111\",\"0x2222222222222222222222222222222222222222\"]'"
	}
	if strings.HasPrefix(typeName, "uint") {
		return fmt.Sprintf("%d", idx+1)
	}
	return "value"
}
