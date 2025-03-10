# Media Events API

This is a simple REST API built with Go and Gin to manage users, media, and events.

## 🚀 How to Run
### 1️⃣ Clone the repository
```sh
git clone https://github.com/your-username/media-events-api.git
cd media-events-api
```

### 2️⃣ Initialize Go module
```sh
go mod init media-events-api
```

### 3️⃣ Install dependencies
```sh
go mod tidy
```

### 4️⃣ Run the server
```sh
go run main.go
```

### 5️⃣ Test the API
#### Get all users
```sh
curl -X GET http://localhost:8080/users
```

#### Get a specific user by ID
```sh
curl -X GET http://localhost:8080/users/{uuid}
```

## 📌 Notes
- Data is stored in memory (no database required for now).
