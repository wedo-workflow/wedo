package model

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// MarshalBinary User implements the BinaryMarshaler interface
func (u *User) MarshalBinary() ([]byte, error) {
	if u == nil {
		return nil, errors.New("User is nil")
	}
	userBytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return userBytes, nil
}

// UnmarshalBinary User implements the BinaryUnmarshaler interface
func (u *User) UnmarshalBinary(data []byte) error {
	if u == nil {
		return errors.New("User is nil")
	}
	err := json.Unmarshal(data, u)
	if err != nil {
		return err
	}
	return nil
}

// UserListOptions is used to query users
type UserListOptions struct {
	Id     string `json:"id"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}
