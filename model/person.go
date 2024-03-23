package model

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}
