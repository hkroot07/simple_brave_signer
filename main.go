package main

import (
	"simple_brave_signer/cmd"
	"simple_brave_signer/cmd/keys"
	"simple_brave_signer/cmd/signatures"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)
}
