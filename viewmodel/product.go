package viewmodel

import (
	"html/template"

	"github.com/kenigbolo/go-web-app/model"
)

// ProductViewModel struct
type ProductViewModel struct {
	Title        string
	Active       string
	Alert        string
	AlertMessage string
	AlertDanger  string
	AlertSuccess string
	Product      Product
}

// Product struct
type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  template.HTML
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageURL         string
	ID               int
}

func productToVM(product *model.Product) Product {
	return Product{
		Name:             product.Name,
		DescriptionShort: product.DescriptionShort,
		DescriptionLong:  template.HTML(product.DescriptionLong),
		PricePerLiter:    product.PricePerLiter,
		PricePer10Liter:  product.PricePer10Liter,
		Origin:           product.Origin,
		IsOrganic:        product.IsOrganic,
		ImageURL:         product.ImageURL,
		ID:               product.ID,
	}
}

// NewProduct func
func NewProduct(product *model.Product) ProductViewModel {
	return ProductViewModel{
		Title:   "Lemonade Stand Supply - Shop",
		Active:  "shop",
		Product: productToVM(product),
		Alert:   "invisible",
	}
}
