package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserDetail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UserDetails []UserDetail

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/get_account", getUserDetails)
	myRouter.HandleFunc("/create_account", createUserDetails).Methods("POST")
	myRouter.HandleFunc("/delete_account", deleteUserAccount).Methods("DELETE")
	myRouter.HandleFunc("/update_account", updateUserDetails).Methods("UPDATE")
	log.Fatal(http.ListenAndServe(":100", myRouter))
}

func createUserDetails(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var detail UserDetail
	json.Unmarshal(reqBody, &detail)

	if contains(UserDetails, detail) {

		fmt.Fprint(w, "User already has an account created")

	} else {
		UserDetails = append(UserDetails, detail)
		fmt.Fprint(w, "User account created Successfully")
	}
	fmt.Println(UserDetails)
}
func getUserDetails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(UserDetails)

}
func deleteUserAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var detail UserDetail
	json.Unmarshal(reqBody, &detail)
	for index, val := range UserDetails {
		if detail.Email == val.Email {
			UserDetails = append(UserDetails[:index], UserDetails[index+1:]...)
			fmt.Fprint(w, "User account Deleted Successfully")
		}
	}

}
func updateUserDetails(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var detail UserDetail
	json.Unmarshal(reqBody, &detail)
	for index, val := range UserDetails {
		if detail.Email == val.Email {
			UserDetails[index].Password = detail.Password
			fmt.Fprintf(w, "User password Changed Successfully")
		}
	}
}
func contains(s []UserDetail, val UserDetail) bool {
	for _, v := range s {
		if v.Email == val.Email {
			return true
		}
	}
	return false
}
func main() {
	UserDetails = []UserDetail{}
	handleRequest()

}
