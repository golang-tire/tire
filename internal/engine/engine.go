package engine

import (
	"io"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func Render(tpl string, data interface{}, wr io.Writer) error {
	tmpl, err := template.New("render").Funcs(sprig.TxtFuncMap()).Parse(tpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(wr, data)
}
