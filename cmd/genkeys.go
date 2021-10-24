package cmd

import (
	"fmt"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/spf13/cobra"
)

// genkeyCmd represents the genkey command
var genkeysCmd = &cobra.Command{
	Use:   "genkeys",
	Short: "Generates RSA key pair",
	Long: `
This command generated new RSA key pair.
It's that simple`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating Keys...")
		priv, pub := GenerateRsaKeyPair()
		// Export the keys to pem string
		priv_pem := ExportRsaPrivateKeyAsPemStr(priv)
		pub_pem, _ := ExportRsaPublicKeyAsPemStr(pub)
		writeFileOut("rsa_key", []byte(priv_pem))
		writeFileOut("rsa_key.pub", []byte(pub_pem))
	},
}

func GenKey() *cobra.Command {
	return genkeysCmd
}

func init() {
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)
	return string(pubkey_pem), nil
}
