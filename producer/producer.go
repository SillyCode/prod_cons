package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FILEPATH = "/app/messages.txt"

func write_to_file(message string) {
	file_handler, err := os.OpenFile(FILEPATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file_handler.Close()

	writer := bufio.NewWriter(file_handler)
	_, err = writer.WriteString(fmt.Sprintf("%s", message))
	if err != nil {
		panic(err)
	}
	writer.Flush()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write any message. To exit type CTRL+C")
	for {
		body, _ := reader.ReadString('\n')
		write_to_file(body)
		log.Printf(" [x] Wrote message: %s", body)
	}
}
