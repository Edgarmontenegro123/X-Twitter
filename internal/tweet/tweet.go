package tweet

type Tweet struct {
	ID      int
	UserId  int
	Content string
	Tweets  []Tweet
}

/*type Tweet struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}*/
