package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type SwipesRepository interface {
	InsertSwipesToDB(ctx context.Context, tx *sql.Tx, record record.SwipeRecord) error
}
