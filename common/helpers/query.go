package helpers

type QueryListBase struct {
	Amount int `query:"amount"`
	Page   int `query:"page"`
}
