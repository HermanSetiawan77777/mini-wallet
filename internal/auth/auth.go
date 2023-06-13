package auth

import (
	"context"
	"herman-technical-julo/internal/errors"
	"herman-technical-julo/internal/token"
	"herman-technical-julo/internal/wallet"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	WalletId    string
	CustomerXid string
}

func (s *Session) Validate() error {
	if s.WalletId == "" {
		return errors.ErrInvalidSession
	}
	return nil
}

type AuthIService interface {
	Authenticate(ctx context.Context, customerXid string) (token string, err error)
	Validate(ctx context.Context, token string) (session *Session, isValid bool, err error)
}

type AuthService struct {
	walletService wallet.WalletIService
	tokener       token.Tokener[*Session]
}

func NewAuthService(
	walletService wallet.WalletIService,
	tokener token.Tokener[*Session],

) *AuthService {
	return &AuthService{walletService, tokener}
}

func (s *AuthService) Authenticate(ctx context.Context, customerXid string) (token string, err error) {
	currentSession, err := s.walletService.GetByCustomerXid(ctx, customerXid)
	if err != nil {
		return "", err
	}
	if currentSession == nil {
		id := uuid.New()
		idString := id.String()
		err = s.walletService.InitializeWallet(ctx, &wallet.InitializeWalletParam{
			WalletId:    idString,
			CustomerXid: customerXid,
			StatusId:    1,
			Balance:     0,
		})
		if err != nil {
			return "", err
		}

		currentSession, err = s.walletService.GetByCustomerXid(ctx, customerXid)
		if err != nil {
			return "", err
		}
	}

	session := &Session{
		WalletId:    currentSession.WalletId,
		CustomerXid: currentSession.CustomerXid,
	}
	expiryTime := 24 * time.Hour
	token, err = s.tokener.GenerateToken(ctx, session, &expiryTime)

	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) Validate(ctx context.Context, token string) (session *Session, isValid bool, err error) {
	currentSession, err := s.tokener.Validate(ctx, token)
	if err != nil {
		return nil, false, err
	}
	if currentSession == nil {
		return nil, false, errors.ErrInvalidSession
	}
	if currentSession.Validate() != nil {
		return nil, false, errors.ErrInvalidSession
	}

	currentWallet, err := s.walletService.GetByCustomerXid(ctx, currentSession.CustomerXid)
	if err != nil {
		return nil, false, err
	}
	if currentWallet == nil {
		return nil, false, errors.ErrInvalidSession
	}

	return currentSession, true, nil
}
