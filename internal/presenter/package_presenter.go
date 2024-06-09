package presenters

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/packages"
	"godating-dealls/internal/domain"
	"net/http"
)

type PackagePresenter struct {
	w http.ResponseWriter
}

func NewPackagePresenter(w http.ResponseWriter) packages.BoundaryPackageOutput {
	return &PackagePresenter{w: w}
}

func (p PackagePresenter) PackageResponse(responses []domain.PackageResponse, err error) {
	common.HandleInternalServerError(err, p.w)
	if responses == nil {
		common.WriteJSONResponse(p.w, http.StatusOK, "Get packages successfully", domain.UserViewNilResponse{
			Message: "Package not found",
		}, int64(len(responses)))
	} else {
		common.WriteJSONResponse(p.w, http.StatusOK, "Get packages successfully", responses, int64(len(responses)))
	}
}

func (p PackagePresenter) PurchasePackageResponse(response domain.PurchasePackageResponse, err error) {
	common.HandleInternalServerError(err, p.w)
	common.WriteJSONResponse(p.w, http.StatusOK, "Purchase packages successfully", response, 1)
}
