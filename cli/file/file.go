package Filer

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func Make(addr string, templ string, strc any) error {

	f, fErr := fileCreator(addr)
	if fErr != nil {
		return fErr
	}
	defer f.Close()
	tmp, tErr := templateCreator(templ)
	if tErr != nil {
		return tErr
	}
	err := templateExecuter(tmp, f, strc)

	if err != nil {
		return err
	}

	return nil
}

func fileCreator(address string) (*os.File, error) {
	// استخراج مسیر دایرکتوری از آدرس فایل
	dir := filepath.Dir(address)

	// ساختن دایرکتوری در صورت نبود
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("❌ Error creating directory: %w", err)
	}

	// حالا ساختن فایل
	f, err := os.Create(address + ".go")
	if err != nil {
		return nil, fmt.Errorf("❌ Error creating file: %w", err)
	}
	return f, nil
}
func templateCreator(templ string) (*template.Template, error) {
	tmpl, err := template.New("controller").Parse(templ)
	if err != nil {
		return nil, errors.New(fmt.Sprint("❌ Template parse error:", err))
	}
	return tmpl, nil
}
func templateExecuter(temple *template.Template, file *os.File, data any) error {

	err := temple.Execute(file, data)
	if err != nil {
		return errors.New(fmt.Sprint("❌ Template execution error:", err))
	}
	return nil
}
