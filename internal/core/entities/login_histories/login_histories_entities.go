package login_histories

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type LoginHistoriesEntities interface {
	SaveLoginHistoriesEntities(ctx context.Context, tx *sql.Tx, dto domain.LoginHistoriesDto) error
	UpdateLoginHistoriesEntities(ctx context.Context, tx *sql.Tx, dto domain.LoginHistoriesDto) error
}
