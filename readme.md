# Go Payment API

Go Payment API adalah sistem backend sederhana untuk mengelola pembayaran atau transfer uang antar users

## Fitur
- **Authentication**: Register, Login, Logout dengan JWT.
- **Payment**: Simpan dan kelola transaksi pembayaran.
- **Activity Logging**: Mencatat setiap aktivitas pengguna dalam sistem.

## Struktur Direktori
```
/go-payment-api
│-- handlers/       # Handler untuk request API
│-- helpers/        # Fungsi yang digunakan untuk membantu penanganan response
│-- middlewares/    # Middleware aplikasi
│-- routes/         # Routes / api path aplikasi
│-- vendor/         # Kumpulan dependensi
│-- models/         # Struktur data aplikasi serta method untuk struktur data
│-- storage/        # File JSON untuk penyimpanan
│-- utils/          # Utility functions
│-- main.go         # Entry point aplikasi
│-- go.mod          # Dependency management
│-- README.md       # Dokumentasi API
```

## Dokumentasi API

### 1. Register
**Endpoint**: `POST /register`

**Request Body:**
```json
{
  "username": "user1",
  "name": "User One",
  "password": "password123"
}
```
**Response:**
```json
{
  "status": 201,
  "message": "User Registered Successfully",
  "data": {
    "username": "user1",
    "name": "User One"
  }
}
```

### 2. Login
**Endpoint**: `POST /login`

**Request Body:**
```json
{
  "username": "user1",
  "password": "password123"
}
```
**Response:**
```json
{
  "status": 200,
  "message": "Login successful",
  "data": {
    "username": "user1",
    "name": "User One",
    "token": "your-jwt-token"
  }
}
```

### 3. Logout
**Endpoint**: `POST /logout`

**Request Header:**
```
Authorization: Bearer your-jwt-token
```
**Response:**
```json
{
  "status": 200,
  "message": "Logout successful"
}
```

### 4. Payment
**Endpoint**: `POST /payment`

**Request Header:**
```
Authorization: Bearer your-jwt-token
```

**Request Body:**
```json
{
  "amount": 50000,
  "payment_method": "credit_card"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Payment processed successfully",
  "data": {
    "transaction_id": "293012893",
    "amount": 50000,
    "payment_method": "credit_card",
    "status": "success"
  }
}
```

## Cara Menjalankan
1. Clone repositori:
   ```sh
   git clone https://github.com/IsraYehovaAsmaranda/go-payment-api.git
   cd go-payment-api
   ```
2. Install dependensi:
   ```sh
   go mod tidy
   ```
3. Jalankan aplikasi:
   ```sh
   go run main.go
   ```
4. API dapat diakses di `http://localhost:8080`

