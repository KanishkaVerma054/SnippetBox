package mocks

import (
	"KanishkaVerma054/snipperBox.dev/internal/models"
	"time"
)

/*
	// 14.5 Mocking dependencies: Mocking the database models

	// create a simple struct which implements the same methods as production models.SnippetModel,
	// but have the methods return some fixed dummy data instead.
*/
var mockSnippet = &models.Snippet {
	ID: 1,
	Title: "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func(m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func(m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func(m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}