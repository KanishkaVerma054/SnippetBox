package mocks

import "KanishkaVerma054/snipperBox.dev/internal/models"

/*
	// 14.5 Mocking dependencies: Mocking the database models

	// create a simple struct which implements the same methods as production models.UserModel,
	// but have the methods return some fixed dummy data instead.
*/

type UserModel struct{}

func(m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func(m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$woord" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func(m *UserModel) Exists(id int) (bool, error) {
	switch id  {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}