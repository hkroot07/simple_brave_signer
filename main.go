package main

import (
	"simple_brave_signer/cmd"
	"simple_brave_signer/cmd/keys"
	"simple_brave_signer/cmd/signatures"
	"simple_brave_signer/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(err, "Initial setup failed")
	}
}
