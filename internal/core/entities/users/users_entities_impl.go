package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain/users"
	"godating-dealls/internal/infra/mysql/record"
	repository "godating-dealls/internal/infra/mysql/repo"
	"log"
)

type UserEntitiesImpl struct {
	repository repository.UserRepository
	validate   *validator.Validate
}

func NewUserEntitiesImpl(repository repository.UserRepository, validate *validator.Validate) UserEntities {
	return &UserEntitiesImpl{repository: repository, validate: validate}
}

func (u UserEntitiesImpl) SaveUserEntities(ctx context.Context, tx *sql.Tx, dto users.UserDto) error {
	// validate request dto
	err := u.validate.Struct(dto)
	if err != nil {
		return err
	}

	// Add validate find user by account id
	userIsExist := u.repository.FindUserByUserIDFromDB(ctx, tx, dto.AccountID)
	if userIsExist {
		return errors.New("users already exists")
	}

	// Set default user data
	records := record.UserRecord{
		AccountID:   dto.AccountID,
		DateOfBirth: nil,
		Age:         0,
		Gender:      "",
		Address:     "",
		Bio:         "",
	}
	log.Printf("user record saved: %+v", records)

	_, err = u.repository.CreateUserToDB(ctx, tx, records)
	err = common.HandleErrorDefault(err)
	return err
}

func (u UserEntitiesImpl) FindUserEntities(ctx context.Context, tx *sql.Tx, accountId int64) (users.Users, error) {
	user, err := u.repository.GetUserByAccountIdFromDB(ctx, tx, accountId)
	if err != nil {
		return users.Users{}, err
	}

	usr := users.Users{
		UserID:    user.UserID,
		AccountID: user.AccountID,
	}

	return usr, nil
}
