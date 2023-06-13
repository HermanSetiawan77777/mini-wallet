package auth_test

import (
	"context"
	"fmt"
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/token"
	"herman-technical-julo/internal/wallet"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestAuthService_Authenticate(t *testing.T) {
	mockedErr := fmt.Errorf("mocked")
	convey.Convey("When failed to get wallet data", t, func() {
		service := auth.NewAuthService(
			&wallet.WalletServiceMock{
				GetByCustomerXidFunc: func(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
					return nil, mockedErr
				},
			},
			&token.TokenerMock[*auth.Session]{},
		)
		convey.Convey("Should return the error", func() {
			token, err := service.Authenticate(context.TODO(), "Xid")
			convey.So(token, convey.ShouldBeEmpty)
			convey.So(err, convey.ShouldEqual, mockedErr)
		})
	})

	convey.Convey("When failed to generate token", t, func() {
		service := auth.NewAuthService(
			&wallet.WalletServiceMock{
				GetByCustomerXidFunc: func(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
					return &wallet.Wallet{
						WalletId:    "1212",
						CustomerXid: "1212",
						StatusId:    1,
						Balance:     0,
					}, nil
				},
			},
			&token.TokenerMock[*auth.Session]{
				GenerateTokenFunc: func(ctx context.Context, payload *auth.Session, expiryTime *time.Duration) (string, error) {
					return "", mockedErr
				},
			},
		)

		convey.Convey("Should return the error", func() {
			token, err := service.Authenticate(context.TODO(), "haha")
			convey.So(token, convey.ShouldBeEmpty)
			convey.So(err, convey.ShouldEqual, mockedErr)
		})
	})

	convey.Convey("When success", t, func() {
		existingWallet := &wallet.Wallet{
			WalletId:    "1212",
			CustomerXid: "1212",
			StatusId:    1,
			Balance:     0,
		}
		service := auth.NewAuthService(
			&wallet.WalletServiceMock{
				GetByCustomerXidFunc: func(ctx context.Context, customerXid string) (*wallet.Wallet, error) {
					return existingWallet, nil
				},
			},

			&token.TokenerMock[*auth.Session]{},
		)

		convey.Convey("Should return token", func() {
			token, err := service.Authenticate(context.TODO(), "haha")
			convey.So(err, convey.ShouldBeNil)
			convey.So(token, convey.ShouldNotBeEmpty)
		})
	})
}
