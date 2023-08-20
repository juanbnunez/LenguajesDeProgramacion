package main

import "fmt"

// Struct for representing a product
type product struct {
	name     string
	quantity int
	price    int
}

// A collection of products
type productInventory []product

var products productInventory
var lowStockProducts productInventory

const minimumStock int = 10 // The minimum stock level triggering decisions

// AddProduct adds a product to the inventory or updates its details
func (p *productInventory) AddProduct(name string, quantity int, price int) {
	found := false
	for n := 0; n < len(products); n++ {
		if name == (*p)[n].name {
			(*p)[n].quantity += quantity
			found = true
			if price != (*p)[n].price {
				(*p)[n].price = price
			}
		}
	}
	if !found {
		*p = append(*p, product{name: name, quantity: quantity, price: price})
	}
}

// IncrementProduct increases the quantity of an existing product
func (p *productInventory) IncrementProduct(name string, quantity int) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].name == name {
			(*p)[i].quantity += quantity
		}
	}
}

// FindProduct searches for a product by name and returns a pointer to it and an error status
func (p *productInventory) FindProduct(name string) (*product, int) {
	var err = -1
	var prod *product = nil
	for i := 0; i < len(*p); i++ {
		if (*p)[i].name == name {
			prod = &((*p)[i])
			err = 0
		}
	}
	return prod, err
}

// SellProduct handles the process of selling a product
func (p *productInventory) SellProduct(name string, quantity int) {
	prod, err := p.FindProduct(name)
	if err != -1 {
		if (*prod).quantity >= quantity {
			(*prod).quantity -= quantity
		}
		if (*prod).quantity == 0 {
			fmt.Printf("The product '%s' has been depleted and removed from the list.\n", prod.name)
			p.DeleteProduct(prod)
		} else {
			fmt.Println("Insufficient quantity for the sale")
		}
	}
}

// DeleteProduct removes a product from the inventory
func (p *productInventory) DeleteProduct(prod *product) {
	for i, pr := range *p {
		if pr.name == prod.name {
			*p = append((*p)[:i], (*p)[i+1:]...)
			break
		}
	}
}

// IncrementPrice updates the price of a product
func (p *productInventory) IncrementPrice(name string, price int) {
	prod, err := p.FindProduct(name)
	if err != -1 {
		(*prod).price = price
		fmt.Println("Price updated")
	}
}

// GetLowStockList retrieves a list of products with low stock
func (p *productInventory) GetLowStockList() productInventory {
	for _, pr := range *p {
		if pr.quantity <= minimumStock {
			lowStockProducts = append(lowStockProducts, pr)
		}
	}
	return lowStockProducts
}

// ReplenishLowStockInventory replenishes the inventory for products with low stock
func (p *productInventory) ReplenishLowStockInventory(minimumStockList productInventory) {
	for i := 0; i < len(minimumStockList); i++ {
		minimumStockList[i].quantity = minimumStock
	}
}

// FillData initializes the product inventory with sample data
func FillData() {
	products.AddProduct("rice", 15, 2500)
	products.AddProduct("beans", 4, 2000)
	products.AddProduct("milk", 8, 1200)
	products.AddProduct("coffee", 12, 4500)
	products.AddProduct("coffee", 12, 5000)
}

func main() {
	FillData()

	fmt.Println("Products:", products)

	// Sell products
	products.SellProduct("rice", 15)
	fmt.Println("Updated list:", products)

	// List low stock products
	products.GetLowStockList()
	fmt.Println("Low stock list:", lowStockProducts)

	// Replenish inventory
	products.ReplenishLowStockInventory(lowStockProducts)
	fmt.Println("Updated list:", lowStockProducts)
}
