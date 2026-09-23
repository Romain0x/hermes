package identity

import (
	"net"
	"sync"
)

type User struct {
	name string
	conn net.Conn
	writeMutex sync.Mutex
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetConn() net.Conn {
	return u.conn
}

/* ------------------------------------ */

var userList = []*User{}

// Add a user to the list of users
func addUser(u *User) {
	userList = append( userList, u )
}

// Get a user based on the name provided as a parameter
func getUser(name string) *User {

	for u := range len(userList) {

		if name == userList[u].GetName() {
			return userList[u]
		}
	}

	return nil
}

// Returns true if a user has the name provided as a parameter.
func existsUserByName(name string) bool {

	for u := range len(userList) {

		if name == userList[u].GetName() {
			return true
		}
	}

	return false
}

// Returns all users as an array of User objects, except for the user provided as a parameter.
func GetAllUserExcept(u *User) []*User {

	myList := []*User{}

	for i := range len(userList) {

		if userList[i].GetName() != u.GetName() {
			myList = append(myList, userList[i])
		}
	}

	return myList
}

/* ------------------------------------ */

// SafeWrite writes b to the connection, preventing concurrency issues.
func (u *User) SafeWrite(b []byte) (int, error) {
	u.writeMutex.Lock()
	defer u.writeMutex.Unlock()
	return u.GetConn().Write(b)
}
