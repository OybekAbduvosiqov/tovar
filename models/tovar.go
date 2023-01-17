package models

type TovarPrimeryKey struct {
	Id string `json:"id"`
}

type CreateTovar struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Photo       string   `json:"photo"`
	CategoryIds []string `json:"category_ids"`
}
type CategoryTovar struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Photo       string  `json:"photo"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
type Tovar struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	Photo       string       `json:"photo"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Categories  []*Category1 `json:"categories"`
}

type UpdateTovar struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Photo       string  `json:"photo"`
}

type UpdateTovarSwag struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Photo       string  `json:"photo"`
}

type GetListTovarRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListTovarResponse struct {
	Count  int64          `json:"count"`
	Tovars []*UpdateTovar `json:"tovars"`
}

type Empty struct{}
