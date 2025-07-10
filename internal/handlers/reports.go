package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handphone-shop/internal/models"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

func (h *Handler) GeneratePDFReport(c *gin.Context) {
	products, err := models.GetAllProducts(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Handphone Shop - Products Report")
	pdf.Ln(20)

	// Table header
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(20, 10, "ID")
	pdf.Cell(40, 10, "Name")
	pdf.Cell(30, 10, "Brand")
	pdf.Cell(30, 10, "Model")
	pdf.Cell(30, 10, "Price")
	pdf.Cell(20, 10, "Stock")
	pdf.Ln(10)

	// Table data
	pdf.SetFont("Arial", "", 10)
	for _, product := range products {
		pdf.Cell(20, 10, fmt.Sprintf("%d", product.ID))
		pdf.Cell(40, 10, product.Name)
		pdf.Cell(30, 10, product.Brand)
		pdf.Cell(30, 10, product.Model)
		pdf.Cell(30, 10, fmt.Sprintf("%.0f", product.Price))
		pdf.Cell(20, 10, fmt.Sprintf("%d", product.Stock))
		pdf.Ln(10)
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=products-report-%s.pdf", time.Now().Format("2006-01-02")))
	
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) GenerateExcelReport(c *gin.Context) {
	products, err := models.GetAllProducts(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f := excelize.NewFile()
	sheetName := "Products"
	
	// Create a new worksheet
	index, err := f.NewSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set headers
	headers := []string{"ID", "Name", "Brand", "Model", "Price", "Stock", "Description", "Created At"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Set data
	for i, product := range products {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), product.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), product.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), product.Brand)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), product.Model)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), product.Price)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), product.Stock)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), product.Description)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), product.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// Set active sheet
	f.SetActiveSheet(index)

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=products-report-%s.xlsx", time.Now().Format("2006-01-02")))
	
	err = f.Write(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}