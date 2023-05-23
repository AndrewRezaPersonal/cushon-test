package testDatabase

import
(
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"types"
	"time"
)

func GetFunds(db *sql.DB) []types.Fund {
	var funds []types.Fund
	results, err := db.Query("SELECT ID, Description FROM cushon_test.funds WHERE IsAvailable = true")
	if err != nil {
		return funds
	}
	for results.Next() {
		var fund types.Fund
		err = results.Scan(&fund.ID, &fund.Description)
		if err != nil {
			return funds
		}
	funds = append(funds, fund)
	}
	return funds
}

func GetAuthorisation(db *sql.DB, customerID int) string {
	var authorisation string
	statement, err := db.Prepare("SELECT Authorisation FROM cushon_test.customers WHERE ID = ?")
	if err != nil {
		return ""
	}
	defer statement.Close()
	err = statement.QueryRow(customerID).Scan(&authorisation)
	if err != nil {
		return ""
	}
	return authorisation
}

func MakeDeposit(db *sql.DB, customerID int, deposit types.Deposit) bool {
	timeNow := time.Now()
	statement, err := db.Prepare("INSERT INTO cushon_test.deposits (CusomterID, FundID, Amount, CreatedAt) VALUES (?,?,?,?)")
	if err != nil {
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(customerID, deposit.Fund, deposit.Amount, timeNow.Format("2006-01-02 15:04:05"))
	if err != nil {
		return false
	}
	return true
}
