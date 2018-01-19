package main

import "github.com/12GardenSeed/go_agenda/service/service"

func main() {
    port := ":8080"
    server :=service.NewServer()
    server.Run(port)
}
