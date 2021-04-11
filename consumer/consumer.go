package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const FILEPATH = "/app/messages.txt"

func read_from_file(offset int64) int64 {
	f, e := os.Open(FILEPATH)
	if e != nil {
		return 0
	}

	defer f.Close()

	_, err := f.Seek(offset, 0)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		fmt.Printf("%s", line)

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		panic(err)
	}
	new_pos, err := f.Seek(0, 2)
	if err != nil {
		panic(err)
	}
	return new_pos
}

func main() {
	c := time.Tick(3 * time.Second)
	pos := int64(0)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for _ = range c {
		pos = read_from_file(pos)
	}
}
