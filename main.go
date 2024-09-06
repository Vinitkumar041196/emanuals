package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func getFileList(dir string) ([]string, error) {
	dirFs, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	files, err := dirFs.ReadDir(0)
	if err != nil {
		return nil, err
	}

	outArr := []string{}
	for _, file := range files {
		outArr = append(outArr, strings.Split(file.Name(), ".")[0])
	}
	return outArr, nil
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Required resource not found: %s", err.Error())
		return
	}
	files, err := getFileList("static/manuals")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "No data to load: %s", err.Error())
		return
	}

	tmp.Execute(w, files)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))
	http.Handle("/manuals/", http.StripPrefix("/manuals/", http.FileServer(http.Dir("./static/manuals"))))

	http.HandleFunc("/", handleRoot)

	log.Println("listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
