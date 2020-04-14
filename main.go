package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/beka/project/models"
)

func main() {
	models.InitDB("root@tcp(localhost:3306)/Salon")
	templates := template.Must(template.ParseFiles("templates/main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "main.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/catalog", CatalogIndex)
	http.HandleFunc("/master", MasterIndex)
	http.ListenAndServe(":8080", nil)
}
func CatalogIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	sum, err := models.AllCatalog(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, su := range sum {
		fmt.Fprintln(w, su.Name, su.Surname, su.Specialty, su.Style, su.Price, su.WorkTime)

	}
}
func MasterIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	ms, err := models.Masters(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, m := range ms {
		fmt.Fprintln(w, m.MasterID, m.Name, m.Surname, m.Phone, m.Email)
	}
}

func EditIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	mas, err := models.EditPage(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, mn := range mas {
		fmt.Fprintln(w, mn.MasterID, mn.Name, mn.Surname, mn.Phone, mn.Email)
	}
}
