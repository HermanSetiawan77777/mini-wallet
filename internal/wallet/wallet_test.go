package wallet_test

import (
	"context"
	"fmt"
	"herman-technical-julo/internal/wallet"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestWalletService_GetWalletByCustomerXid(t *testing.T) {
	convey.Convey("When failed to get wallet", t, func() {
		mockedErr := fmt.Errorf("mocked")
		service := wallet.NewWalletService(
			&wallet.WalletRepositoryMock{
				GetByCustomerXidFunc: func(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
					return nil, mockedErr
				},
			},
		)

		convey.Convey("It should return the error", func() {
			newWallet, err := service.GetByCustomerXid(context.TODO(), "Xid")
			convey.So(newWallet, convey.ShouldBeNil)
			convey.So(err, convey.ShouldEqual, mockedErr)
		})
	})

	convey.Convey("When wallet does not exists", t, func() {
		service := wallet.NewWalletService(
			&wallet.WalletRepositoryMock{
				GetByCustomerXidFunc: func(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
					return nil, nil
				},
			},
		)

		convey.Convey("It should return nil", func() {
			newWallet, err := service.GetByCustomerXid(context.TODO(), "Xid")
			convey.So(newWallet, convey.ShouldBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})

	convey.Convey("When wallet exists", t, func() {
		currentWallet := &wallet.Wallet{
			WalletId:    "1212",
			CustomerXid: "1212",
			StatusId:    1,
			Balance:     0,
		}
		service := wallet.NewWalletService(
			&wallet.WalletRepositoryMock{
				GetByCustomerXidFunc: func(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
					return currentWallet, nil
				},
			},
		)

		convey.Convey("It should return the wallet", func() {
			newWallet, err := service.GetByCustomerXid(context.TODO(), "Xid")
			convey.So(err, convey.ShouldBeNil)
			convey.So(newWallet, convey.ShouldEqual, currentWallet)
		})
	})
}

func TestWalletService_Create(t *testing.T) {
	mockedErr := fmt.Errorf("mocked")
	params := &wallet.InitializeWalletParam{
		WalletId:    "12121",
		CustomerXid: "12112",
		StatusId:    0,
		Balance:     0,
	}
	service := wallet.NewWalletService(&wallet.WalletRepositoryMock{})

	convey.Convey("When wallet id is empty", t, func() {
		convey.Convey("It should return error", func() {
			err := service.InitializeWallet(context.TODO(), &wallet.InitializeWalletParam{
				WalletId:    "",
				CustomerXid: "1212",
				StatusId:    0,
				Balance:     0,
			})
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(err.Error(), convey.ShouldEqual, "Wallet id cannot be empty")
		})
	})

	convey.Convey("When customer xid is empty", t, func() {
		convey.Convey("It should return error", func() {
			err := service.InitializeWallet(context.TODO(), &wallet.InitializeWalletParam{
				WalletId:    "1212",
				CustomerXid: "",
				StatusId:    0,
				Balance:     0,
			})
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(err.Error(), convey.ShouldEqual, "customer xid cannot be empty")
		})
	})

	convey.Convey("When failed to create wallet", t, func() {
		service := wallet.NewWalletService(
			&wallet.WalletRepositoryMock{
				CreateFunc: func(ctx context.Context, params *wallet.Wallet) error {
					return mockedErr
				},
			},
		)
		convey.Convey("It should return the error", func() {
			err := service.InitializeWallet(context.TODO(), params)
			convey.So(err, convey.ShouldResemble, mockedErr)
		})
	})

	convey.Convey("When success", t, func() {
		var createdWallet *wallet.Wallet
		service := wallet.NewWalletService(
			&wallet.WalletRepositoryMock{
				CreateFunc: func(ctx context.Context, params *wallet.Wallet) error {
					createdWallet = params
					return nil
				},
			},
		)
		convey.Convey("It should create wallet  based on given params", func() {
			err := service.InitializeWallet(context.TODO(), params)
			convey.So(err, convey.ShouldBeNil)

			convey.So(createdWallet, convey.ShouldNotBeNil)
			convey.So(createdWallet.WalletId, convey.ShouldEqual, params.WalletId)
			convey.So(createdWallet.Balance, convey.ShouldEqual, params.Balance)
			convey.So(createdWallet.CustomerXid, convey.ShouldEqual, params.CustomerXid)
			convey.So(createdWallet.StatusId, convey.ShouldEqual, params.StatusId)

		})
	})
}
