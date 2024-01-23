package user

import (
	"errors"
)

type UserDB interface {
	SaveUser(user User)
	GetUserByID(userID int) User
}

type UserService struct {
	users map[int]User
	db    UserDB
}

func NewUserService(db UserDB) *UserService {
	return &UserService{
		users: make(map[int]User),
		db:    db,
	}
}

func (us *UserService) SaveUser(user User) {
	us.db.SaveUser(user)
}

func (us *UserService) Follow(userID, followerID int) error {
	userToFollow := us.db.GetUserByID(userID)
	if userToFollow.ID == 0 {
		return errors.New("el usuario a seguir no existe")
	}

	// Comprobar existencia del seguidor
	follower := us.db.GetUserByID(followerID)
	if follower.ID == 0 {
		return errors.New("el seguidor no existe")
	}

	// Verificar si el seguidor ya está siguiendo al usuario
	for _, existingFollower := range userToFollow.Followers {
		if existingFollower == followerID {
			return errors.New("el seguidor ya está siguiendo al usuario")
		}
	}

	// Agregar el seguidor al usuario
	userToFollow.Followers = append(userToFollow.Followers, followerID)

	// Almacenar el usuario actualizado en la DB
	us.db.SaveUser(userToFollow)

	// También actualizar la lista de seguidos en el seguidor
	follower.Following = append(follower.Following, userID)
	us.db.SaveUser(follower)

	return nil
}

func (us *UserService) GetFollowers(userID int) ([]User, error) {
	user := us.db.GetUserByID(userID)
	if user.ID == 0 {
		return nil, errors.New("el usuario no existe")
	}

	followers := []User{}
	for _, followerID := range user.Followers {
		follower := us.db.GetUserByID(followerID)
		if follower.ID != 0 {
			followers = append(followers, follower)
		}
	}

	return followers, nil
}
