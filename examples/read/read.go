package main

import (
	"github.com/davecheney/serial"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %v DEVICE", os.Args[0])
	}
	s, err := serial.Open(os.Args[1], 115200)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var b [1]byte
		_, err := s.Read(b[:])
		if err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(2 * time.Second)
	log.Println("timeout expired")
	s.Close()
	log.Println("port closed")
	wg.Wait()
}
