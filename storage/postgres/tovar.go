package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/OybekAbduvosiqov/tovar/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type TovarRepo struct {
	db *pgxpool.Pool
}

func NewTovarRepo(db *pgxpool.Pool) *TovarRepo {
	return &TovarRepo{
		db: db,
	}
}

func (r *TovarRepo) Insert(ctx context.Context, req *models.CreateTovar) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO tovar (
		id,
		name,
		description,
		price,
		photo,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Description,
		req.Price,
		req.Photo,
	)

	if err != nil {
		return "", err
	}

	if len(req.CategoryIds) > 0 {

		tovarCategoryQuery := `
				INSERT INTO tovar_category (
					category_id, 
					tovar_id
				) VALUES
		`

		for _, categoryId := range req.CategoryIds {
			tovarCategoryQuery += fmt.Sprintf("('%s', '%s'),", categoryId, id)
		}

		tovarCategoryQuery = tovarCategoryQuery[:len(tovarCategoryQuery)-1]

		_, err := r.db.Exec(ctx, tovarCategoryQuery)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func (r *TovarRepo) GetByID(ctx context.Context, req *models.TovarPrimeryKey) (*models.Tovar, error) {

	query := `
		SELECT
			tovar.id,
			tovar.name,
			tovar.description,
			tovar.price,
			tovar.photo,
			tovar.created_at,
			tovar.updated_at,
			(
				SELECT
					ARRAY_AGG(category_id)
				FROM tovar_category 
				WHERE tovar_category.tovar_id = $1
			) AS category_ids
		FROM
			tovar
		WHERE tovar.id = $1
	`

	var (
		id          sql.NullString
		name        sql.NullString
		description sql.NullString
		price       sql.NullFloat64
		photo       sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		categoryIds []string
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&description,
			&price,
			&photo,
			&createdAt,
			&updatedAt,
			(*pq.StringArray)(&categoryIds),
		)

	if err != nil {
		return nil, err
	}

	tovar := &models.Tovar{
		Id:          id.String,
		Name:        name.String,
		Description: description.String,
		Price:       price.Float64,
		Photo:       photo.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}

	if len(categoryIds) > 0 {
		categoryQuery := `
			SELECT
				id
			FROM
				category
			WHERE id IN (`

		for _, categoryId := range categoryIds {
			categoryQuery += fmt.Sprintf("'%s',", categoryId)
		}
		categoryQuery = categoryQuery[:len(categoryQuery)-1]
		categoryQuery += ")"

		rows, err := r.db.Query(ctx, categoryQuery)
		if err != nil {
			return nil, err
		}

		for rows.Next() {

			var (
				id sql.NullString
			)

			err = rows.Scan(
				&id,
			)

			tovar.Categories = append(tovar.Categories, &models.Category1{
				Id: id.String,
			})
		}
	}

	return tovar, nil
}

func (r *TovarRepo) GetList(ctx context.Context, req *models.GetListTovarRequest) (*models.GetListTovarResponse, error) {

	var (
		resp   models.GetListTovarResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			description,
			price,
			photo
		FROM tovar
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return &models.GetListTovarResponse{}, err
	}

	for rows.Next() {
		var tovar models.UpdateTovar

		err = rows.Scan(
			&resp.Count,
			&tovar.Id,
			&tovar.Name,
			&tovar.Description,
			&tovar.Price,
			&tovar.Photo,
			// &tovar.CreatedAt,
			// &tovar.UpdatedAt,
		)

		if err != nil {
			return &models.GetListTovarResponse{}, err
		}
		resp.Tovars = append(resp.Tovars, &tovar)

	}
	return &resp, nil
}

func (r *TovarRepo) Update(ctx context.Context, tovar *models.UpdateTovar) error {
	query := `
		UPDATE 
			tovar
		SET 
			name = $2,
			description = $4,
			price = $3,
			photo = $4,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		tovar.Id,
		tovar.Name,
		tovar.Description,
		tovar.Price,
		tovar.Photo,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *TovarRepo) Delete(ctx context.Context, req *models.TovarPrimeryKey) error {

	var (
		count int
	)

	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM tovar_category WHERE tovar_id = $1", req.Id).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err := r.db.Exec(ctx, "DELETE FROM tovar_category WHERE tovar_id  = $1 ", req.Id)
		if err != nil {
			return err
		}

	}
	_, err = r.db.Exec(ctx, "DELETE FROM tovar WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
