package responses

import (
	userModel "ayman-elmalah-build-a-good-structure-with-golang/internal/modules/user/models"
	"fmt"
)

type User struct {
	ID    uint
	Image string
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
		Name:  user.Name,
		Email: user.Email,
	}
}
