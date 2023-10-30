package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	port := ":8080" // Порт, на котором будет работать сервер

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Читаем содержимое файла "index.html"
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Не удается прочитать HTML файл", http.StatusInternalServerError)
			return
		}

		// Устанавливаем правильную кодировку для HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(content)
	})

	// Запуск сервера
	fmt.Printf("Сервер запущен на порту %s\n", port)
	http.ListenAndServe(port, nil)
}
