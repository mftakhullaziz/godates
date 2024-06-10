package packages

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
	"time"
)

type PackageEntityImpl struct {
	PackagesRepository         repo.PackagesRepository
	PurchasePackagesRepository repo.PurchasePackagesRepository
}

func NewPackageEntityImpl(packagesRepository repo.PackagesRepository,
	purchasePackagesRepository repo.PurchasePackagesRepository) PackageEntity {
	return &PackageEntityImpl{
		PackagesRepository:         packagesRepository,
		PurchasePackagesRepository: purchasePackagesRepository,
	}
}

func (p PackageEntityImpl) GetAllPackagesEntity(ctx context.Context, tx *sql.Tx) ([]domain.PackageDto, error) {
	records, err := p.PackagesRepository.GetAllPackages(ctx, tx)
	if err != nil {
		return nil, errors.New("get all packages entity is failed")
	}
	var packages []domain.PackageDto
	for _, rec := range records {
		dto := domain.PackageDto{
			PackageID:                rec.PackageID,
			Description:              rec.Description,
			PackageName:              rec.PackageName,
			Price:                    rec.Price,
			PackageDurationInMonthly: rec.PackageDurationInMonthly,
			UnlimitedSwipes:          rec.UnlimitedSwipes,
			Status:                   rec.Status,
		}
		packages = append(packages, dto)
	}
	return packages, nil
}

func (p PackageEntityImpl) PurchasePackage(ctx context.Context, tx *sql.Tx, dto domain.PackageDto) error {
	rec := record.AccountPremiumRecord{
		AccountID:             dto.AccountID,
		PackageID:             dto.PackageID,
		PurchaseDate:          time.Now(),
		ExpiryDate:            time.Now().AddDate(0, int(dto.PackageDurationInMonthly), 0),
		UnlimitedSwipesActive: dto.UnlimitedSwipes,
		Status:                dto.Status,
	}
	err := p.PurchasePackagesRepository.PurchasePackagesByAccount(ctx, tx, rec)
	if err != nil {
		return errors.New("purchase packages by account is failed")
	}
	return nil
}

func (p PackageEntityImpl) FindAccountPremiumPackage(ctx context.Context, tx *sql.Tx, accountId int64) (domain.AccountPurchasePackage, error) {
	premiumPackage, err := p.PurchasePackagesRepository.FindAccountPremiumByAccountId(ctx, tx, accountId)
	if err != nil {
		return domain.AccountPurchasePackage{}, errors.New("find account premium package is failed")
	}
	res := domain.AccountPurchasePackage{
		AccountID: premiumPackage.AccountID,
		ExpiresIn: premiumPackage.ExpiryDate,
	}
	return res, nil
}
