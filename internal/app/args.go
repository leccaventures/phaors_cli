package app

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func parseMethodArgs(method abi.Method, rawArgs []string) ([]any, error) {
	if len(rawArgs) != len(method.Inputs) {
		hasNamed := true
		for _, arg := range rawArgs {
			if !strings.Contains(arg, "=") {
				hasNamed = false
				break
			}
		}
		if !hasNamed {
			return nil, fmt.Errorf("method %s expects %d argument(s), got %d", method.Name, len(method.Inputs), len(rawArgs))
		}
	}

	values := make([]any, len(method.Inputs))
	if len(rawArgs) == 0 && len(method.Inputs) == 0 {
		return values, nil
	}

	useNamed := true
	for _, arg := range rawArgs {
		if !strings.Contains(arg, "=") {
			useNamed = false
			break
		}
	}

	if useNamed {
		rawByName := make(map[string]string, len(rawArgs))
		for _, arg := range rawArgs {
			parts := strings.SplitN(arg, "=", 2)
			rawByName[parts[0]] = parts[1]
		}
		for i, input := range method.Inputs {
			raw, ok := rawByName[input.Name]
			if !ok {
				return nil, fmt.Errorf("missing named argument %q", input.Name)
			}
			value, err := parseValue(input.Type, raw)
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", input.Name, err)
			}
			values[i] = value
		}
		return values, nil
	}

	if len(rawArgs) != len(method.Inputs) {
		return nil, fmt.Errorf("method %s expects %d argument(s), got %d", method.Name, len(method.Inputs), len(rawArgs))
	}

	for i, input := range method.Inputs {
		value, err := parseValue(input.Type, rawArgs[i])
		if err != nil {
			name := input.Name
			if name == "" {
				name = fmt.Sprintf("arg%d", i)
			}
			return nil, fmt.Errorf("parse %s: %w", name, err)
		}
		values[i] = value
	}
	return values, nil
}

func parseValue(t abi.Type, raw string) (any, error) {
	if strings.HasSuffix(t.String(), "[]") {
		return parseSliceValue(t, raw)
	}

	switch t.String() {
	case "address":
		if !common.IsHexAddress(raw) {
			return nil, fmt.Errorf("invalid address %q", raw)
		}
		return common.HexToAddress(raw), nil
	case "bool":
		return strconv.ParseBool(raw)
	case "bytes32":
		return parseFixedBytes32(raw)
	case "bytes4":
		return parseFixedBytes4(raw)
	}

	if strings.HasPrefix(t.String(), "uint") {
		return parseBigInt(raw)
	}

	return nil, fmt.Errorf("unsupported input type %s", t.String())
}

func parseSliceValue(t abi.Type, raw string) (any, error) {
	var items []string
	trimmed := strings.TrimSpace(raw)
	if strings.HasPrefix(trimmed, "[") {
		if err := json.Unmarshal([]byte(trimmed), &items); err != nil {
			return nil, fmt.Errorf("invalid JSON array: %w", err)
		}
	} else if trimmed == "" {
		items = []string{}
	} else {
		items = strings.Split(trimmed, ",")
	}

	values := reflect.MakeSlice(reflect.SliceOf(goTypeForABI(*t.Elem)), 0, len(items))
	for _, item := range items {
		parsed, err := parseValue(*t.Elem, strings.TrimSpace(item))
		if err != nil {
			return nil, err
		}
		values = reflect.Append(values, reflect.ValueOf(parsed))
	}
	return values.Interface(), nil
}

func goTypeForABI(t abi.Type) reflect.Type {
	switch t.String() {
	case "address":
		return reflect.TypeOf(common.Address{})
	case "bool":
		return reflect.TypeOf(false)
	case "bytes32":
		return reflect.TypeOf([32]byte{})
	case "bytes4":
		return reflect.TypeOf([4]byte{})
	}
	if strings.HasPrefix(t.String(), "uint") {
		return reflect.TypeOf(&big.Int{})
	}
	if strings.HasSuffix(t.String(), "[]") && t.Elem != nil {
		return reflect.SliceOf(goTypeForABI(*t.Elem))
	}
	panic("unsupported ABI type: " + t.String())
}

func parseBigInt(raw string) (*big.Int, error) {
	value := new(big.Int)
	base := 10
	input := raw
	if strings.HasPrefix(raw, "0x") || strings.HasPrefix(raw, "0X") {
		base = 16
		input = raw[2:]
	}
	if _, ok := value.SetString(input, base); !ok {
		return nil, fmt.Errorf("invalid integer %q", raw)
	}
	return value, nil
}

func parseFixedBytes32(raw string) ([32]byte, error) {
	var out [32]byte
	bytes, err := parseHexBytes(raw, 32)
	if err != nil {
		return out, err
	}
	copy(out[:], bytes)
	return out, nil
}

func parseFixedBytes4(raw string) ([4]byte, error) {
	var out [4]byte
	bytes, err := parseHexBytes(raw, 4)
	if err != nil {
		return out, err
	}
	copy(out[:], bytes)
	return out, nil
}

func parseHexBytes(raw string, expectedLen int) ([]byte, error) {
	trimmed := strings.TrimPrefix(strings.TrimPrefix(raw, "0x"), "0X")
	if len(trimmed) != expectedLen*2 {
		return nil, fmt.Errorf("expected %d-byte hex string", expectedLen)
	}
	decoded, err := hex.DecodeString(trimmed)
	if err != nil {
		return nil, fmt.Errorf("invalid hex bytes: %w", err)
	}
	return decoded, nil
}
