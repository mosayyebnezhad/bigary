package moduler

import (
	Filer "clean/cmd/cli/file"
	"fmt"
)

//2- make interface in  {{moduleName}} / repository|service / {{controllerName}} / repository.go|service.go

func Interface_repository(moduleName string, InReName string) {
	templ := `package I{{.InReName}}

type {{.InReName}}Reposirory interface{
// write your functionalities here ...

}`

	addr := fmt.Sprintf("module/%s/interface/%s/repository", moduleName, InReName)

	structur := struct {
		//interface repository name
		InReName string
	}{
		InReName: InReName,
	}

	sErr := Filer.Make(addr, templ, structur)

	if sErr != nil {
		fmt.Println(sErr)

		return
	}

	fmt.Println("✅ on", addr)
}

func Interface_service(moduleName string, InReName string) {

	templ := `package I{{.InReName}}

type {{.InReName}}Service interface{
// write your functionalities here ...
}`

	addr := fmt.Sprintf("module/%s/interface/%s/service", moduleName, InReName)

	structur := struct {
		//interface service name
		InReName string
	}{
		InReName: InReName,
	}

	sErr := Filer.Make(addr, templ, structur)

	if sErr != nil {
		fmt.Println(sErr)

		return
	}

	fmt.Println("✅ on", addr)

}
