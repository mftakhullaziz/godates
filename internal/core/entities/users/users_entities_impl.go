package users

import (
	"context"
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

func (u UserEntitiesImpl) SaveUserEntities(ctx context.Context, dto users.UserDto) error {
	// validate request dto
	err := u.validate.Struct(dto)
	if err != nil {
		return err
	}

	// Add validate find user by account id

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

	_, err = u.repository.CreateUserToDB(ctx, records)
	common.HandleErrorReturn(err)

	return nil

}
