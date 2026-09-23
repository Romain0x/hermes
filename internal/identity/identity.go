package identity

import (
	"bufio"
	"net"
	"strings"
)


func CreatePerson(conn net.Conn) (*User, *bufio.Reader) {

	var name string

	reader := bufio.NewReader(conn)

	for {
		conn.Write([]byte("Entrer un nom : "))

		message, err := reader.ReadString('\n')
		if err != nil {
			return nil, nil
		}

		name = strings.TrimSpace(message)

		if existsUserByName(name) {
			conn.Write([]byte("Utilisateur existant\n"))
			continue
		}

		break
	}

	user := &User{
		name: name,
		conn: conn,
	}

	addUser(user)

	return user, reader
}
