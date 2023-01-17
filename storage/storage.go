package storage

import (
	"context"

	"github.com/OybekAbduvosiqov/tovar/models"
)

type StorageI interface {
	CloseDB()
	Tovar() TovarRepoI
	// Category() CategoryRepoI
}

type TovarRepoI interface {
	Insert(context.Context, *models.CreateTovar) (string, error)
	GetByID(context.Context, *models.TovarPrimeryKey) (*models.Tovar, error)
	GetList(ctx context.Context, req *models.GetListTovarRequest) (*models.GetListTovarResponse, error)
	Update(ctx context.Context, book *models.UpdateTovar) error
	Delete(ctx context.Context, req *models.TovarPrimeryKey) error
}

// type CategoryRepoI interface {
// 	Insert(context.Context, *models.CreateCategory) (string, error)
// 	GetByID(context.Context, *models.CategoryPrimeryKey) (*models.Category, error)
// 	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
// 	Update(ctx context.Context, category *models.UpdateCategory) error
// 	Delete(ctx context.Context, req *models.CategoryPrimeryKey) error
// }
