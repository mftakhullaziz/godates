package domain

import "time"

type AccountDataResponse struct {
	UserID      int64      `json:"user_id"`
	AccountID   int64      `json:"account_id"`
	FullName    *string    `json:"full_name"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Age         int        `json:"age"`
	Gender      string     `json:"gender"`
	Address     string     `json:"address"`
	Bio         string     `json:"bio"`
	Verified    bool       `json:"verified"`
	DateOfBirth *time.Time `json:"date_of_birth"`
}

type AccountViewResponse struct {
	TotalSwipeLike   int64 `json:"total_swipe_like"`
	TotalSwipePassed int64 `json:"total_swipe_passed"`
}

type AccountResponse struct {
	AccountData AccountDataResponse `json:"account_data"`
	AccountView AccountViewResponse `json:"account_view"`
}
