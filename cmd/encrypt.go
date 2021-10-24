package cmd

import (
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"errors"

	"github.com/spf13/cobra"
)

var algorythmName string
var outFile string
var inFile string
var plaintext string
var encryptionKey string
var passphrase string

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts string/file",
	Long: `
Encrypt as the name suggest allows you to encrypt string or file.
You can use specific algorythms, keys and other cryptographic parameters to your liking.
Change the above by passing apropriate flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if algorythmName == "aes" {
			if encryptionKey == "" {
				// * User didn't provide the encryption key so we generate using passphrase
				if passphrase != "" {
					// * Generate the key using sha256
					hash := sha256.Sum256([]byte(passphrase))
					encryptionKey = hex.EncodeToString(hash[:])
				} else {
					panic("Please provide key or a passphrase")
				}
			}
			if plaintext != "" {
				// * Encrypts the string and encodes it to hex and prints it
				fmt.Println("Encrypted, hex:")
				fmt.Println(hex.EncodeToString(encryptString([]byte(plaintext))))
			} else {
				if inFile != "" && outFile != "" {
					// If user provided input file and output file script will write result to the file
					fmt.Println("File mode")
					ciphertext := encryptString(readFileIn(inFile))
					writeFileOut(outFile, ciphertext)
				} else {
					panic("Please use input file and output file or pass appropriate file")
				}
			}
		} else if algorythmName == "rsa" {
			// Use the RSA encryption
			if encryptionKey != "" {
				// In the case of encryption key not being null read the file and use it as public key
				f, err := ioutil.ReadFile(encryptionKey)
				if err != nil {
					panic(err)
				}
				publicKey, err := ParseRsaPublicKeyFromPemStr(string(f))
				if err != nil {
					panic(err)
				}
				if inFile != "" && outFile != "" {
					// * Read from file and load result to the out file
					inFileRaw, err := ioutil.ReadFile(inFile)
					if err != nil {
						panic(err)
					}
					encryptedBytes, err := rsa.EncryptOAEP(
						sha256.New(),
						rand.Reader,
						publicKey,
						inFileRaw,
						nil)
					if err != nil {
						panic(err)
					}
					writeFileOut(outFile, encryptedBytes)
				} else if plaintext != "" {
					// If plaintext is provided and files haven't been provided encrypt using the encryption key
					encryptedBytes, err := rsa.EncryptOAEP(
						sha256.New(),
						rand.Reader,
						publicKey,
						[]byte(plaintext),
						nil)
					if err != nil {
						panic(err)
					}
					fmt.Println(hex.EncodeToString(encryptedBytes))
				} else {
					// User provided wrong flags
					panic("Please use apropriate flags")
				}
			} else {
				// User didn't provide public key
				panic("Please provide a public key to encrypt message")
			}
		} else {
			// User entered the wrong algorythm name
			panic("algorythm name not recognized")
		}
	},
}

func Encrypt() *cobra.Command {
	return encryptCmd
}

func init() {

	// * Defining flags to use
	encryptCmd.Flags().StringVarP(&algorythmName, "algorythm", "a", "aes", "Algorythm to use when encrypting")
	encryptCmd.Flags().StringVarP(&inFile, "inputFile", "f", "", "File input to use when encrypting")
	encryptCmd.Flags().StringVarP(&outFile, "outputFile", "o", "", "File output to use when encrypting")
	encryptCmd.Flags().StringVarP(&plaintext, "data", "d", "", "Data to encrypt")
	encryptCmd.Flags().StringVarP(&passphrase, "passphrase", "p", "", "Passphrase to use when encrypting")
	encryptCmd.Flags().StringVar(&encryptionKey, "key", "", "Key to use when encrypting")
}

func readFileIn(FilePath string) []byte {
	// Reads file and returns byte array
	f, err := ioutil.ReadFile(FilePath)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFileOut(FileName string, Data []byte) {
	if err := os.WriteFile(FileName, Data, 0666); err != nil {
		panic(err)
	}
}

func encryptString(Data []byte) []byte {
	// Encrypt the byte array and return ciphertext as byte array
	switch algorythmName {
	case "aes":
		key, err := hex.DecodeString(encryptionKey)
		if err != nil {
			panic(err)
		}
		cphr, err := aes.NewCipher(key)
		if err != nil {
			panic(err)
		}
		gcm, err := cipher.NewGCM(cphr)
		if err != nil {
			panic(err)
		}
		nonce := make([]byte, gcm.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err)
		}
		ciphertext := gcm.Seal(nonce, nonce, Data, nil)
		return ciphertext
	default:
		panic("Cannot find algorythm")
	}
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	// Parses the RSA private key from file
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	// Parses the RSA public key from file
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("key type is not rsa")
}
