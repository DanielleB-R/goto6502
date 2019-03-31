# goto6502

This is an (experimental, in-progress) emulator of the classic 6502
microprocessor in Go.

## Prerequisites

To run this code, you must have the following installed:
- Go (preferably the latest version)
- The `cc65` toolchain from https://cc65.github.io/

## Running the emulator

Currently the emulator only has test code in the main function. Build
the tests by running `make all` in the `asm` subdirectory. Then, run
`go run *.go` in the root directory.
