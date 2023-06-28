package model

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsLogin  bool   `json:"is_login"`
}

type Transaction struct {
	ID          string  `json:"id"`
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Merchant    string  `json:"merchant"`
	Amount      float64 `json:"amount"`
	Timestamp   string  `json:"timestamp"`
}

type TransactionsHistory struct {
	Transactions []Transaction `json:"transactions"`
}
