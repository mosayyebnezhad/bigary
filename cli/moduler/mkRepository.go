package moduler

import (
	Filer "clean/cmd/cli/file"
	"fmt"
)

func Repository(moduleName, ControllerName string) {
	templ := `package repository

import "gorm.io/gorm"

type {{.ControllerName}}Repository struct {
	db *gorm.DB
}

func New{{.ControllerName}}Repository(db *gorm.DB) I{{.ControllerName}}.Repository {
	return &{{.ControllerName}}Repository{db: db}
}

// func (o *{{.ControllerName}}Repository) TestExample(e entity.Otp) error {panic("")}`

	addr := fmt.Sprintf("module/%s/repository/%sRepository", moduleName, ControllerName)

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
