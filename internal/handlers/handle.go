package handlers

import (
	"bufio"
	"fmt"
	"hermes/internal/design"
	"hermes/internal/identity"
	"log"
	"net"
	"strings"
)

const maxNameLen = 10

// Manages the connection, input, and output.
func HandleConnection(conn net.Conn, user *identity.User, reader *bufio.Reader) {

	defer conn.Close()

	sendWelcome(user)

	for {

		msg, err := readMessage(reader)
		if err != nil {
			log.Fatal("Erreur sur la lecture")
		}

		if "/leave" == msg {
			return
		}

		user.SafeWrite(fmt.Appendf(nil, "%-*s : ", maxNameLen, user.GetName()))

		broadcastMessage(identity.GetAllUserExcept(user), user, msg)
	}
}


func sendWelcome(user *identity.User) {

	user.SafeWrite([]byte("\033[H\033[2J")) // Clear all
	user.SafeWrite([]byte(design.Banner))
	user.SafeWrite(fmt.Appendf(nil, "\t Log in as %s.\n\n%-*s : ", user.GetName(), maxNameLen, user.GetName()))
}


func readMessage(reader *bufio.Reader) (string, error) {

	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error : %v", err)
		return "", err
	}

	return strings.TrimSpace(message), nil
}


func broadcastMessage(listUser []*identity.User, user *identity.User, message string) {

	clearLine := "\r\033[2K" // ASCII character to clear a line

	for i := range len(listUser) {
		response := fmt.Sprintf("%s%-*s : %s\n%-*s : ", clearLine, maxNameLen, user.GetName(), message, maxNameLen, listUser[i].GetName())

		_, err := listUser[i].SafeWrite([]byte(response))
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}

			log.Printf("Server write error : %v", err)
		}

	}
}
