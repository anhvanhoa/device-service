# Device Service

Microservice quản lý hệ sinh thái thiết bị IoT: loại thiết bị, thiết bị, lịch sử trạng thái và dữ liệu cảm biến. Viết bằng Go theo Clean Architecture, giao tiếp qua gRPC, lưu trữ PostgreSQL.

## 🏗️ Kiến trúc

```
device-service/
├── bootstrap/                 # Khởi tạo ứng dụng (app/env)
├── cmd/                       # Điểm vào ứng dụng và client test
│   ├── main.go                # Service chính
│   └── client/                # gRPC client CLI
├── domain/                    # Tầng nghiệp vụ (Clean Architecture)
│   ├── entity/                # Thực thể cốt lõi: device type, device, history, sensor data
│   ├── repository/            # Giao diện repository
│   └── usecase/               # Use cases cho từng domain
├── infrastructure/
│   ├── grpc_service/          # gRPC handlers cho các domain
│   └── repo/                  # Triển khai repository (PostgreSQL)
├── migrations/                # Schema DB + seed
├── script/seed/               # Script seed dữ liệu dev
├── doc/                       # Tài liệu nội bộ
└── logs/                      # Log ứng dụng
```

## 🚀 Tính năng

- **Device Type**: Tạo/đọc/cập nhật/xóa, liệt kê.
- **IoT Device**: Tạo/đọc/cập nhật/xóa, liệt kê, gắn `device_type`.
- **IoT Device History**: Ghi nhận và truy vấn lịch sử thay đổi/trạng thái, đọc/xóa, liệt kê.
- **Sensor Data**: Ghi dữ liệu cảm biến, đọc/xóa, liệt kê theo thiết bị/khoảng thời gian.
- **Phân trang/bộ lọc**: Hỗ trợ query linh hoạt theo ID thiết bị, loại, thời gian.
- **Rõ ràng tầng nghiệp vụ**: Use case thuần Go, dễ test và mở rộng.

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.25.0
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC (`github.com/anhvanhoa/sf-proto`)
- **Cấu hình**: Viper
- **Logging**: Zap

## 📋 Yêu cầu

- Go 1.25.0+
- PostgreSQL 12+
- [golang-migrate](https://github.com/golang-migrate/migrate)

## ⚙️ Cấu hình

Sao chép file cấu hình mẫu và chỉnh sửa:
```bash
cp dev.config.yml config.yml
```

Các khóa chính trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://pg:123456@localhost:5432/device-service_db?sslmode=disable"
name_service: "DeviceService"
port_grpc: 50057
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

## 🚀 Hướng dẫn nhanh

```bash
# 1) Clone & cài đặt
git clone <repository-url>
cd device-service
go mod download

# 2) Cơ sở dữ liệu
make create-db
make up          # chạy migrations

# 3) Chạy ứng dụng
make run         # service gRPC
make client      # client CLI để thử nhanh
```

Chạy bằng Docker Compose (tùy chọn):
```bash
docker-compose up -d
```

## 🗄️ Migrations & Seed

- Migrations nằm trong `migrations/` (bảng `device_types`, `iot_devices`, `iot_device_history`, `sensor_data`).
- Seed mẫu nằm trong `migrations/seed/` và `script/seed/`.

Lệnh hữu ích:
```bash
make up            # apply all migrations
make down          # rollback 1 step
make reset         # drop + up lại
make create name=migration_name
make force version=1
make seed          # chèn dữ liệu mẫu
make seed-reset    # reset + seed
make docker-seed   # seed khi chạy cùng Docker
```

## 🔌 gRPC API

Service triển khai các nhóm endpoint sau (tham chiếu proto từ `sf-proto`):

- **DeviceTypeService**: `Create`, `Get`, `Update`, `Delete`, `List`
- **IoTDeviceService**: `Create`, `Get`, `Update`, `Delete`, `List`
- **IoTDeviceHistoryService**: `Create`, `Get`, `Delete`, `List`
- **SensorDataService**: `Create`, `Get`, `Delete`, `List`

gRPC server được cấu hình tại `infrastructure/grpc_service/*/` và khởi tạo trong `bootstrap/app.go`.

## 🔧 Các lệnh Make chính

```bash
make build     # Build ứng dụng
make run       # Chạy service
make client    # Chạy gRPC client CLI
make test      # Chạy tests (nếu có)
make help      # Liệt kê lệnh
```

## 🧪 Thử nhanh qua client

```bash
make client
```

Client CLI trong `cmd/client/` cho phép gọi nhanh các RPC để kiểm thử luồng CRUD và truy vấn.

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Commit thay đổi + thêm test nếu cần
4. Mở pull request

## 📄 Giấy phép

MIT License

## 🆘 Hỗ trợ

Vui lòng mở issue khi gặp sự cố hoặc cần tính năng mới.
