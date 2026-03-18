# Pharos Contracts CLI

Small Go CLI for reading Pharos staking contract methods from the embedded ABI.

## Layout

- `cmd/pharoscli`: CLI entrypoint
- `internal/app`: private application logic
- `internal/app/assets`: embedded ABI and method reference data
- `docs`: supporting markdown documentation
- `test`: exported JSON fixtures generated from live read calls

## Build

```bash
make build
```

This writes the local binary to `bin/pharoscli`.

## Run

```bash
make run ARGS='methods'
make run ARGS='help getValidator'
go run ./cmd/pharoscli currentEpoch
```

## Install

```bash
make install
```

## Release Builds

```bash
make release
```

This writes cross-built binaries to `dist/` for:

- `darwin/amd64`
- `darwin/arm64`
- `linux/amd64`
- `windows/amd64`

## Export Test Fixtures

The fixture export uses these fixed inputs:

```text
poolid=0x8467c9cf1536642e27ed004d13af86753c312d2d7176a0ca48fe806894b573e2
poolid=0x0a2c55a00df40b658738e1417622b69531d11b37c2cfac825a8e9f565a8064eb
address=0x00000B834695138Ffd7E4BF07CB4470c292F4eE4
```

Generate the requested JSON snapshots with:

```bash
make fixtures
```

This creates:

- `test/getAllValidators.json`
- `test/getActiveValidators.json`
- `test/getValidator.json`
- `test/validators.json`
- `test/isValidatorActive.json`
- `test/getActiveValidatorCount.json`
- `test/getValidatorCounts.json`
- `test/getDelegator.json`
- `test/getPendingStakeInfo.json`

The generated calls are equivalent to:

```bash
go run ./cmd/pharoscli getAllValidators
go run ./cmd/pharoscli getActiveValidators
go run ./cmd/pharoscli getValidator ${poolid}
go run ./cmd/pharoscli validators ${poolid}
go run ./cmd/pharoscli isValidatorActive ${poolid}
go run ./cmd/pharoscli getActiveValidatorCount
go run ./cmd/pharoscli getValidatorCounts
go run ./cmd/pharoscli getDelegator ${poolid} ${address}
go run ./cmd/pharoscli getPendingStakeInfo ${poolid} ${address}
```

## Verify

```bash
make test
make fixtures
```
