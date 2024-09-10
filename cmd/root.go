package cmd

import (
	"context"
	"simple_brave_signer/config"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "simple_brave_signer",
		Short: "Bravely generate key pairs, sign file and check signatures.",
		Long:  `A collection of tools to generate key pairs in PEM files, sign files and verify signatures.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}
}

func initializeConfig(cmd *cobra.Command) error {
	localViper, err := config.LoadYamlConfig()
	if err != nil {
		return err
	}

	ctx := context.WithValue(cmd.Context(), config.ViperKey, localViper)
	cmd.SetContext(ctx)
	return nil
}
