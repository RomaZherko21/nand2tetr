package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании адреса:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Ошибка при создании сервера:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP сервер запущен на :8080")

	for {
		buffer := make([]byte, 1024) // Увеличим буфер для больших сообщений
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Получено %d байт от %v: %s\n", n, remoteAddr, message)

		// Отправляем эхо-ответ обратно клиенту
		_, err = conn.WriteToUDP([]byte("Hello, client!"), remoteAddr)
		if err != nil {
			fmt.Println("Ошибка при отправке ответа:", err)
			continue
		}
	}
}
