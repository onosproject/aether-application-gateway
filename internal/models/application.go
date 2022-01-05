package models

// Application defines model for Application.
type Application struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Applications = []Application{
	{ID: "1", Name: "app-1"},
	{ID: "2", Name: "app-2"},
}
