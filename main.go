package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)

	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charst=UTF8)
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

	//fmt.Fprintf(w, `welcome`)
}

func features(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=UF")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error)
	}
	ptmp.Execute(w, nil)

	//fmt.Fprintf(w, `welcome`)
}

func docs(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charst=UTF")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}
