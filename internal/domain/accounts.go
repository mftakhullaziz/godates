package domain

type AccountDataResponse struct {
	UserID    int64   `json:"user_id"`
	AccountID int64   `json:"account_id"`
	FullName  *string `json:"full_name"`
	Username  string  `json:"username"`
	Age       int     `json:"age"`
	Gender    string  `json:"gender"`
	Address   string  `json:"address"`
	Bio       string  `json:"bio"`
	Verified  bool    `json:"verified"`
}

type AccountViewResponse struct {
	TotalView int64 `json:"total_view"`
}

type AccountResponse struct {
	AccountData AccountDataResponse `json:"account_data"`
	AccountView AccountViewResponse `json:"account_view"`
}
