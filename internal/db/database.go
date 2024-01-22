package db

import (
	"errors"
	"github.com/Edgarmontenegro123/X-Twitter/internal/tweet"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
)

type Database struct {
	Users     map[int]user.User
	Tweets    map[int][]tweet.Tweet
	followers map[int][]int
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
	client, ok := db.Users[userID]
	if !ok {
		return nil
	}

	userTweets := db.Tweets[client.ID]
	return userTweets
}

func (db *Database) GetFollowers(userID int) ([]user.User, error) {
	// Cambio user por client para evitar conflictos con la importación del package
	client := db.GetUserByID(userID)
	if client.ID == 0 {
		return nil, errors.New("el usuario no existe")
	}

	followers := []user.User{}
	for _, followerID := range client.Followers {
		follower := db.GetUserByID(followerID)
		if follower.ID != 0 {
			followers = append(followers, follower)
		}
	}
	return followers, nil
}
