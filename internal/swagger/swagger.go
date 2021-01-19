package swagger

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-tire/tire/internal/engine"

	"github.com/golang-tire/pkg/log"
)

const template = `// Code generated by "tire embed swagger". DO NOT EDIT.
package {{.PkgName}}

import (
	"encoding/json"
)

const {{.Filename}}_paths = {{.Paths | quote }}
const {{.Filename}}_definitions = {{.Definitions | quote}}

func Register_{{.Filename}}_swagger(paths *map[string]interface{}, definitions *map[string]interface{}) error{
	err := json.Unmarshal([]byte({{.Filename}}_paths), paths)
	if err != nil{
		return err
	}
	err = json.Unmarshal([]byte({{.Filename}}_definitions), definitions)
	if err != nil{
		return err
	}
	return nil
}
`

type swaggerFile struct {
	Swagger     string                 `json:"swagger"`
	Schemes     []string               `json:"schemes"`
	Consumes    []string               `json:"consumes"`
	Produces    []string               `json:"produces"`
	Paths       map[string]interface{} `json:"paths"`
	Definitions map[string]interface{} `json:"definitions"`
}

type data struct {
	PkgName     string
	Filename    string
	Paths       string
	Definitions string
}

func EmbedSwagger(jsonFilename, output, packageName string) error {

	err := log.Init(context.Background(), true)
	if err != nil {
		return err
	}

	filename := strings.Split(filepath.Base(jsonFilename), ".")[0]
	fl, err := os.Open(jsonFilename)
	if err != nil {
		return err
	}

	var swag swaggerFile
	if err = json.NewDecoder(fl).Decode(&swag); err != nil {
		return err
	}

	path, err := json.Marshal(swag.Paths)
	if err != nil {
		return err
	}

	def, err := json.Marshal(swag.Definitions)
	if err != nil {
		return err
	}

	oData := data{
		PkgName:     packageName,
		Filename:    filename,
		Paths:       string(path),
		Definitions: string(def),
	}

	if output == "Stdout" {
		return engine.Render(template, oData, os.Stdout)
	}

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}

	defer outFile.Close()
	return engine.Render(template, oData, outFile)
}