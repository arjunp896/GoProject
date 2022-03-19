package models

type Car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Model *Model `json:"model"`
	// ID string `json:"id"`

}

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
