package usecase

import (
	"github.com/Edgarmontenegro123/X-Twitter/internal/entity"
)

type TweetUseCase struct {
}

func (uc *TweetUseCase) PublishTweet(userID, content string) (*entity.Tweet, error) {
	newTweet := &entity.Tweet{
		ID:      generateUniqueID(),
		UserId:  userID,
		Content: content,
	}
	return newTweet, nil
}

func generateUniqueID() string {
	return "tweet123"
}
