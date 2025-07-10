# Go E-Commerce Handphone Shop

A modern e-commerce web application built with Go (Golang) for selling mobile phones. Features a clean, responsive design with complete CRUD operations, SQLite database, and report generation capabilities.

## Features

### Core Functionality
- **Create**: Add new products to the database
- **Read**: Display products with relational data
- **Update**: Edit existing product information
- **Delete**: Remove products from the database
- **Reports**: Generate PDF and Excel reports

### Additional Features
- Modern responsive design with Tailwind CSS
- SQLite database for data storage
- Admin dashboard for product management
- Sample product data with real mobile phone information
- Image support for product listings
- Stock management
- Search functionality (frontend ready)

## Tech Stack

- **Backend**: Go (Golang) with Gin framework
- **Database**: SQLite with SQL driver
- **Frontend**: HTML templates with Tailwind CSS
- **Reports**: PDF generation with gofpdf, Excel export with excelize

## Project Structure

```
├── main.go                 # Main application entry point
├── go.mod                  # Go module dependencies
├── internal/
│   ├── database/
│   │   └── db.go          # Database connection
│   ├── handlers/
│   │   ├── handlers.go    # HTTP handlers
│   │   └── reports.go     # Report generation handlers
│   └── models/
│       └── product.go     # Product model and database operations
├── templates/
│   ├── index.html         # Homepage template
│   ├── admin.html         # Admin dashboard
│   ├── admin_products.html # Product management
│   ├── product_form.html  # Add/Edit product form
│   └── reports.html       # Reports page
├── static/
│   └── css/
│       └── style.css      # Custom CSS styles
└── README.md
```

## Installation & Setup

1. **Prerequisites**:
   - Go 1.21 or higher
   - SQLite (included with go-sqlite3 driver)

2. **Clone and Install**:
   ```bash
   git clone <repository-url>
   cd handphone-shop
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```

4. **Access the Application**:
   - Homepage: http://localhost:8080
   - Admin Dashboard: http://localhost:8080/admin

## API Endpoints

### Public Endpoints
- `GET /` - Homepage with product listing
- `GET /products` - Get all products (JSON)
- `GET /products/:id` - Get product by ID (JSON)

### Admin Endpoints
- `GET /admin` - Admin dashboard
- `GET /admin/products` - Product management page
- `GET /admin/products/new` - Add new product form
- `POST /admin/products` - Create new product
- `GET /admin/products/:id/edit` - Edit product form
- `PUT /admin/products/:id` - Update product
- `DELETE /admin/products/:id` - Delete product
- `GET /admin/reports` - Reports page
- `GET /admin/reports/pdf` - Generate PDF report
- `GET /admin/reports/excel` - Generate Excel report

## Database Schema

The application uses a single `products` table with the following structure:

```sql
CREATE TABLE products (
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
```

## Sample Data

The application includes sample data for popular mobile phones:
- iPhone 15 Pro
- Samsung Galaxy S24 Ultra
- Xiaomi 14 Ultra
- Google Pixel 8 Pro
- OnePlus 12

## Features in Detail

### Product Management
- Add new products with complete information
- Edit existing products
- Delete products with confirmation
- View product details
- Stock level tracking

### Reports
- **PDF Reports**: Formatted PDF with product table
- **Excel Reports**: Spreadsheet export with all product data
- Real-time data generation
- Downloadable files with timestamps

### Admin Dashboard
- Product statistics overview
- Quick action buttons
- Recent products table
- Low stock alerts
- Responsive design

## Dependencies

```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/mattn/go-sqlite3 v1.14.17
    github.com/jung-kurt/gofpdf v1.16.2
    github.com/xuri/excelize/v2 v2.8.0
)
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the MIT License.

## Note

This application was created as a demonstration of Go web development with CRUD operations, database integration, and report generation. It's suitable for educational purposes and can be extended for production use with additional features like authentication, payment integration, and more advanced e-commerce functionality.