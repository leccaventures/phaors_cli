package app

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	defaultRPCURL      = "https://atlantic-rpc.dplabs-internal.com"
	defaultContractHex = "0x4100000000000000000000000000000000000000"
	defaultTimeout     = 15 * time.Second
)

type config struct {
	rpcURL      string
	contract    common.Address
	timeout     time.Duration
	showHelp    bool
	command     string
	commandArgs []string
}

func parseCLI(args []string) (config, error) {
	cfg := config{
		rpcURL:   defaultRPCURL,
		contract: common.HexToAddress(defaultContractHex),
		timeout:  defaultTimeout,
	}

	for len(args) > 0 {
		arg := args[0]
		if !strings.HasPrefix(arg, "-") {
			break
		}
		args = args[1:]

		switch {
		case arg == "--help" || arg == "-help" || arg == "-h":
			cfg.showHelp = true
		case strings.HasPrefix(arg, "--rpc="):
			cfg.rpcURL = strings.TrimPrefix(arg, "--rpc=")
		case arg == "--rpc":
			if len(args) == 0 {
				return cfg, errors.New("missing value for --rpc")
			}
			cfg.rpcURL = args[0]
			args = args[1:]
		case strings.HasPrefix(arg, "--contract="):
			value := strings.TrimPrefix(arg, "--contract=")
			if !common.IsHexAddress(value) {
				return cfg, fmt.Errorf("invalid contract address: %s", value)
			}
			cfg.contract = common.HexToAddress(value)
		case arg == "--contract":
			if len(args) == 0 {
				return cfg, errors.New("missing value for --contract")
			}
			if !common.IsHexAddress(args[0]) {
				return cfg, fmt.Errorf("invalid contract address: %s", args[0])
			}
			cfg.contract = common.HexToAddress(args[0])
			args = args[1:]
		case strings.HasPrefix(arg, "--timeout="):
			value := strings.TrimPrefix(arg, "--timeout=")
			d, err := time.ParseDuration(value)
			if err != nil {
				return cfg, fmt.Errorf("invalid timeout %q: %w", value, err)
			}
			cfg.timeout = d
		case arg == "--timeout":
			if len(args) == 0 {
				return cfg, errors.New("missing value for --timeout")
			}
			d, err := time.ParseDuration(args[0])
			if err != nil {
				return cfg, fmt.Errorf("invalid timeout %q: %w", args[0], err)
			}
			cfg.timeout = d
			args = args[1:]
		default:
			return cfg, fmt.Errorf("unknown flag: %s", arg)
		}
	}

	if len(args) > 0 {
		cfg.command = args[0]
		cfg.commandArgs = args[1:]
	}

	return cfg, nil
}
