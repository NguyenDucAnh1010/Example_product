# Example_product

## Giới thiệu
1 dự án demo về việc sử dụng API để quản lý sản phẩm, cài đặt thêm tính năng đăng nhập bằng JWT và sử dụng WebSocket để cập nhật danh sách sản phẩm thời gian thực.

## Tính năng chính
- Tính năng 1: sử dụng API để quản lý sản phẩm
- Tính năng 2: tích hợp đăng nhập bằng JWT
- Tính năng 3: sử dụng WebSocket để cập nhật danh sách sản phẩm thời gian thực

## Yêu cầu hệ thống
- MongoDB >= 4.0 trên cổng 27017:27017 và không có tài khoản mật khẩu
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

## Test API (bằng Postman)
1. Tạo tài khoản bằng method Post trên "http://localhost:8080/register"
   ```body
   {
    "username": "admin",
    "password": "admin"
   }
2. Đăng nhập bằng method Post trên "http://localhost:8080/login"
   Body:
   {
    "username": "admin",
    "password": "admin"
   }
   copy token trả về vào Header
   Header:
   {
     "key": "Authorization",
     "value": "Token trả về khi đăng nhập"
   }
3. 
