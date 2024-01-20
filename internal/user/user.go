package user

type User struct {
	ID        int
	Username  string
	Followers []int
	// Tweets    []tweet.Tweet
}

/*type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}*/
