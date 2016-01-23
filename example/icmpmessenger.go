package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Maksadbek/tcpovericmp/icmp"
)

func main() {
	i, err := icmp.New("127.0.0.1", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			msg, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}
			err = i.Write(msg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	go func() {
		for {
			msg, err := i.Read()
			if err != nil {
				log.Fatal(err)
			}
			if msg == nil {
				continue
			}
			fmt.Println(string(msg))
		}
	}()
	<-make(chan struct{})
}
