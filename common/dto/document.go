package dto

type CreateDocument struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Status string `json:"status" validate:"required" enum:"draft,completed,doing,rejected"`
}
