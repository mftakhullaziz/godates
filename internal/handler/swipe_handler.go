package handler

import (
	"encoding/json"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/swipes"
	"godating-dealls/internal/domain"
	presenters "godating-dealls/internal/presenter"
	"net/http"
)

type SwipeHandler struct {
	InputSwipeBoundary swipes.InputSwipeBoundary
}

func NewSwipeHandler(InputSwipeBoundary swipes.InputSwipeBoundary) *SwipeHandler {
	return &SwipeHandler{InputSwipeBoundary: InputSwipeBoundary}
}

func (sh *SwipeHandler) SwipeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var request domain.SwipeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	common.PrintJSON("Handler | Swipe Request", request)

	presenter := presenters.NewSwipePresenter(w)

	// Call the use case method passing the presenter
	err := sh.InputSwipeBoundary.ExecuteSwipes(ctx, token, request, presenter)
	common.HandleInternalServerError(err, w)
}
