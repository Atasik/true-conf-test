package domain

import "time"

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}

	UserList map[int]User

	UserStore struct {
		Increment int      `json:"increment"`
		List      UserList `json:"list"`
	}

	UpdateUserInput struct {
		DisplayName *string
	}
)

func (i UpdateUserInput) Validate() error {
	if i.DisplayName == nil {
		return ErrNoUpdateValues
	}
	return nil
}
