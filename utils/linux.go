package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func IsRoot() bool {
	return strings.Contains(GetProcessOwner(), ("root"))
}

func GetProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(stdout)
}

func WriteToFile(file string, data string) {

	err := ioutil.WriteFile(file, []byte(data), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}

func GetUbuntuRealseCodeName() string {

	stdout, err := exec.Command("lsb_release", "-a").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var codeName string
	var codeNameIdentifier string = "Codename:"

	scanner := bufio.NewScanner(strings.NewReader(string(stdout)))
	for scanner.Scan() {

		_line := scanner.Text()

		if strings.Contains(_line, codeNameIdentifier) {

			_line = strings.ReplaceAll(_line, codeNameIdentifier, "")
			_line = strings.Trim(_line, " ")
			codeName = strings.Trim(_line, "	")

		}

	}

	return codeName

}

func UpdateSystem() bool {

	std, err := exec.Command("sudo", "apt-get", "update").Output()

	if err != nil {
		fmt.Println(string(std), err)
		return false
	}

	return true

}

func AddNginxPpaKey() bool {

	std, err := exec.Command("wget", "https://nginx.org/keys/nginx_signing.key", "-O", "-", "|", "sudo", "apt-key", "add", "-").Output()

	if err != nil {
		fmt.Println(string(std), err)
		return false
	}

	return true
}

func Exec(cmd string, args ...string) (bool, string) {

	std, err := exec.Command(cmd, args...).Output()

	if err != nil {
		return false, string(std)
	}

	return true, string(std)
}
