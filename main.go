package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	qr "github.com/skip2/go-qrcode"
)

func main() {
	// add the content before the app is started.
	err := qr.WriteFile("http://184.72.118.42:80/htetmyat", qr.Medium, 256, "htetmyatthar1.png")
	if err != nil {
		log.Fatalln("qr code generation gone wrong.")
	}
	tmpl := template.Must(template.ParseFiles("index.html", "redirect.html"))

	http.HandleFunc("/htetmyat", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w, r, "/htetmyatthar/new", http.StatusFound)
		return
	})

	http.HandleFunc("/htetmyatthar/org", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	})

	http.HandleFunc("/htetmyatthar/new", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "redirect.html", nil)
		return
	})

	fmt.Println("server started on port 80.")
	log.Fatal(http.ListenAndServe(":80", nil))
}
