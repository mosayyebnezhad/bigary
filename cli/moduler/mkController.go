package moduler

import (
	Filer "clean/cmd/cli/file"
	"fmt"
)

//1- make controller in  {{moduleName}} / api / controller / {{ControllerName}}Controller.go

func Controller(moduleName string, ControllerName string) {
	templ := `package controller

type {{.ControllerName}}Controller struct {
	service I{{.ControllerName}}.{{.ControllerName}}Service
}

func New{{.ControllerName}}Controller(service I{{.ControllerName}}.{{.ControllerName}}Service) *{{.ControllerName}}Controller {
	return &{{.ControllerName}}Controller{service: service}
}`

	addr := fmt.Sprintf("module/%s/api/controller/%sController", moduleName, ControllerName)

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
