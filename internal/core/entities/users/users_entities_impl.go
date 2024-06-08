package users

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

type UserEntitiesImpl struct {
	repository repository.UserRepository
	validate   *validator.Validate
}

func NewUserEntitiesImpl(repository repository.UserRepository, validate *validator.Validate) UserEntities {
	return &UserEntitiesImpl{repository: repository, validate: validate}
}

func (u UserEntitiesImpl) SaveUserEntities(ctx context.Context, tx *sql.Tx, dto domain.UserDto) error {
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
		FullName:    *dto.FullName,
	}
	common.PrintJSON("user entities | user record to be saved", records)

	_, err = u.repository.CreateUserToDB(ctx, tx, records)
	err = common.HandleErrorDefault(err)
	return err
}

func (u UserEntitiesImpl) FindUserEntities(ctx context.Context, tx *sql.Tx, accountId int64) (domain.Users, error) {
	user, err := u.repository.GetUserByAccountIdFromDB(ctx, tx, accountId)
	if err != nil {
		return domain.Users{}, err
	}

	usr := domain.Users{
		UserID:    user.UserID,
		AccountID: user.AccountID,
	}

	return usr, nil
}

func (u UserEntitiesImpl) FindAllUserEntities(ctx context.Context, tx *sql.Tx) ([]domain.AllUsers, error) {
	allUsers, err := u.repository.GetAllUsersFromDB(ctx, tx)
	if err != nil {
		return nil, errors.New("could not get all users")
	}

	var allUser []domain.AllUsers
	for _, user := range allUsers {
		usr := domain.AllUsers{
			UserID:    user.UserID,
			AccountID: user.AccountID,
			Verified:  user.Verified,
		}

		allUser = append(allUser, usr)
	}

	return allUser, nil
}
