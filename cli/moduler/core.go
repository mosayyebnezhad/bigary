package moduler

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//2- make interface in  {{moduleName}} / interfaca / {{controllerName}} / ripository.go & service.go

//3- make repository in {{moduleName}} / api / repository / {{controllerName}}repository.go

//3- make service in {{moduleName}} / api / service / {{controllerName}}Service.go

// func Engine() {
// 	// Controller("auth", "hi")
// 	// Interface_service("auth", "alnaghi")
// 	// Interface_repository("auth", "alnaghi")
// 	// Service("auth", "chang")
// 	Repository("auth", "pesh")

// }

var Moduler = &cobra.Command{
	Use:   "moduler [type] [name] [...options]",
	Short: "generate module component",
	Args:  cobra.RangeArgs(1, 4),
	Run: func(cmd *cobra.Command, args []string) {

		lenth := len(args)
		// 0: moduler
		//starting a wizard to fillign out all of inputs

		// 1: moduler users
		// 1: moduler users ...

		// 2: moduler auth users
		// 2: moduler auth users ...

		// 3: moduler interface auth users
		// 3: moduler controller auth users ...

		if lenth == 2 {
			moduleWith2Inp(cmd, args)
		}
		if lenth == 3 {
			moduleWith4Inp(cmd, args)
		}

	},
}

func moduleWith2Inp(cmd *cobra.Command, args []string) {

	b := Sure(args)
	if !b {
		cmd.PrintErrln("❌ process stopped !")
		os.Exit(1)
	}

	f, s := args[0], args[1]
	Interface_repository(f, s)
	Interface_service(f, s)
	Service(f, s)
	Repository(f, s)
	Controller(f, s)

}
func moduleWith4Inp(cmd *cobra.Command, args []string) {

	allowedTypes := map[string]bool{
		"controller": true,
		"service":    true,
		"repository": true,
		"interface":  true,
	}

	kind := args[0]
	if !allowedTypes[kind] {
		cmd.PrintErrln("❌ invalid type. must be one of: controller, service, repository, interface")
		os.Exit(1)
	}

	b := Sure(args)
	if !b {
		cmd.PrintErrln("❌ process stopped !")
		os.Exit(1)
	}

	switch kind {
	case "controller":
	case "service":
	case "repository":
	case "interface":
	}

}

func Sure(addr []string) bool {
	t := len(addr)
	scan := bufio.NewScanner(os.Stdin)

	fmt.Printf("‼️ are you sure to create full {  %s  } in {  %s  } module ? \n", addr[t-1], addr[t-2])
	scan.Scan()

	answer := scan.Text()

	if answer == "yes" || answer == "yeah" || answer == "y" || answer == "بله" || answer == "آره" {
		return true
	}
	return false

}
