package tweet

import (
	"errors"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
)

type TweetDB interface {
	/*SaveUser(user user.User)
	GetUserByID(userID int) user.User*/
	SaveTweet(tweet Tweet)
	GetTweetsByUserID(userID int) []Tweet
	GetUserByID(userID int) user.User
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
	// Obtenemos el usuario desde la DB
	// user := ts.db.GetUserByID(userID)
	if userID == 0 {
		return errors.New("el usuario no existe") // Manejar el caso usuario no existe
	}

	// Crear un nuevo tweet
	newTweet := Tweet{
		ID:      len(ts.db.GetTweetsByUserID(userID)) + 1,
		Content: content,
		UserId:  userID,
	}

	/*// Agregar el tweet al usuario
	user.Tweets = append(user.Tweets, newTweet)*/

	// Guardar el usuario actualizado en la DB
	ts.db.SaveTweet(newTweet)

	return nil
}

func (ts *TweetService) GetTweets(userID int) ([]Tweet, error) {
	// Obtener tweets desde la DB
	return ts.db.GetTweetsByUserID(userID), nil
}
