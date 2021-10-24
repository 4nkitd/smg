package cmd

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts string/file",
	Long: `
Decrypt as the name suggest allows you to decrypt string or file.
You need to specify the algorithm and key to decrypt string or file,
by passing the appropriate flag to the command
`,
	Run: func(cmd *cobra.Command, args []string) {
		if algorythmName == "aes" {
			// Use wants to use the AES encryption
			if encryptionKey == "" {
				// * User didn't provide the encryption key so we generate passphrase
				if passphrase == "" {
					// * Generate the key using sha256
					hash := sha256.Sum256([]byte(passphrase))
					encryptionKey = hex.EncodeToString(hash[:])
				} else {
					panic("Please provide key or a passphrase")
				}
			}
			if plaintext != "" {
				// Plaintext isn't empty so encrypt it
				fmt.Println("Decrypted, hex:")
				plaintext2, err := hex.DecodeString(plaintext)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(decryptString(plaintext2)))
			} else {
				// Check if user provided files
				if inFile != "" && outFile != "" {
					fmt.Println("File mode")
					ciphertext := decryptString(readFileIn(inFile))
					writeFileOut(outFile, ciphertext)
				} else {
					panic("Please use input file and output file")
				}
			}
		} else if algorythmName == "rsa" {
			// * Checks if user provided the private RSA key
			if encryptionKey != "" {
				decryptString(nil)
			} else {
				// Panic because user didn't provide the encryption key
				panic("Please provide a private key to encrypt message")
			}
		} else {
			// User entered wrong algorythm
			panic("algorithm name not recognized or null")
		}
	},
}

func Decrypt() *cobra.Command {
	return decryptCmd
}

func init() {

	// Here you will define your flags and configuration settings.
	decryptCmd.Flags().StringVarP(&algorythmName, "algorythm", "a", "aes", "Algorithm to use when decrypting")
	decryptCmd.Flags().StringVarP(&inFile, "inputFile", "f", "", "File input to use when decrypting")
	decryptCmd.Flags().StringVarP(&outFile, "outputFile", "o", "", "File output to use when decrypting")
	decryptCmd.Flags().StringVarP(&plaintext, "data", "d", "", "Data to decrypt")
	decryptCmd.Flags().StringVarP(&passphrase, "passphrase", "p", "", "Passphrase to use when decrypting")
	decryptCmd.Flags().StringVar(&encryptionKey, "key", "", "Key to use when decrypting")
}

func decryptString(ciphertext []byte) []byte {
	// Decrypts the byte array and returns plaintext
	// TODO: Maybe return hex string to simplify the code
	switch algorythmName {
	case "aes":
		key, err := hex.DecodeString(encryptionKey)
		if err != nil {
			panic(err)
		}
		c, err := aes.NewCipher(key)
		if err != nil {
			panic(err)
		}
		gcmDecrypt, err := cipher.NewGCM(c)
		if err != nil {
			panic(err)
		}
		nonceSize := gcmDecrypt.NonceSize()
		if len(ciphertext) < nonceSize {
			panic(err)
		}
		nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]
		plaintext, err := gcmDecrypt.Open(nil, []byte(nonce), []byte(encryptedMessage), nil)
		if err != nil {
			panic(err)
		}
		return plaintext
	case "rsa":
		f, err := ioutil.ReadFile(encryptionKey)
		if err != nil {
			panic(err)
		}
		privateKey, err := ParseRsaPrivateKeyFromPemStr(string(f))
		if err != nil {
			// User entered wrong key
			panic(err)
		}
		if inFile != "" && outFile != "" {
			// * Read from file and load result to the out file
			inFileRaw, err := ioutil.ReadFile(inFile)
			if err != nil {
				panic(err)
			}
			decryptedBytes, err := privateKey.Decrypt(nil, inFileRaw, &rsa.OAEPOptions{Hash: crypto.SHA256})
			if err != nil {
				panic(err)
			}
			writeFileOut(outFile, decryptedBytes)
		} else if plaintext != "" {
			// User provided the plaintext string so decrypt it using the private key
			decryptedBytes, err := privateKey.Decrypt(nil, []byte(plaintext), &rsa.OAEPOptions{Hash: crypto.SHA256})
			if err != nil {
				panic(err)
			}
			fmt.Println(string(decryptedBytes))
		} else {
			// User entered wrong flags
			panic("Please use apropriate flags")
		}
		return nil
	default:
		panic("Cannot find algorythm")
	}
}
