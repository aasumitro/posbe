package model

type (
	Addon struct {
		ID          int64   `json:"id"`
		Price       float64 `json:"price"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Usage       int64   `json:"usage"`
	}

	Product struct {
		ID            int64  `json:"id"`
		CategoryID    int64  `json:"category_id" `
		SubcategoryID int64  `json:"subcategory_id"`
		SKU           string `json:"sku"`
		Image         string `json:"image"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		// embedding data
		Category    *Category    `json:"category,omitempty"`
		Subcategory *Subcategory `json:"subcategory,omitempty"`
		Variants    []*Variant   `json:"variants,omitempty"`
	}

	Variant struct {
		ID          int64   `json:"id"`
		ProductID   int64   `json:"product_id"`
		UnitID      int64   `json:"unit_id"`
		UnitSize    float64 `json:"unit_size"`
		Type        string  `json:"type"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		// embedding data
		Unit *Unit `json:"unit,omitempty"`
	}

	//ICatalogCommonService interface {
	//	UnitList(ctx context.Context) (units []*Unit, errData *utils.ServiceError)
	//	AddUnit(ctx context.Context, data *Unit) (units *Unit, errData *utils.ServiceError)
	//	EditUnit(ctx context.Context, data *Unit) (units *Unit, errData *utils.ServiceError)
	//	DeleteUnit(ctx context.Context, data *Unit) *utils.ServiceError
	//
	//	CategoryList(ctx context.Context) (units []*Category, errData *utils.ServiceError)
	//	AddCategory(ctx context.Context, data *Category) (units *Category, errData *utils.ServiceError)
	//	EditCategory(ctx context.Context, data *Category) (units *Category, errData *utils.ServiceError)
	//	DeleteCategory(ctx context.Context, data *Category) *utils.ServiceError
	//
	//	SubcategoryList(ctx context.Context) (units []*Subcategory, errData *utils.ServiceError)
	//	AddSubcategory(ctx context.Context, data *Subcategory) (units *Subcategory, errData *utils.ServiceError)
	//	EditSubcategory(ctx context.Context, data *Subcategory) (units *Subcategory, errData *utils.ServiceError)
	//	DeleteSubcategory(ctx context.Context, data *Subcategory) *utils.ServiceError
	//
	//	AddonList(ctx context.Context) (units []*Addon, errData *utils.ServiceError)
	//	AddAddon(ctx context.Context, data *Addon) (units *Addon, errData *utils.ServiceError)
	//	EditAddon(ctx context.Context, data *Addon) (units *Addon, errData *utils.ServiceError)
	//	DeleteAddon(ctx context.Context, data *Addon) *utils.ServiceError
	//}

	//ICatalogProductService interface {
	//	AddProductVariant(ctx context.Context, data *ProductVariant) (variant *ProductVariant, errData *utils.ServiceError)
	//	EditProductVariant(ctx context.Context, data *ProductVariant) (variant *ProductVariant, errData *utils.ServiceError)
	//	DeleteProductVariant(ctx context.Context, data *ProductVariant) *utils.ServiceError
	//
	//	ProductSearch(ctx context.Context, keys []FindWith, values []any) (products []*Product, errData *utils.ServiceError)
	//	ProductList(ctx context.Context) (products []*Product, errData *utils.ServiceError)
	//	ProductDetail(ctx context.Context, id int) (product *Product, errData *utils.ServiceError)
	//	AddProduct(ctx context.Context, data *Product) (product *Product, errData *utils.ServiceError)
	//	EditProduct(ctx context.Context, data *Product) (product *Product, errData *utils.ServiceError)
	//	DeleteProduct(ctx context.Context, data *Product) *utils.ServiceError
	//}
)
