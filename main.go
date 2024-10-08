package main

import "log"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	service := NewChuckNorrisJokeService("https://api.chucknorris.io/jokes/random")
	service = NewLoggingService(service)

	apiServer := NewApiServer(service)
	log.Fatal(apiServer.Start(":4200"))
}
