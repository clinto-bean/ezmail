package main

import (
	"log"
	"os"
	"time"
)

func main() {
	a := API{
		port: "8080",
	}
	d := DB{
		DB_URL: os.Getenv("DB_URL"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}
	apiRetries, dbRetries := 0, 0

	apiErrChan := make(chan error)
	dbErrChan := make(chan error)
	go func() {
		for {
			err := a.Start()
			if apiRetries > 2 {
				log.Printf("fatal api error: %v", err)
				break
			}
			apiErrChan <- err
			apiRetries++
			time.Sleep(1 * time.Second) // prevent tight restart loop
		}
	}()
	go func() {
		for {
			err := d.Conn()
			if dbRetries > 2 {
				log.Printf("fatal DB error: %v", err)
				break
			}
			dbErrChan <- err
			
			dbRetries++
			time.Sleep(1 * time.Second) // prevent restart loop as with a.Start, db will be migrated out of application as todo
		}
	}()
	for {
		select {
		case err := <-apiErrChan:
			log.Printf("API crashed: %v. Restarting.", err)
		case err := <- dbErrChan:
			log.Printf("DB crashed: %v. Restarting.", err)
		}

	}


}