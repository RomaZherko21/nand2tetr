package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer ln.Close()

	fmt.Println("TCP сервер запущен и ожидает подключения...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			return
		}

		fmt.Printf("Получено сообщение: %s", message)

		upperMessage := strings.ToUpper(message)
		_, err = conn.Write([]byte(upperMessage))
		if err != nil {
			fmt.Println("Ошибка при отправке ответа:", err)
			return
		}
	}
}
