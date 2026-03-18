package app

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

func invokeReadMethod(cfg config, contractABI abi.ABI, methodName string, rawArgs []string) error {
	method, err := lookupReadMethod(contractABI, methodName)
	if err != nil {
		return err
	}

	parsedArgs, err := parseMethodArgs(method, rawArgs)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()

	client, err := ethclient.DialContext(ctx, cfg.rpcURL)
	if err != nil {
		return fmt.Errorf("dial rpc: %w", err)
	}
	defer client.Close()

	callData, err := contractABI.Pack(methodName, parsedArgs...)
	if err != nil {
		return fmt.Errorf("pack call data: %w", err)
	}

	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &cfg.contract,
		Data: callData,
	}, nil)
	if err != nil {
		return fmt.Errorf("contract call failed: %w", err)
	}

	decoded, err := method.Outputs.UnpackValues(result)
	if err != nil {
		return fmt.Errorf("decode result: %w", err)
	}

	payload := map[string]any{
		"rpc":         cfg.rpcURL,
		"contract":    cfg.contract.Hex(),
		"method":      methodName,
		"signature":   method.Sig,
		"inputs":      marshalInputs(method, parsedArgs),
		"outputs":     marshalOutputs(method, decoded),
		"rawHex":      "0x" + hex.EncodeToString(result),
		"calledAtUTC": time.Now().UTC().Format(time.RFC3339),
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(payload)
}
