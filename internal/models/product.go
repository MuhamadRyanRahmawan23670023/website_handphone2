package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MigrateDB(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		brand TEXT NOT NULL,
		model TEXT NOT NULL,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		description TEXT,
		image TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	// Insert sample data
	insertSampleData(db)
}

func insertSampleData(db *sql.DB) {
	// Check if data already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil || count > 0 {
		return
	}

	sampleProducts := []Product{
		{Name: "iPhone 15 Pro", Brand: "Apple", Model: "A3108", Price: 15999000, Stock: 25, Description: "Latest iPhone with titanium design and A17 Pro chip", Image: "https://images.pexels.com/photos/788946/pexels-photo-788946.jpeg"},
		{Name: "Samsung Galaxy S24 Ultra", Brand: "Samsung", Model: "SM-S928B", Price: 18999000, Stock: 30, Description: "Premium Android flagship with S Pen and 200MP camera", Image: "https://images.pexels.com/photos/1092644/pexels-photo-1092644.jpeg"},
		{Name: "Xiaomi 14 Ultra", Brand: "Xiaomi", Model: "2405CPX3DG", Price: 12999000, Stock: 20, Description: "Photography-focused flagship with Leica cameras", Image: "https://images.pexels.com/photos/1275229/pexels-photo-1275229.jpeg"},
		{Name: "Google Pixel 8 Pro", Brand: "Google", Model: "GC3VE", Price: 13999000, Stock: 15, Description: "AI-powered smartphone with pure Android experience", Image: "https://images.pexels.com/photos/1207583/pexels-photo-1207583.jpeg"},
		{Name: "OnePlus 12", Brand: "OnePlus", Model: "CPH2573", Price: 11999000, Stock: 18, Description: "Fast charging flagship with OxygenOS", Image: "https://images.pexels.com/photos/1279107/pexels-photo-1279107.jpeg"},
	}

	for _, product := range sampleProducts {
		_, err := db.Exec(`
			INSERT INTO products (name, brand, model, price, stock, description, image)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, product.Name, product.Brand, product.Model, product.Price, product.Stock, product.Description, product.Image)
		if err != nil {
			continue
		}
	}
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, brand, model, price, stock, description, image, created_at, updated_at FROM products ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Model, &p.Price, &p.Stock, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(db *sql.DB, id int) (*Product, error) {
	var p Product
	err := db.QueryRow("SELECT id, name, brand, model, price, stock, description, image, created_at, updated_at FROM products WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Brand, &p.Model, &p.Price, &p.Stock, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func CreateProduct(db *sql.DB, product *Product) error {
	result, err := db.Exec(`
		INSERT INTO products (name, brand, model, price, stock, description, image)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, product.Name, product.Brand, product.Model, product.Price, product.Stock, product.Description, product.Image)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)
	return nil
}

func UpdateProduct(db *sql.DB, id int, product *Product) error {
	_, err := db.Exec(`
		UPDATE products 
		SET name = ?, brand = ?, model = ?, price = ?, stock = ?, description = ?, image = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, product.Name, product.Brand, product.Model, product.Price, product.Stock, product.Description, product.Image, id)
	return err
}

func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}