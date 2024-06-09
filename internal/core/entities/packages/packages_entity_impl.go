package packages

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/repo"
)

type PackageEntityImpl struct {
	PackagesRepository repo.PackagesRepository
}

func NewPackageEntityImpl(packagesRepository repo.PackagesRepository) PackageEntity {
	return &PackageEntityImpl{
		PackagesRepository: packagesRepository,
	}
}

func (p PackageEntityImpl) GetAllPackagesEntity(ctx context.Context, tx *sql.Tx) ([]domain.PackageDto, error) {
	records, err := p.PackagesRepository.GetAllPackages(ctx, tx)
	if err != nil {
		return nil, errors.New("get all packages entity is failed")
	}
	var packages []domain.PackageDto
	for _, record := range records {
		dto := domain.PackageDto{
			PackageID:                record.PackageID,
			Description:              record.Description,
			PackageName:              record.PackageName,
			Price:                    record.Price,
			PackageDurationInMonthly: record.PackageDurationInMonthly,
			UnlimitedSwipes:          record.UnlimitedSwipes,
			Status:                   record.Status,
		}
		packages = append(packages, dto)
	}
	return packages, nil
}
