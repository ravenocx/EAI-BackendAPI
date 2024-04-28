package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/ravenocx/EAI-BackendAPI/internal/app"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer func(){
			if err := recover();err != nil{
				log.Printf("Recover from panic in StartApplication : %+v", err)
			}
		}()
		app.StartApplication()
	}()

	// keep ping to the /ping
	go func() {
		defer wg.Done()
		defer func(){
			if err := recover();err != nil{
				log.Printf("Recover from panic in Ping to /ping : %+v", err)
			}
		}()
		for {
			resp, err := http.Get("https://be-esd-website.onrender.com/ping")
			if err != nil {
				log.Printf("Error test ping: %s\n", err)
				continue
			}
			log.Printf("Ping!!!")
			// Close the response body
			resp.Body.Close()
			time.Sleep(1 * time.Minute)
		}
	}()

	wg.Wait() //keep the goroutines alive
}
