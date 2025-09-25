# Sinar API

A comprehensive food ordering and management system for universities built with Go and Gin framework.

## 🚀 Features

- **User Management**: Student registration and authentication with OTP
- **University Management**: Multi-university support
- **Restaurant Management**: University-specific restaurants
- **Food Management**: Food catalog and ordering system
- **User Food Orders**: Purchase tracking with expiration management
- **OTP Authentication**: Secure phone-based verification

## 📚 API Documentation

### Swagger UI
Access the interactive API documentation at: `http://localhost:8080/swagger/index.html`

### API Endpoints

#### 🔐 OTP Authentication
- `POST /otp/create` - Request OTP code
- `POST /otp/verify` - Verify OTP code

#### 👤 User Management
- `GET /user/{student_number}` - Get user by student number

#### 🏫 University Management
- `GET /university/{id}` - Get university by ID

#### 🍽️ Restaurant Management
- `GET /restaurant/{university_id}` - Get restaurants by university ID

#### 🍕 Food Management
- `GET /food/` - Get all available foods

#### 🛒 User Food Orders
- `GET /userfood/active` - Get active (non-expired) user foods
- `GET /userfood/{id}` - Get user food by ID
- `POST /userfood/` - Create new user food purchase
- `POST /userfood/{id}/use` - Mark food as used/expired

## 🛠️ Installation & Setup

### Prerequisites
- Go 1.24.5+
- PostgreSQL
- Redis

### Installation

1. **Clone the repository**
```bash
git clone <repository-url>
cd sinar
```

2. **Install dependencies**
```bash
go mod download
```

3. **Set up database**
```bash
# Create PostgreSQL database
createdb mydb

# Run database migrations
psql -d mydb -f Database/tables.sql
```

4. **Configure environment**
Create a `.env` file with your configuration:
```env
REDIS_HOST=localhost
REDIS_PORT=6379
SMS_API_KEY=your_sms_api_key
```

5. **Run the application**
```bash
go run main.go
```

The API will be available at `http://localhost:8080`

## 📖 API Usage Examples

### 1. Request OTP
```bash
curl -X POST http://localhost:8080/otp/create \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1234567890"}'
```

### 2. Verify OTP
```bash
curl -X POST http://localhost:8080/otp/verify \
  -H "Content-Type: application/json" \
  -d '{"phone": "+1234567890", "otp": "123456"}'
```

### 3. Get User by Student Number
```bash
curl http://localhost:8080/user/98123456
```

### 4. Get University
```bash
curl http://localhost:8080/university/1
```

### 5. Get Restaurants by University
```bash
curl http://localhost:8080/restaurant/1
```

### 6. Get All Foods
```bash
curl http://localhost:8080/food/
```

### 7. Create User Food Purchase (Single Object)
```bash
curl -X POST http://localhost:8080/userfood/ \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "food_id": 2,
    "restaurant_id": 1,
    "price": 85000,
    "sinar_price": 65000,
    "code": "ABC123",
    "expiration_hours": 24
  }'
```

### 8. Create User Food Purchase (Array Format)
```bash
curl -X POST http://localhost:8080/userfood/ \
  -H "Content-Type: application/json" \
  -d '[{
    "user_id": 1,
    "food_id": 2,
    "restaurant_id": 1,
    "price": 85000,
    "sinar_price": 65000,
    "code": "ABC123",
    "expires_at": "2025-09-24T10:15:30Z"
  }]'
```

### 9. Get Active User Foods
```bash
curl http://localhost:8080/userfood/active
```

### 10. Mark Food as Used
```bash
curl -X POST http://localhost:8080/userfood/1/use
```

## 🏗️ Project Structure

```
sinar/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── domain/           # Domain models
│   ├── dto/              # Data Transfer Objects
│   ├── interface/        # External interfaces
│   │   ├── postgres/     # Database repositories
│   │   ├── redis/        # Redis client
│   │   └── server/       # HTTP handlers
│   └── usecase/          # Business logic
├── pkg/                  # Public library code
│   ├── logger/           # Logging utilities
│   └── sms/              # SMS service
├── Database/             # Database schemas and migrations
├── docs/                 # Swagger documentation
├── main.go              # Application entry point
└── README.md            # This file
```

## 🔧 Development

### Generate Swagger Documentation
```bash
swag init
```

### Run Tests
```bash
go test ./...
```

### Build Application
```bash
go build -o sinar main.go
```

## 📝 Data Models

### User
- ID, FirstName, LastName, Phone, ProfilePic, StudentNum, Sex, UniversityID

### University
- ID, Name, Location, Logo

### Restaurant
- ID, UniversityID, Name, Sex, Color

### Food
- ID, Name

### UserFood
- ID, UserID, FoodID, RestaurantID, Price, SinarPrice, Code, CreatedAt, ExpiresAt

## 🚦 Status Codes

- `200 OK` - Success
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request
- `401 Unauthorized` - Authentication required
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource conflict (e.g., already used)
- `500 Internal Server Error` - Server error

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions, please contact the development team or create an issue in the repository.