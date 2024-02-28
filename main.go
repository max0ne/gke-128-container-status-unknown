package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// self destruct in few seconds
	go func() {
		time.Sleep(10 * time.Second)
		_, err := allocateEphemeralStorage(1024 * 1024 * 1024)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Create a channel to receive the SIGTERM signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Wait for the signal
	sig := <-sigChan
	log.Println("Received signal", sig)
	for {
		log.Println("sleeping...")
		time.Sleep(time.Second)
	}
}

func allocateEphemeralStorage(size int64) (int64, error) {
	log.Printf("Generating %d bytes Ephemearl storage usage", size)

	file := "big-file"
	block := int64(4096 * 1024)

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return 0, fmt.Errorf("unable to open file: %s", err)
	}

	for i := int64(0); i < size; i += block {
		_, err := f.Write(make([]byte, block))
		if err != nil {
			return 0, fmt.Errorf("unable to write to file: %s", err)
		}
	}

	info, err := os.Stat(file)
	if err != nil {
		return 0, fmt.Errorf("unable to get file info: %s", err)
	}

	return info.Size(), nil
}
