package lump

import (
	"fmt"
	"smg/utils"
)

func InstallMysql() {

	status, resp := utils.Exec("sudo", "apt-get", "install", "mysql-server", "-y")

	if status == true {
		fmt.Println(resp, "\n", " Installation was a success!")
		return
	}

	fmt.Println("Error :- ", resp, "\n", "Installation Failed.")

	mysqlCliInstall()

}

func mysqlCliInstall() {

	utils.Exec("sudo", "mysql_secure_installation")

}
