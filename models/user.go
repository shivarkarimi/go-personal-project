package models

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID int32 = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = int(nextID)
	nextID++
	users = append(users, &u)
	return u, nil
}

func DeleteUserById(id int) error {

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v  not found", id)
}

func UpdateUser(user User) (User, error) {

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v  not found", user.ID)
}
