package app

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func marshalInputs(method abi.Method, values []any) map[string]any {
	result := make(map[string]any, len(values))
	for i, value := range values {
		name := method.Inputs[i].Name
		if name == "" {
			name = fmt.Sprintf("arg%d", i)
		}
		result[name] = normalizeValue(reflect.ValueOf(value))
	}
	return result
}

func marshalOutputs(method abi.Method, values []any) map[string]any {
	result := make(map[string]any, len(values))
	for i, value := range values {
		name := method.Outputs[i].Name
		if name == "" {
			name = fmt.Sprintf("ret%d", i)
		}
		result[name] = normalizeValue(reflect.ValueOf(value))
	}
	return result
}

func normalizeValue(v reflect.Value) any {
	if !v.IsValid() {
		return nil
	}

	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil
		}
		if bi, ok := v.Interface().(*big.Int); ok {
			return bi.String()
		}
		return normalizeValue(v.Elem())
	}

	if v.CanInterface() {
		switch value := v.Interface().(type) {
		case common.Address:
			return value.Hex()
		case [32]byte:
			return "0x" + hex.EncodeToString(value[:])
		case [4]byte:
			return "0x" + hex.EncodeToString(value[:])
		}
	}

	switch v.Kind() {
	case reflect.Struct:
		result := make(map[string]any, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			if !field.IsExported() {
				continue
			}
			result[lowerCamel(field.Name)] = normalizeValue(v.Field(i))
		}
		return result
	case reflect.Slice, reflect.Array:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			bytes := make([]byte, v.Len())
			reflect.Copy(reflect.ValueOf(bytes), v)
			return "0x" + hex.EncodeToString(bytes)
		}
		result := make([]any, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			result = append(result, normalizeValue(v.Index(i)))
		}
		return result
	case reflect.Map:
		result := make(map[string]any, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			result[fmt.Sprint(iter.Key().Interface())] = normalizeValue(iter.Value())
		}
		return result
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	default:
		if v.CanInterface() {
			return v.Interface()
		}
		return fmt.Sprint(v)
	}
}

func lowerCamel(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
