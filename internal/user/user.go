package user

type User struct {
	ID        int
	Username  string
	Followers []int
	Following []int
}
