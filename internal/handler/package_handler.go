package handler

import (
	"encoding/json"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/packages"
	"godating-dealls/internal/domain"
	presenters "godating-dealls/internal/presenter"
	"net/http"
)

type PackageHandler struct {
	InputPackageBoundary packages.InputPackageBoundary
}

func NewPackageHandler(inputPackageBoundary packages.InputPackageBoundary) *PackageHandler {
	return &PackageHandler{InputPackageBoundary: inputPackageBoundary}
}

func (ph *PackageHandler) GetPackageHandler(w http.ResponseWriter, r *http.Request) {
	// If user premium is unlimited, if not is just 10 data
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Instantiate the presenter
	presenter := presenters.NewPackagePresenter(w)

	// Call the use case method passing the presenter
	err := ph.InputPackageBoundary.ExecuteGetAllPackages(ctx, token, presenter)
	common.HandleInternalServerError(err, w)
}

func (ph *PackageHandler) PurchasePackages(w http.ResponseWriter, r *http.Request) {
	var request domain.PurchasePackageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	common.PrintJSON("Handler | Purchase Request", request)

	// If user premium is unlimited, if not is just 10 data
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Instantiate the presenter
	presenter := presenters.NewPackagePresenter(w)

	// Call the use case method passing the presenter
	err := ph.InputPackageBoundary.ExecutePurchasedPackages(ctx, token, request, presenter)
	common.HandleInternalServerError(err, w)
}
