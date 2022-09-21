package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")
	w.Write([]byte("This working world"))
}

func ussd_callback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")

	session_id := r.FormValue("sessionId")
	service_code := r.FormValue("serviceCode")
	phone_number := r.FormValue("phoneNumber")
	_ = fmt.Sprintf("%s, %s, %s", session_id, service_code, phone_number)
	fmt.Println(session_id, service_code, phone_number)
	text := r.FormValue("text")
	if len(text) == 0 {
		w.Write([]byte("Welcome to Ethio Ekub \n Choose your Actions \n 1. Login \n 2. signup \n 3. info	 "))
	} else {
		switch text {
		case "1":
			w.Write([]byte("You have selected Login"))
			return
		case "2":
			w.Write([]byte("You have selected Signup"))
			return
		case "3":
			w.Write([]byte("You have selected Info"))
			return
		default:
			w.Write([]byte("Invalid Input"))
			return
		}
	}
}

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8083"
	}

	fmt.Println("Ussd app using go lang")

	http.HandleFunc("/", ussd_callback)
	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
