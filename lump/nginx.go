package lump

import (
	"bytes"
	"fmt"
	"smg/utils"
	"text/template"
)

type Ubuntu struct {
	CodeName string
}

var (
	nginxSourceListLocation string = "/etc/apt/sources.list.d/nginx.list"

	nginxSourceList string = `
	## Added by github.com/4nkitd/smg 
	deb https://nginx.org/packages/ubuntu/ {{ .CodeName }} nginx
	deb-src https://nginx.org/packages/ubuntu/ {{ .CodeName }} nginx`

	parsedNginxSourceData bytes.Buffer
)

func UpdateNginxSourceList() {

	ubuntu := Ubuntu{CodeName: utils.GetUbuntuRealseCodeName()}

	_parser, _ := template.New("Ubuntu").Parse(nginxSourceList)

	_parser.Execute(&parsedNginxSourceData, ubuntu)

	utils.WriteToFile(nginxSourceListLocation, parsedNginxSourceData.String())

	fmt.Println(nginxSourceListLocation, "Update")
}

func InstallNginx() {

	UpdateNginxSourceList()

	status, resp := utils.Exec("sudo", "apt", "install", "nginx")

	if status == true {
		fmt.Println(resp, "\n", " Installation was a success!")
		return
	}

	fmt.Println("Error :- ", resp, "\n", "Installation Failed.")

}
