package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Обработчик для статических файлов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Обработчик для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./index.html")
			return
		}
		http.NotFound(w, r)
	})

	// Обработчики для страниц ошибок
	http.HandleFunc("/401", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		http.ServeFile(w, r, "./401.html")
	})

	http.HandleFunc("/403", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		http.ServeFile(w, r, "./403.html")
	})

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "./500.html")
	})

	fmt.Println("Сервер запущен на http://localhost:80")
	http.ListenAndServe(":80", nil)
}
