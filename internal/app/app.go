package app

import (
	"errors"
)

func Run(args []string) error {
	cfg, err := parseCLI(args)
	if err != nil {
		return err
	}

	contractABI, err := loadABI()
	if err != nil {
		return err
	}

	if cfg.showHelp || cfg.command == "" {
		printUsage(contractABI)
		return nil
	}

	switch cfg.command {
	case "help", "--help", "-h":
		if len(cfg.commandArgs) > 0 {
			return printMethodHelp(contractABI, cfg.commandArgs[0])
		}
		printUsage(contractABI)
		return nil
	case "methods", "list":
		printMethods(contractABI)
		return nil
	case "call":
		if len(cfg.commandArgs) == 0 {
			return errors.New("missing method name after 'call'")
		}
		return invokeReadMethod(cfg, contractABI, cfg.commandArgs[0], cfg.commandArgs[1:])
	default:
		return invokeReadMethod(cfg, contractABI, cfg.command, cfg.commandArgs)
	}
}
