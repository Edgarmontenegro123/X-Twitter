package db

import (
	"github.com/Edgarmontenegro123/X-Twitter/internal/tweet"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
)

type Database struct {
	Users  map[int]user.User
	Tweets map[int][]tweet.Tweet
}

func NewDatabase() *Database {
	return &Database{
		Users:  make(map[int]user.User),
		Tweets: make(map[int][]tweet.Tweet),
	}
}

func (db *Database) SaveUser(user user.User) {
	db.Users[user.ID] = user
}

func (db *Database) GetUserByID(userID int) user.User {
	return db.Users[userID]
}

func (db *Database) GetUsers() map[int]user.User {
	return db.Users
}

func (db *Database) SaveTweet(tweet tweet.Tweet) {
	db.Tweets[tweet.UserId] = append(db.Tweets[tweet.UserId], tweet)
}

func (db *Database) GetTweetsByUserID(userID int) []tweet.Tweet {
	return db.Tweets[userID]
}
