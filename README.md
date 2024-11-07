# Example_product

## Giới thiệu
1 dự án demo về việc sử dụng API để quản lý sản phẩm, cài đặt thêm tính năng đăng nhập bằng JWT và sử dụng WebSocket để cập nhật danh sách sản phẩm thời gian thực.

## Tính năng chính
- Tính năng 1: sử dụng API để quản lý sản phẩm
- Tính năng 2: tích hợp đăng nhập bằng JWT
- Tính năng 3: sử dụng WebSocket để cập nhật danh sách sản phẩm thời gian thực

## Yêu cầu hệ thống
- MongoDB >= 4.0 trên cổng 27017:27017 và không có tài khoản mật khẩu (em cài trên Docker)
- Golang

## Cài đặt
1. Clone repository:
   ```bash
   git clone https://github.com/NguyenDucAnh1010/Example_product.git
2. Tải thư viện xuống:
   ```bash
   go mod tidy
4. Khởi chạy ứng dụng:
   ```bash
   go run cmd/main.go

## Database schema của sản phẩm
	type Product struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Name        string             `json:"name,omitempty" bson:"name,omitempty"`
		Description string             `json:"description,omitempty" bson:"description,omitempty"`
		Category    string             `json:"category,omitempty" bson:"category,omitempty"`
		Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
		Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
		Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
		Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
		CreatedAt   primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
		UpdatedAt   primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	}

## Test API (bằng Postman)
1. Tạo tài khoản bằng method Post trên "http://localhost:8080/register" với Body:
   ```body
   {
    "username": "admin",
    "password": "admin"
   }
2. Đăng nhập bằng method Post trên "http://localhost:8080/login" với Body:
   ```body
   {
      "username": "admin",
      "password": "admin"
   }
3. Copy token trả về vào Header:
   ```header
   {
     "key": "Authorization",
     "value": "Token trả về khi đăng nhập"
   }
4. Kết nối tới Websoket trên "ws://localhost:8080/ws" (cái này không cần đăng nhập để vào)
5. Tạo 1 sản phẩm mới bằng method Post trên "http://localhost:8080/product" với Body:
   ```body
   {
      "name": "Laptop XI tặng",
      "description": "Laptop chuyên game với cấu hình mạnh mẽ",
      "category": "Electronics",
      "price": 1200.99,
      "stock": 50,
      "images": [
         "https://example.com/images/laptop1.jpg",
         "https://example.com/images/laptop2.jpg"
      ],
      "tags": ["gaming", "laptop", "electronics"],
      "created_at": "2024-11-01T08:30:00Z",
      "updated_at": "2024-11-01T08:30:00Z"
   }
6. Kiểm tra tất cả sản phẩm bằng method Get trên "http://localhost:8080/product"
7. Kiểm tra 1 sản phẩm theo id bằng method Get trên "http://localhost:8080/product/{_id}" (_id của sản phẩm mới tạo)
8. Cập nhật 1 sản phẩm theo id bằng method Put trên "http://localhost:8080/product/{_id}" (_id của sản phẩm mới tạo) với Body:
   ```body
   {
      "name": "Laptop XI sửa",
      "description": "Laptop chuyên game với cấu hình mạnh mẽ",
      "category": "Electronics",
      "price": 1200.99,
      "stock": 50,
      "images": [
         "https://example.com/images/laptop1.jpg",
         "https://example.com/images/laptop2.jpg"
      ],
      "tags": ["gaming", "laptop", "electronics"],
      "created_at": "2024-11-01T08:30:00Z",
      "updated_at": "2024-11-01T08:30:00Z"
   }
9. Xoá 1 sản phẩm theo id bằng method Delete trên "http://localhost:8080/product/{_id}" (_id của sản phẩm mới tạo)
