package tweet

import "github.com/Edgarmontenegro123/X-Twitter/internal/user"

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

/*type Tweet struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}*/
