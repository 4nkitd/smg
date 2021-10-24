package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/spf13/cobra"
)

var algorythm string
var filePath string
var dataString string
var outputFile string

var inputFile string
var hashSumFile string

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash allows you to create or verify hashes",
	Long: `
Hash command, makes hashing files easier and verifying hashes of files.
Just pass a flag and have fun.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use an appropriate command")
	},
}

func Hash() *cobra.Command {
	return hashCmd
}

func Verify() *cobra.Command {
	return verifyCmd
}

func Gen() *cobra.Command {
	return genCmd
}

func init() {

	// Here you define your flags and configuration settings.
	hashCmd.PersistentFlags().StringVarP(&algorythm, "algorythm", "a", "sha256", "Algorythm to use when hashing")

	genCmd.Flags().StringVarP(&dataString, "data", "d", "", "Data to be hashed")
	genCmd.Flags().StringVarP(&filePath, "file", "f", "", "File to be hashed")
	genCmd.Flags().StringVarP(&outputFile, "out", "o", "", "File to save the hashed output")

	verifyCmd.Flags().StringVarP(&inputFile, "file", "f", "", "File to be verified")
	verifyCmd.Flags().StringVarP(&hashSumFile, "hashSum", "s", "", "Hash sum to be checked")
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate the hash from string/file",
	Long: `
This subcommand generates hashes from strings and files.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating Hash...")
		if dataString != "" {
			// * Check if dataString is null, if not hash the data
			hash := hashString(dataString)
			if hash != "" {
				fmt.Println("Hashed String: ")
				fmt.Println(hash)
			} else {
				panic("Cannot find algorythm")
			}
		} else {
			// * Search if the file was specified exists
			if filePath != "" {
				if outputFile != "" {
					// * File output exists, hash file and create new file with name of output file
					fmt.Println("Writing file1")
					writeFile(readFile(filePath), outputFile)
				} else {
					// * File output doesnt exists, create new file with default name
					fileNameSplit := strings.Split(filePath, ".")
					fileName := fileNameSplit[0]
					fmt.Println("Writing file2")
					writeFile(readFile((filePath)), fileName+".hash")
				}
			} else {
				panic("User didn't specify string or file input")
			}
		}
	},
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the hash from string/file",
	Long: `
This subcommand verifies file hash sums.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Verifing Hash...")
		if inputFile != "" && hashSumFile != "" && algorythm != "" {
			// * First Calculate hash of the inputFile
			fileHash := readFile(inputFile)
			hash := readFileRaw(hashSumFile)
			if fileHash == hash {
				fmt.Println("Hash verified successfully")
			} else {
				fmt.Println("Hash failed verification")
				fmt.Println("Provided hash: " + hash)
				fmt.Println("Calculated hash: " + fileHash)
				fmt.Println("Algorythm: " + algorythm)
			}
		} else {
			panic("Please enter input file, hash sum file and algorythm")
		}
	},
}

func hashString(data string) string {
	// Switch the algorythm and hash the data
	switch algorythm {
	case "sha256":
		hashedString := sha256.Sum256([]byte(data))
		return hex.EncodeToString(hashedString[:])
	case "sha512":
		hashedString := sha512.Sum512([]byte(data))
		return hex.EncodeToString(hashedString[:])
	case "md5":
		hashedString := md5.Sum([]byte(data))
		return hex.EncodeToString(hashedString[:])
	default:
		return ""
	}
}

func readFile(FilePath string) string {
	// Reads the file and returns the hash string
	f, err := os.Open(FilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	switch algorythm {
	case "sha256":
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			panic(err)
		}
		return hex.EncodeToString(h.Sum(nil))
	case "sha512":
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			panic(err)
		}
		return hex.EncodeToString(h.Sum(nil))
	case "md5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			panic(err)
		}
		return hex.EncodeToString(h.Sum(nil))
	default:
		panic("Cannot find algorythm")
	}
}

func writeFile(Data string, Name string) {
	// Writes data to the filename name
	if Data != "" {
		if err := os.WriteFile(Name, []byte(Data), 0666); err != nil {
			panic(err)
		}
	} else {
		panic("Cannot write null to file!")
	}
}

func readFileRaw(FilePath string) string {
	// Reads file and returns raw contents of the file
	content, err := ioutil.ReadFile(FilePath)
	if err != nil {
		panic(err)
	}
	text := string(content)
	return text
}
