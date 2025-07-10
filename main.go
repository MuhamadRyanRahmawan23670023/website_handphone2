package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handphone-shop/internal/database"
	"github.com/handphone-shop/internal/handlers"
	"github.com/handphone-shop/internal/models"
)

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Auto-migrate tables
	models.MigrateDB(db)

	// Initialize Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static", "./static")

	// Initialize handlers
	h := handlers.New(db)

	// Routes
	router.GET("/", h.Home)
	router.GET("/products", h.GetProducts)
	router.GET("/products/:id", h.GetProduct)
	router.POST("/products", h.CreateProduct)
	router.PUT("/products/:id", h.UpdateProduct)
	router.DELETE("/products/:id", h.DeleteProduct)

	// Admin routes
	admin := router.Group("/admin")
	{
		admin.GET("/", h.AdminDashboard)
		admin.GET("/products", h.AdminProducts)
		admin.GET("/products/new", h.NewProductForm)
		admin.POST("/products", h.CreateProduct)
		admin.GET("/products/:id/edit", h.EditProductForm)
		admin.PUT("/products/:id", h.UpdateProduct)
		admin.DELETE("/products/:id", h.DeleteProduct)
		admin.GET("/reports", h.ReportsPage)
		admin.GET("/reports/pdf", h.GeneratePDFReport)
		admin.GET("/reports/excel", h.GenerateExcelReport)
	}

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}