package accounts

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	repository "godating-dealls/internal/infra/mysql/repo"
)

type AccountEntityImpl struct {
	repository repository.AccountRepository
	validate   *validator.Validate
}

func NewAccountsEntityImpl(repository repository.AccountRepository, validate *validator.Validate) AccountEntity {
	return &AccountEntityImpl{repository: repository, validate: validate}
}

// SaveAccountEntities this is business rules enterprise of accounts
func (a AccountEntityImpl) SaveAccountEntities(ctx context.Context, tx *sql.Tx, dto domain.AccountDto) (domain.Accounts, error) {
	// validate request dto
	err := a.validate.Struct(dto)
	if err != nil {
		return domain.Accounts{}, err
	}

	// add validate username and email
	emailIsExist := a.repository.IsExistAccountByEmailFromDB(ctx, tx, *dto.Email)
	common.PrintJSON("auth entities | email is exist", emailIsExist)

	usernameIsExist := a.repository.IsExistAccountByUsernameFromDB(ctx, tx, *dto.Username)
	common.PrintJSON("auth entities | username is exist", usernameIsExist)

	if emailIsExist || usernameIsExist {
		return domain.Accounts{}, errors.New("email or username already exists")
	}

	records := record.AccountRecord{
		Username:     *dto.Username,
		PasswordHash: common.HashingPassword([]byte(dto.Password)),
		Email:        *dto.Email,
		Verified:     false,
	}
	common.PrintJSON("auth entities | account record to be saved", records)

	account, err := a.repository.CreateAccountToDB(ctx, tx, records)
	if err != nil {
		return domain.Accounts{}, err
	}

	result := domain.Accounts{
		AccountId: account.AccountID,
		Username:  account.Username,
		Email:     account.Email,
		Password:  account.PasswordHash,
		CreateAt:  account.CreatedAt,
		UpdateAt:  account.UpdatedAt,
	}
	return result, err
}

func (a AccountEntityImpl) AuthenticateAccount(ctx context.Context, tx *sql.Tx, dto domain.AccountDto) (domain.Accounts, error) {
	// validate request dto
	err := a.validate.Struct(dto)
	if err != nil {
		return domain.Accounts{}, err
	}

	if dto.Username != nil && *dto.Username != "" && dto.Email != nil && *dto.Email != "" {
		account, err := a.repository.FindAccountByUsernameAndEmailFromDB(ctx, tx, *dto.Email, *dto.Username)
		if err != nil {
			return domain.Accounts{}, err
		}
		res := domain.Accounts{
			AccountId: account.AccountID,
			Username:  account.Username,
			Email:     account.Email,
			Password:  account.PasswordHash,
			CreateAt:  account.CreatedAt,
			UpdateAt:  account.UpdatedAt,
		}
		return res, err
	} else if dto.Username != nil && *dto.Username != "" {
		account, err := a.repository.FindAccountByUsernameFromDB(ctx, tx, *dto.Username)
		if err != nil {
			return domain.Accounts{}, err
		}
		res := domain.Accounts{
			AccountId: account.AccountID,
			Username:  account.Username,
			Email:     account.Email,
			Password:  account.PasswordHash,
			CreateAt:  account.CreatedAt,
			UpdateAt:  account.UpdatedAt,
		}
		return res, err
	} else if dto.Email != nil && *dto.Email != "" {
		account, err := a.repository.FindAccountByEmailFromDB(ctx, tx, *dto.Email)
		if err != nil {
			return domain.Accounts{}, err
		}
		res := domain.Accounts{
			AccountId: account.AccountID,
			Username:  account.Username,
			Email:     account.Email,
			Password:  account.PasswordHash,
			CreateAt:  account.CreatedAt,
			UpdateAt:  account.UpdatedAt,
		}
		return res, err
	}

	return domain.Accounts{}, err
}

func (a AccountEntityImpl) FindAccountVerifiedEntities(ctx context.Context, tx *sql.Tx, accountId int64) (bool, error) {
	verified, err := a.repository.FindAccountVerifiedByAccountIdFromDB(ctx, tx, accountId)
	if err != nil {
		return false, errors.New("failed to find account verified entities")
	}
	return verified, nil
}

func (a AccountEntityImpl) UpdateAccountVerified(ctx context.Context, tx *sql.Tx, accountId int64) error {
	err := a.repository.UpdateAccountVerifiedByAccountIdFromDB(ctx, tx, accountId)
	if err != nil {
		return errors.New("failed to update account verified entities")
	}
	return nil
}

func (a AccountEntityImpl) FindAccountDetails(ctx context.Context, tx *sql.Tx, accountId int64) (domain.AccountDetail, error) {
	account, err := a.repository.FindAccountByIdFromDB(ctx, tx, accountId)
	if err != nil {
		return domain.AccountDetail{}, errors.New("failed to find account details")
	}
	result := domain.AccountDetail{
		AccountId: account.AccountID,
		Username:  account.Username,
		Email:     account.Email,
		Verified:  account.Verified,
	}
	return result, err
}
