package models

import (
	"html/template"
	"log"
	"net/http"
)

type Master struct {
	Name    string
	Surname string
}

type Service struct {
	Specialty string
	Style     string
	WorkTime  string
	Price     string
}
type Catalog struct {
	MasterID  int
	ServiceID int
}

type Summary struct {
	Master
	Service
	Catalog
}

func AllCatalog(w http.ResponseWriter, r *http.Request) ([]*Summary, error) {
	rows, err := db.Query("SELECT m.name, m.Surname, s.Specialty, s.Style, s.Worktime, s.Price FROM Catalog c join service s on c.ServiceID = s.ServiceID join master m on c.MasterID = m.MasterID")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	sum := make([]*Summary, 0)
	for rows.Next() {
		su := new(Summary)
		err := rows.Scan(&su.Name, &su.Surname, &su.Specialty, &su.Style, &su.Price, &su.WorkTime)
		if err != nil {
			return nil, err
		}
		sum = append(sum, su)
	}
	tmpl, _ := template.ParseFiles("templates/catalog.html")
	tmpl.Execute(w, sum)
	return sum, nil
}
