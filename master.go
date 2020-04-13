package models

import (
	"html/template"
	"log"
	"net/http"
)

type master struct {
	MasterID int
	Name     string
	Surname  string
	Phone    int
	Email    string
}

func Masters(w http.ResponseWriter, r *http.Request) ([]*master, error) {
	rows, err := db.Query("SELECT * FROM Master")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	ms := make([]*master, 0)
	for rows.Next() {
		m := new(master)
		err := rows.Scan(&m.MasterID, &m.Name, &m.Surname, &m.Phone, &m.Email)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	tmpl, _ := template.ParseFiles("templates/master.html")
	tmpl.Execute(w, ms)
	return ms, nil
}
