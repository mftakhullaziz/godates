package domain

type SwipeRequest struct {
	ActionType     string `json:"action_type"`
	AccountIdSwipe int64  `json:"account_id_swipe"`
}

type SwipeResponse struct {
	Message string `json:"message"`
}
