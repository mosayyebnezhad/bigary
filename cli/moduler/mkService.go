package moduler

import (
	Filer "clean/cmd/cli/file"
	"fmt"
)

func Service(moduleName, ControllerName string) {
	templ := `package service


type {{.ControllerName}}Service struct {
	repo I{{.ControllerName}}.{{.ControllerName}}Reposirory
}

func New{{.ControllerName}}Service(repo I{{.ControllerName}}.{{.ControllerName}}Reposirory) I{{.ControllerName}}.{{.ControllerName}}Service {
	return &{{.ControllerName}}Service{repo: repo}
}

//func (l *{{.ControllerName}}Service) Example(userID uint) (string, error) {

`

	addr := fmt.Sprintf("module/%s/service/%sService", moduleName, ControllerName)

	structur := struct {
		ControllerName string
	}{
		ControllerName: ControllerName,
	}

	sErr := Filer.Make(addr, templ, structur)

	if sErr != nil {
		fmt.Println(sErr)

		return
	}

	fmt.Println("✅ on", addr)
}
