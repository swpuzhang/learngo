package usetemplate

import (
	"fmt"
	"html/template"
	"os"
)

// UseTemplate 使用template
func UseTemplate() {
	useTem()
}

func useTem() {
	templ := template.Must(template.New("tem").Parse(`<p>A:{{.A}}</p><p>B:{{.B}}</p>`))
	var data struct {
		A string
		B template.HTML
	}
	data.A = `<b>Hello</b>
	`
	data.B = `<b>Hello</b>
	`
	if err := templ.Execute(os.Stdout, data); err != nil {
		fmt.Println(err)
		return
	}

}
