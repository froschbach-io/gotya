package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Values map[string]interface{}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: " + os.Args[0] + " [values file (yaml)] [template file]")
		os.Exit(1)
	}
	valuesFile := os.Args[1]
	templateFile := os.Args[2]

	valuesData, err := ioutil.ReadFile(valuesFile)
	if err != nil {
		fmt.Printf("Failed to read %v: %v\n", valuesFile, err)
		os.Exit(1)
	}

	values := Values{}
	err = yaml.Unmarshal([]byte(valuesData), &values)
	if err != nil {
		fmt.Printf("Failed to parse %v: %v\n", valuesFile, err)
		os.Exit(1)
	}

	tmplData, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Printf("Failed to read %v: %v\n", templateFile, err)
		os.Exit(1)
	}

	// custom function that can be called from within the template
	funcMap := template.FuncMap{
		"add": func(a int, b int) int {
			return a + b
		},
	}

	tmpl := template.New(templateFile).Funcs(funcMap).Delims("[[", "]]")
	tmpl, err = tmpl.Parse(string(tmplData))
	if err != nil {
		fmt.Printf("Failed to parse template file: %v\n", err)
		os.Exit(1)
	}

	out := new(bytes.Buffer)
	err = tmpl.Execute(out, values)
	if err != nil {
		fmt.Printf("Failed to apply values to template: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(out)
}
