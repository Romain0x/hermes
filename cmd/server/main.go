package main

import (
	"hermes/internal/design"
	"hermes/internal/handlers"
	"hermes/internal/identity"

	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening : ", err)
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection : ", err)
			continue
		}

		conn.Write([]byte(design.Banner))

		user, reader := identity.CreatePerson(conn)
		if user == nil {
			log.Fatal("Erreur de création de personne")
		}

		go handlers.HandleConnection(conn, user, reader)
	}
}
