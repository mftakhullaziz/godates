package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	repository "godating-dealls/internal/infra/mysql/repo"
	"time"
)

type UserEntityImpl struct {
	repository repository.UserRepository
	validate   *validator.Validate
}

func NewUserEntityImpl(repository repository.UserRepository, validate *validator.Validate) UserEntity {
	return &UserEntityImpl{repository: repository, validate: validate}
}

func (u UserEntityImpl) SaveUserEntities(ctx context.Context, tx *sql.Tx, dto domain.UserDto) error {
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
		FullName:    dto.FullName,
	}
	common.PrintJSON("user entities | user record to be saved", records)

	_, err = u.repository.CreateUserToDB(ctx, tx, records)
	err = common.HandleErrorDefault(err)
	return err
}

func (u UserEntityImpl) FindUserEntities(ctx context.Context, tx *sql.Tx, accountId int64) (domain.Users, error) {
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

func (u UserEntityImpl) FindAllUserEntities(ctx context.Context, tx *sql.Tx) ([]domain.AllUsers, error) {
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

func (u UserEntityImpl) FindAllUserViewsEntities(ctx context.Context, tx *sql.Tx, verified bool, shouldNext bool, accountIdIdentifier int64) ([]domain.AllUserViews, error) {
	var allUsers []record.UserAccountRecord
	if shouldNext {
		allUsers1, err := u.repository.GetAllUsersViewsFromDB(ctx, verified, accountIdIdentifier, tx)
		if err != nil {
			return nil, errors.New("could not get all users")
		}
		allUsers = append(allUsers, allUsers1...)
	} else {
		allUsers2, err := u.repository.GetAllUsersNextViewsFromDB(ctx, verified, accountIdIdentifier, tx)
		if err != nil {
			return nil, errors.New("could not get all users")
		}
		allUsers = append(allUsers, allUsers2...)
	}

	var allUser []domain.AllUserViews
	for _, user := range allUsers {
		usr := domain.AllUserViews{
			UserID:    user.UserID,
			AccountID: user.AccountID,
			Username:  user.Username,
			FullName:  user.FullName,
			Gender:    user.Gender,
			Age:       user.Age,
			Bio:       user.Bio,
			Verified:  user.Verified,
		}

		allUser = append(allUser, usr)
	}

	return allUser, nil
}

func (u UserEntityImpl) FindUserDetailEntity(ctx context.Context, tx *sql.Tx, accountId int64) (domain.Users, error) {
	user, err := u.repository.GetUserByAccountIdFromDB(ctx, tx, accountId)
	if err != nil {
		return domain.Users{}, err
	}

	usr := domain.Users{
		UserID:      user.UserID,
		AccountID:   user.AccountID,
		FullName:    user.FullName,
		DateOfBirth: user.DateOfBirth,
		Age:         user.Age,
		Gender:      user.Gender,
		Address:     user.Address,
		Bio:         user.Bio,
	}

	return usr, nil
}

func (u UserEntityImpl) UpdateUserEntities(ctx context.Context, tx *sql.Tx, dto domain.PatchUser) (domain.PatchUserDto, error) {
	err := u.validate.Struct(dto)
	if err != nil {
		return domain.PatchUserDto{}, err
	}

	dateOfBirth := parseDateOfBirth(*dto.DateOfBirth)
	rec := record.UserRecord{
		UserID:      dto.UserID,
		FullName:    dto.FullName,
		Gender:      *dto.Gender,
		Bio:         *dto.Bio,
		Address:     *dto.Address,
		DateOfBirth: &dateOfBirth,
		Age:         calculateAge(dateOfBirth),
	}

	user, err := u.repository.UpdateUserToDB(ctx, tx, rec)
	if err != nil {
		return domain.PatchUserDto{}, errors.New("could not update user")
	}

	patchUserDto := domain.PatchUserDto{
		UserID:      user.UserID,
		AccountID:   user.AccountID,
		FullName:    user.FullName,
		Age:         int64(user.Age),
		Gender:      &user.Gender,
		Address:     &user.Address,
		DateOfBirth: user.DateOfBirth,
		Bio:         &user.Bio,
		UpdatedAt:   &user.UpdatedAt,
	}

	return patchUserDto, nil
}

func calculateAge(dateOfBirth time.Time) int {
	if dateOfBirth.IsZero() {
		return 0
	}

	now := time.Now()
	years := now.Year() - dateOfBirth.Year()

	// Adjust the age if the birthdate hasn't occurred yet this year.
	if now.YearDay() < dateOfBirth.YearDay() {
		years--
	}

	return years
}

func parseDateOfBirth(date string) time.Time {
	layout := "2006-01-02"

	// Parse the date string
	dateOfBirth, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return dateOfBirth
}

func formatTimePointer(t *time.Time) string {
	if t == nil {
		return ""
	}
	// Format the time as YYYY-MM-DD
	return t.Format("2006-01-02")
}
