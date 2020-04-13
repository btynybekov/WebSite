package models

import (
	"log"
	"net/http"
)

func VacationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
 
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
		}
		Name := r.FormValue("Name")
		Surname := r.FormValue("Surname")
		Specialty := r.FormValue("Specialty")
		Style := r.FormValue("Style")
		Worktime := r.FormValue("Worktime")
		Price := r.FormValue("Price")
		_, err = db.Exec("Insert into Salon.Master (Name, Surname) values(?, ?) Salon.Service (Specialty, Style, Price, Worktime) values (?, ?, ?, ?)", Name, Surname, Specialty, Style, Worktime, Price)
       
        if err != nil {
            log.Println(err)
          }
        http.Redirect(w, r, "/", 301)
    }else{
        http.ServeFile(w,r, "templates/create.html")
	}
}
