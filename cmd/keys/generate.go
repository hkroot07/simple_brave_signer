package keys

import "github.com/spf13/cobra"

func init() {
	keysCmd.AddCommand(keyGenerateCmd)
	keyGenerateCmd.Flags().String("pub-out", "pub_key.pem", "Path to save the public key")
	keyGenerateCmd.Flags().String("priv-out", "priv_key.pem", "Path to save the private key")
	keyGenerateCmd.Flags().Int("pub-size", 2048, "Private key size in bits")
}

var keyGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates key pair.",
	Long:  `Generate an RSA key pair and store it in PEM files. The private key will be encrypted using a passphrase that you'll need to enter. AES encryption with Argon2 key derivation function is utilized.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
