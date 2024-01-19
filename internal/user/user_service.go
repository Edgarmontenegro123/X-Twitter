package user

import (
	"errors"
)

type UserService struct {
	users map[int]User
	db    *Database
}

func NewUserService(db *Database) *UserService {
	return &UserService{
		users: make(map[int]User),
		db:    db,
	}
}

func (us *UserService) Follow(userID, followerID int) error {
	userToFollow, ok := us.users[userID]
	if !ok {
		return errors.New("el usuario a seguir no existe")
	}

	follower, ok := us.users[followerID]
	if !ok {
		return errors.New("El seguidor no existe")
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

	return nil
}

func (us *UserService) GetFollowers(userID int) ([]User, error) {
	user, ok := us.users[userID]
	if !ok {
		return nil, errors.New("el usuario no existe")
	}

	followers := []User{}
	for _, followerID := range user.Followers {
		follower, ok := us.users[followerID]
		if ok {
			followers = append(followers, follower)
		}
	}

	return followers, nil
}
