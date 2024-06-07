package login_histories

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
)

type LoginHistoriesEntitiesImpl struct {
	LoginRepository repo.LoginHistoriesRepository
	Validate        *validator.Validate
}

func NewLoginHistoriesEntitiesImpl(Validate *validator.Validate, LoginRepository repo.LoginHistoriesRepository) LoginHistoriesEntities {
	return &LoginHistoriesEntitiesImpl{
		Validate:        Validate,
		LoginRepository: LoginRepository,
	}
}

func (l LoginHistoriesEntitiesImpl) SaveLoginHistoriesEntities(ctx context.Context, tx *sql.Tx, dto domain.LoginHistoriesDto) error {
	// validate request dto
	err := l.Validate.Struct(dto)
	if err != nil {
		return err
	}

	records := record.LoginHistoriesRecord{
		UserID:    dto.UserID,
		AccountID: dto.AccountID,
	}
	_, err = l.LoginRepository.CreateLoginHistoryDB(ctx, tx, records)
	_ = common.HandleErrorDefault(err)
	return nil
}
