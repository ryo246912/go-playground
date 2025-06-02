package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.New("index").ParseGlob("*.tmpl"))
	if err := t.ExecuteTemplate(os.Stdout, "index", "test"); err != nil {
		log.Fatalln(err)
	}
}
