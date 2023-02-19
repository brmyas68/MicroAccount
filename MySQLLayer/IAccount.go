package MySQLLayer

type IAccount struct {
	AccountID       int32
	AccountUserID   int32
	AccountOrderID  int32
	AccountDateTime string //time.Time
	AccountPrice    int64
	AccountTypePay  int16
}

func NewAccount(AccountID int32, AccountUserID int32, AccountOrderID int32, AccountDateTime string, AccountPrice int64, AccountTypePay int16) *IAccount {

	return &IAccount{
		AccountID:       AccountID,
		AccountUserID:   AccountUserID,
		AccountOrderID:  AccountOrderID,
		AccountDateTime: AccountDateTime,
		AccountPrice:    AccountPrice,
		AccountTypePay:  AccountTypePay,
	}
}
