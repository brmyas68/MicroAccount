package MySQLLayer

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type IAccountMySQLService interface {
	Insert(acc *IAccount) int32
	GetUserAccounts(userID int32) []IAccount
}

type AccountMySQLServiceStruct struct {
	sqldb *sql.DB
}

func NewAccountMySQLServiceStruct(sqldb *sql.DB) IAccountMySQLService {
	return &AccountMySQLServiceStruct{sqldb: sqldb}
}

func (tsAcc *AccountMySQLServiceStruct) GetUserAccounts(userID int32) []IAccount {

	ResultRows, err := tsAcc.sqldb.Query("SELECT * FROM account WHERE AccountUserID=?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer ResultRows.Close()

	Accounts := []IAccount{}
	for ResultRows.Next() {
		acc := IAccount{}
		err := ResultRows.Scan(&acc.AccountID, &acc.AccountUserID, &acc.AccountOrderID, &acc.AccountDateTime, &acc.AccountPrice, &acc.AccountTypePay)
		if err != nil {
			log.Fatal("  ==== ", err)
		}
		//acc.AccountDateTime = time.Now()
		Accounts = append(Accounts, acc)
	}

	return Accounts
}

func (tsAcc *AccountMySQLServiceStruct) Insert(acc *IAccount) int32 {

	ResultInsert, err := tsAcc.sqldb.Exec("INSERT INTO  account (AccountOrderID,AccountUserID,AccountDateTime,AccountPrice,AccountTypePay) VALUES  (?,?,?,?,?)", acc.AccountOrderID, acc.AccountUserID, acc.AccountDateTime, acc.AccountPrice, acc.AccountTypePay)
	if err != nil {
		return 0 //log.Fatal(err)
	}

	LastAccountID, err := ResultInsert.LastInsertId()
	if err != nil {
		return 0 //log.Fatal(err)
	}
	affectedRow, err := ResultInsert.RowsAffected()
	if err != nil {
		return 0 //log.Fatal(err)
	}

	if affectedRow > 0 {
		return int32(LastAccountID)
	}

	return 0
}
