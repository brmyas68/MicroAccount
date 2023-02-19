package AccountGrpcServer

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"main.go/MySQLLayer"
	"main.go/account/pb"
)

type AccountGrpcServerStruct struct {
}

var AccountMySQLStruct MySQLLayer.IAccountMySQLService

func NewAccountGrpcServerStruct(sqldb *sql.DB) *AccountGrpcServerStruct {

	AccountMySQLStruct = MySQLLayer.NewAccountMySQLServiceStruct(sqldb)
	return &AccountGrpcServerStruct{}
}

func (tsAcc *AccountGrpcServerStruct) InsertAccount(ctx context.Context, Request *pb.RequestAccount) (*pb.ResponseAccount, error) {

	TypePay, err := strconv.ParseInt(Request.Account.GetAccountTypePay(), 10, 64)
	if err != nil {
		panic(err)
	}

	IAccount := MySQLLayer.IAccount{
		AccountUserID:   Request.Account.GetAccountUserID(),
		AccountOrderID:  Request.Account.GetAccountOrderID(),
		AccountDateTime: Request.Account.GetAccountDateTime(), //Request.Account.GetAccountDateTime().AsTime(),
		AccountPrice:    Request.Account.GetAccountPrice(),
		AccountTypePay:  int16(TypePay),
	}
	AccountIDInsert := AccountMySQLStruct.Insert(&IAccount)
	if AccountIDInsert > 0 {
		return &pb.ResponseAccount{
			AccountID: AccountIDInsert,
			Status: &pb.StatusAccount{
				StatusCode:    pb.StatusCodeAccount_Status200,
				StatusMessage: pb.StatusMessageAccount_SUCCESS,
			},
		}, nil
	}

	return &pb.ResponseAccount{
		AccountID: 0,
		Status: &pb.StatusAccount{
			StatusCode:    pb.StatusCodeAccount_Status400,
			StatusMessage: pb.StatusMessageAccount_FAILED,
		},
	}, nil
}
func (tsAcc *AccountGrpcServerStruct) GetUserAccounts(Request *pb.RequestUserAccount, stream pb.AccountService_GetUserAccountsServer) error {
	Accounts := AccountMySQLStruct.GetUserAccounts(Request.GetUserID())
	if Accounts == nil || len(Accounts) <= 0 {
		return nil
	}
	for _, Account := range Accounts {
		//timestamp := timestamppb.New(Account.AccountDateTime)
		Acc := &pb.IAccount{
			AccountID:       Account.AccountID,
			AccountOrderID:  Account.AccountOrderID,
			AccountUserID:   Account.AccountUserID,
			AccountDateTime: Account.AccountDateTime,
			AccountPrice:    Account.AccountPrice,
			AccountTypePay:  fmt.Sprint(Account.AccountTypePay),
		}

		ResponseUserAccount := &pb.ResponseUserAccounts{
			Account: Acc,
			Status: &pb.StatusAccount{
				StatusCode:    pb.StatusCodeAccount_Status200,
				StatusMessage: pb.StatusMessageAccount_SUCCESS,
			},
		}

		err := stream.Send(ResponseUserAccount)
		if err != nil {
			return err
		}
	}

	return nil

}
