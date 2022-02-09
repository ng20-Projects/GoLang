package structs

type User struct {
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	Email         string  `json:"email"`
	Contact       string  `json:"contact"`
	City          string  `json:"city"`
	WalletBalance float32 `json:"walletBalance"`
	ID            int     `json:"id"`
}
