package tweet

import (
	"errors"
	"fmt"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
)

type TweetDB interface {
	SaveTweet(tweet Tweet)
	GetTweetsByUserID(userID int) []Tweet
	GetUserByID(userID int) user.User
	GetFollowers(userID int) ([]user.User, error)
}

type TweetService struct {
	db TweetDB
}

func NewTweetService(db TweetDB) *TweetService {
	return &TweetService{
		db: db,
	}
}

func (ts *TweetService) PublishTweet(userID int, content string) error {
	if userID == 0 {
		return errors.New("el usuario no existe")
	}

	// Crear un nuevo tweet
	newTweet := Tweet{
		ID:      len(ts.db.GetTweetsByUserID(userID)) + 1,
		Content: content,
		UserId:  userID,
	}

	// Guardar el usuario actualizado en la DB
	ts.db.SaveTweet(newTweet)

	return nil
}

func (ts *TweetService) GetTweets(userID int) ([]TimeLineTweet, error) {
	// Obtener tweets desde la DB
	userTweets := ts.db.GetTweetsByUserID(userID)

	// Obtener la lista de usuarios a los que sigue
	following, err := ts.db.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	// Obtener tweets de usuarios a los que sigue
	var timelineTweets []TimeLineTweet
	for _, followingUser := range following {
		followingUserTweets := ts.db.GetTweetsByUserID(followingUser.ID)
		for _, tweet := range followingUserTweets {
			timelineTweets = append(timelineTweets, TimeLineTweet{Tweet: tweet, User: followingUser})
		}
	}

	// Agregar tweets del usuario actual
	for _, tweet := range userTweets {
		timelineTweets = append(timelineTweets, TimeLineTweet{Tweet: tweet, User: ts.db.GetUserByID(userID)})
	}

	return timelineTweets, nil
}

func (ts *TweetService) GetTweetsTimeline(userID int) ([]TimeLineTweet, error) {
	// Obtener la lista de usuarios a los que sigue
	following, err := ts.db.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	// Obtener tweets de usuarios a los que sigue
	var timelineTweets []TimeLineTweet
	for _, follower := range following {
		followerTweets := ts.db.GetTweetsByUserID(follower.ID)
		for _, tweet := range followerTweets {
			timelineTweets = append(timelineTweets, TimeLineTweet{Tweet: tweet, User: ts.db.GetUserByID(follower.ID)})
		}
	}

	// agregar tweets de usuario actual al timeline
	userTweets := ts.db.GetTweetsByUserID(userID)

	for _, tweet := range userTweets {
		timelineTweets = append(timelineTweets, TimeLineTweet{Tweet: tweet, User: ts.db.GetUserByID(userID)})
	}

	fmt.Printf("Total timelineTweets: %d\n", len(timelineTweets))
	return timelineTweets, nil
}

func (ts *TweetService) PrintUserTimeline(userID int) {
	timelineTweets, err := ts.GetTweetsTimeline(userID)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Timeline de usuario%d: %+v\n", userID, timelineTweets)
}
