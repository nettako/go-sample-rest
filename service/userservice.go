package service

import (
	"encoding/json"
	"log"
	"net/http"

	"../config"
	"../model"
)

func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users model.Users
	var arrUser []model.Users
	var response model.Response

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("Select userid,firstname,lastname from core_user_list")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
