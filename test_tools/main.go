package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$")
	//path := r.URL.Path
	path := r.URL.Path[1:] // exclude the first characters frSom the url
	match := re.FindAllStringSubmatch(path, -1)
	if match != nil {
		fmt.Fprintf(w, "Hello, gopher %s\n", match[0][1])
		return
	}
	fmt.Fprintf(w, "Hello dear %s \n", path) //imprime en la web :8088

}

/**
Algunas herramientas de GO:
- go get
- go install

- go list
- go list -f '{{ .Imports }}' 					// imports de archivos
- go list -f '{{ .Name }}: {{ .Doc }}' 			// nombre de los paquetes y la documentación que se encuentra en los archivos
- go list -f '{{ join .Imports "\n" }}' fmt  	// lista los imports que usa el paquete fmt

- go doc
- go doc fmt
- go doc fmt Println

- go net/http

- go checherr
	Si algunas vez tienes un error y la consola no te dice nada puedes ejecutar "checherr" sobre el directorio del archivo

- go vet // differente compile
	this command will give you code that maybe would be wrong, could have possible any issue

- debugger go
	for the configuration of the json only change the direction and is all: "${workspaceRoot}/test_tools",

- Unit testing


*/
