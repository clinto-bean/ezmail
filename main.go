package main

import (
	"log"
	"time"
)

func main() {
	a := API{
		port: "8080",
	}

	apiErrChan := make(chan error)

	go func() {
		for {
			err := a.Start()
			if err != nil {
				return
			}
			apiErrChan <- err
			time.Sleep(1 * time.Second) // prevent tight restart loop
		}
	}()

	for {
		err := <-apiErrChan
		log.Printf("API crashed: %v. Restarting.", err)
	}
}