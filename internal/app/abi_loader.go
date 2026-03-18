package app

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed assets/abi.json
var abiJSON []byte

type abiEnvelope struct {
	ABI json.RawMessage `json:"abi"`
}

func loadABI() (abi.ABI, error) {
	var envelope abiEnvelope
	if err := json.Unmarshal(abiJSON, &envelope); err != nil {
		return abi.ABI{}, fmt.Errorf("decode abi wrapper: %w", err)
	}
	parsed, err := abi.JSON(strings.NewReader(string(envelope.ABI)))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("parse abi: %w", err)
	}
	return parsed, nil
}
