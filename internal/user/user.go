package user

type User struct {
	ID        int
	Username  string
	Followers []int
}

/*type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}*/
