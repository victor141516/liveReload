package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
)

var dirUpdated = 0

func BuildWatcher(watchDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == fsnotify.Write {
					dirUpdated = 1
					log.Println("Refreshing...", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()
	err = watcher.Add(watchDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func checkIfUpdated(w http.ResponseWriter, r *http.Request) {
	message := strconv.Itoa(dirUpdated)
	if dirUpdated == 1 {
		dirUpdated = 0
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(message))
}

func main() {
	watchDir := "."
	port := "5555"
	if len(os.Args) > 1 {
		watchDir = os.Args[1]
	}
	if len(os.Args) > 2 {
		port = os.Args[2]
	}
	go BuildWatcher(watchDir)
	http.HandleFunc("/", checkIfUpdated)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
