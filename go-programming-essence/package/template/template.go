package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func main() {
	tmpl := `
Template 1
{{.}}`
	t := template.Must(template.New("example").Parse(tmpl))
	if err := t.Execute(os.Stdout, "Hello, World!"); err != nil {
		log.Fatal(err)
	}

	t2, _ := template.New("example2").Parse(`
Template 2
{{.Name}}`)
	if err := t2.Execute(os.Stdout, struct {
		Name string
	}{
		Name: "Bob",
	}); err != nil {
		log.Fatal(err)
	}

	t3, _ := template.New("example3").Parse(`
Template 3
{{range .}}
<p>{{.}}</p>{{end}}
index:{{index . 1}}
`)
	if err := t3.Execute(os.Stdout, []string{
		"Hello", "World", "This is a template example.",
	}); err != nil {
		log.Fatal(err)
	}

	t4, _ := template.New("example3").Parse(`
Template 4
{{if gt .Age 20}}
{{.Name}} is older than 20
{{else}}
{{.Name}} is not older than 20
{{end}}

{{if or (eq .Name "Alice") (eq .Name "Bob") }}
 {{.Name}} is Alice or Bob
{{else}}
 {{.Name}} is neither Alice nor Bob
{{end}}
`)
	if err := t4.Execute(os.Stdout, struct {
		Age  int
		Name string
	}{
		Age:  25,
		Name: "Alice",
	}); err != nil {
		log.Fatal(err)
	}

	t5, _ := template.New("example5").Parse(`
Template 5
{{with index .Employees 0}}
{{.Name}}
{{end}}
--------------
{{with $v := index .Employees 0}}
{{$v.Name}}
{{end}}
`)
	if err := t5.Execute(os.Stdout, struct {
		Employees []struct {
			Name string
		}
	}{
		[]struct {
			Name string
		}{
			{Name: "Alice"},
			{Name: "Bob"},
			{Name: "Charlie"},
		},
	}); err != nil {
		log.Fatal(err)
	}

	t6, _ := template.New("example6").Funcs(
		template.FuncMap{
			"FormatDateTime": func(format string, d time.Time) string {
				if d.IsZero() {
					return "xxx"
				}
				return d.Format(format)
			},
		},
	).Parse(`
Template 6
{{FormatDateTime "2006-01-02 15:04:05" .}}`)
	if err := t6.Execute(os.Stdout, time.Now()); err != nil {
		log.Fatal(err)
	}

	tX, _ := template.New("example2").Parse(`
Template X
{{define "T"}}Hello, {{.}}!{{end}}
`)
	if err := tX.ExecuteTemplate(os.Stdout, "T", "Gopher"); err != nil {
		log.Fatal(err)
	}
}
