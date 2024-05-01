package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Тест"))
}

func register(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Тест"))
}

func login(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Тест"))
}