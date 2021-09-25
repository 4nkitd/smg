package lump

import (
	"smg/utils"

	"github.com/spf13/cobra"
)

func InstallerInit(cmd *cobra.Command, args []string) {

	cmd.Annotations = make(map[string]string)

	if !utils.IsRoot() {

		cmd.Annotations["error"] = "This Command Required ROOT Access"
		return

	}

	if utils.SearchString(args, "nginx") == true {

		utils.AddNginxPpaKey()
		utils.UpdateSystem()

		InstallNginx()

	}

	if utils.SearchString(args, "mysql") == true {

		InstallMysql()

	}

	if utils.SearchString(args, "php7.4") == true {

		Installphp()
		Php74()

	}

	if utils.SearchString(args, "php7.3") == true {

		Installphp()
		Php73()

	}

	if utils.SearchString(args, "php7.2") == true {

		Installphp()
		Php72()

	}

}
