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
