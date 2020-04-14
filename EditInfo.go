package models

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func EditPage(w http.ResponseWriter, r *http.Request) ([]*master, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	rows := db.QueryRow("SELECT * FROM Master where id = ?", id)
	mas := make([]*master, 0)
	for rows.Next() {
		mn := new(master)
		err := rows.Scan(&mn.MasterID, &mn.Name, &mn.Surname, &mn.Phone, &mn.Email)
		if err != nil {
			return nil, err
		}
		mas = append(mas, mn)
	}
	tmpl, _ := template.ParseFiles("templates/master.html")
	tmpl.Execute(w, mas)
	return mas, nil
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	Name := r.FormValue("Name")
	Surname := r.FormValue("Surname")
	Phone := r.FormValue("Phone")
	Email := r.FormValue("Email")
	_, err = db.Exec("Update Salon.Master set Name = ?, Surname = ?, Phone = ?, Email = ?", Name, Surname, Phone, Email)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}
