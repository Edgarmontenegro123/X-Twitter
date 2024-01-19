package db

import "github.com/Edgarmontenegro123/X-Twitter/internal/user"

type Database struct {
	users map[int]user.User
}

func NewDatabase() *Database {
	return &Database{
		users: make(map[int]user.User),
	}
}

func (db *Database) SaveUser(user user.User) {
	db.users[user.ID] = user
}

func (db *Database) GetUserByID(userID int) user.User {
	return db.users[userID]
}
