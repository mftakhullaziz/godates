package auths

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain/auths"
	"godating-dealls/internal/infra/mysql/record"
	repository "godating-dealls/internal/infra/mysql/repo"
	"log"
)

type AccountEntitiesImpl struct {
	repository repository.AccountRepository
	validate   *validator.Validate
}

func NewAccountsEntitiesImpl(repository repository.AccountRepository, validate *validator.Validate) AuthEntities {
	return &AccountEntitiesImpl{repository: repository, validate: validate}
}

// SaveAccountEntities this is business rules enterprise of accounts
func (a AccountEntitiesImpl) SaveAccountEntities(ctx context.Context, tx *sql.Tx, dto auths.AccountDto) (auths.Accounts, error) {
	// validate request dto
	err := a.validate.Struct(dto)
	if err != nil {
		return auths.Accounts{}, err
	}

	// add validate username and email
	emailIsExist := a.repository.IsExistAccountByEmailFromDB(ctx, tx, dto.Email)
	usernameIsExist := a.repository.IsExistAccountByUsernameFromDB(ctx, tx, dto.Username)

	if emailIsExist || usernameIsExist {
		return auths.Accounts{}, errors.New("email or username already exists")
	}

	records := record.AccountRecord{
		Username:     dto.Username,
		PasswordHash: common.HashingPassword([]byte(dto.Password)),
		Email:        dto.Email,
		Verified:     false,
	}
	log.Printf("account record saved: %+v", records)

	account, err := a.repository.CreateAccountToDB(ctx, tx, records)
	if err != nil {
		return auths.Accounts{}, err
	}

	result := auths.Accounts{
		AccountId: account.AccountID,
		Username:  account.Username,
		Email:     account.Email,
		Password:  account.PasswordHash,
		CreateAt:  account.CreatedAt,
		UpdateAt:  account.UpdatedAt,
	}
	return result, err
}
