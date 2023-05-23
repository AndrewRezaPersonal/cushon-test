package types

type Fund struct {
    ID int `json:"id"`
    Description string `json:"description"`
}

type Investment struct {
	CustomerID int `json:"customerID"`
	Deposits []Deposit `json:"deposits"`
}

type Deposit struct {
	Amount float64 `json:"amount"`
	Fund int `json:"fund"`
}