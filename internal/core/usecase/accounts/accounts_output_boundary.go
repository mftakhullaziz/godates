package accounts

import "godating-dealls/internal/domain"

type OutputAccountBoundary interface {
	AccountDetailResponse(response domain.AccountResponse, err error)
}
