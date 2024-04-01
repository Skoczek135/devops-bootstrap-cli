package create

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

// createFile creates a file with the given fileName
// if file exists then it does nothing and return false
// if file does not exist then it creates the file and returns true
func createFile(fileName string) (*os.File, error) {
	f, err := os.Stat(fileName)
	if f != nil {
		return nil, errors.New(fmt.Sprintf("File %s already exists", fileName))
	}
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, errors.New(fmt.Sprintf("Something goes wrong while creating file", fileName))
}

// fillFileWithTemplate fills the file with the given template string
func fillFileWithTemplate(file *os.File, tmpl string) error {
	t := template.Must(template.New(tmpl).Parse(tmpl))
	err := t.Execute(file, nil)
	if err != nil {
		return err
	}
	return nil
}
