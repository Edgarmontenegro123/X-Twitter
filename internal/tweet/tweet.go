package tweet

import "github.com/Edgarmontenegro123/X-Twitter/internal/user"

const MaxTweetLength = 280

type Tweet struct {
	ID      int
	UserId  int
	Content string
	Tweets  []Tweet
}

type TimeLineTweet struct {
	Tweet
	User user.User
}
