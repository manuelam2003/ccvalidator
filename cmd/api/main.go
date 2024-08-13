package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"unicode"

	"github.com/manuelam2003/ccvalidator/internal/luhn"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("POST /check", checkPost)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func checkPost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	cardNumber := r.FormValue("cardNumber")
	if isValidInput(cardNumber) {
		if luhn.CheckLuhn(cardNumber) {
			fmt.Fprintf(w, `<div class="alert alert-success">The credit card %s is valid.</div>`, cardNumber)
		} else {
			fmt.Fprintf(w, `<div class="alert alert-danger">The credit card %s is not valid.</div>`, cardNumber)
		}
	} else {
		fmt.Fprintf(w, `<div class="alert alert-danger">It must be a number.</div>`)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Validate the input to ensure it contains only digits
func isValidInput(number string) bool {
	for _, r := range number {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
