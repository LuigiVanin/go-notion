package dto

type CreateDocument struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Status string `json:"status" validate:"required" enum:"draft,completed,doing,rejected"`
}

type UpdateDocument struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Status string `json:"status" enum:"draft,completed,doing,rejected"`
}
