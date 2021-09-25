package lump

import (
	"fmt"
	"smg/utils"
)

func Installphp() {

	utils.Exec("sudo", "apt-get", "install", "software-properties-common")
	utils.Exec("sudo", "add-apt-repository", "ppa:ondrej/php")
	utils.Exec("sudo", "apt-get", "update")
	utils.Exec("sudo", "apt-get", "-y", "php7.4")

}

func Php74() {

	utils.Exec("sudo", "apt-get", "-y", "php7.4")

	fmt.Println("Installing php 7.4 ...")

	Addons("7.4")

}

func Php73() {

	utils.Exec("sudo", "apt-get", "-y", "php7.3")

	fmt.Println("Installing php 7.3 ...")

	Addons("7.3")

}

func Php72() {

	utils.Exec("sudo", "apt-get", "-y", "php7.2")

	fmt.Println("Installing php 7.2 ...")

	Addons("7.2")

}

func Addons(version string) {

	extensions := []string{
		"apt",
		"install",
		"php" + version + "-common",
		"php" + version + "-mysql",
		"php" + version + "-xml",
		"php" + version + "-xmlrpc",
		"php" + version + "-curl",
		"php" + version + "-gd",
		"php" + version + "-imagick",
		"php" + version + "-cli",
		"php" + version + "-dev",
		"php" + version + "-imap",
		"php" + version + "-mbstring",
		"php" + version + "-opcache",
		"php" + version + "-soap",
		"php" + version + "-zip",
		"php" + version + "-intl",
		"php" + version + "-fpm",
	}

	utils.Exec("sudo", extensions...)

	fmt.Println("Php Extensions Installation Complete.")

}
