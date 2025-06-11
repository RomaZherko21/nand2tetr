package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	for {
		message := "Привет, TCP сервер!"
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			return
		}

		fmt.Printf("Отправлено сообщение: %s\n", message)
		time.Sleep(1 * time.Second)

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при получении ответа:", err)
			return
		}

		fmt.Printf("Получен ответ от сервера: %s", response)
	}
}
