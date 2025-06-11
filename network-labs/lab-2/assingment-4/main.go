package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// Создаем обработчик для всех путей
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Создаем URL для проксируемого сервера
		targetURL := &url.URL{
			Scheme: "http",
			Host:   "localhost:80",
			Path:   r.URL.Path,
		}

		// Копируем оригинальный запрос
		proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка создания запроса: %v", err), http.StatusInternalServerError)
			return
		}

		// Копируем заголовки из оригинального запроса
		for header, values := range r.Header {
			for _, value := range values {
				proxyReq.Header.Add(header, value)
			}
		}

		// Добавляем X-Forwarded заголовки
		proxyReq.Header.Set("X-Forwarded-For", r.RemoteAddr)
		proxyReq.Header.Set("X-Forwarded-Host", r.Host)
		proxyReq.Header.Set("X-Forwarded-Proto", r.URL.Scheme)

		// Выполняем запрос к целевому серверу
		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка проксирования: %v", err), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Копируем заголовки ответа
		for header, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(header, value)
			}
		}

		// Устанавливаем статус код
		w.WriteHeader(resp.StatusCode)

		// Копируем тело ответа
		if _, err := io.Copy(w, resp.Body); err != nil {
			fmt.Printf("Ошибка копирования ответа: %v\n", err)
		}
	})

	fmt.Println("Прокси-сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
