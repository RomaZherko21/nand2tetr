package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Ошибка при разрешении адреса:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Ошибка при создании соединения:", err)
		return
	}
	defer conn.Close()

	for {
		message := "Привет, UDP сервер!"
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			return
		}

		fmt.Printf("Отправлено сообщение: %s\n", message)
		time.Sleep(1 * time.Second)
	}
}
