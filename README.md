# Proyek API Sederhana

Proyek ini adalah API dasar yang dibangun menggunakan Go dan mensimulasikan layanan item sederhana dengan beberapa endpoint untuk health check, informasi aplikasi, dan manajemen item. Proyek ini menggunakan router **Gorilla Mux** untuk routing dan **godotenv** untuk mengelola variabel lingkungan.

## Fitur

- **Endpoint Health check**: Mengembalikan status terkini dari API.
- **Endpoint Informasi**: Mengembalikan nama aplikasi, versi, lingkungan, dan port.
- **Manajemen Item**: Endpoint untuk membuat, melihat, dan mengambil item dari database sederhana yang disimulasikan di memori.

## Prasyarat

Pastikan Anda telah menginstal hal-hal berikut:

- [Go](https://golang.org/doc/install) (versi 1.23 atau lebih baru)
- [Git](https://git-scm.com/downloads)

## Instalasi

1. Clone repository:

   ```bash
   git clone https://newrahmat@bitbucket.org/newrahmat/devops-be-sample-manager.git
   cd devops-be-sample-manager
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Buat file `.env` untuk konfigurasi. Contoh:

   ```bash
   APP_NAME=SimpleAPI 
   APP_PORT=8080 
   APP_ENV=development 
   APP_VERSION=1.0.0
   ```

4. Menjalankan Proyek:

   ```bash
   go run main.go
   ```

### **POST** Request: Membuat Item Baru:

   ```bash
   curl -X POST http://localhost:8080/items \
   -H "Content-Type: application/json" \
   -d '{ "Name": "Item Baru" }' 
   ```

#### Respons:
   ```json
   { "ID": "2", "Name": "Item Baru" }
   ```

### **GET** Request: List Item:

   ```bash
   curl -X GET http://localhost:8080/items
   ```

#### Respons:
   ```json
   [ { "ID": "1", "Name": "Item Satu" }, { "ID": "2", "Name": "Item Baru" } ]
   ```

### **GET** Request: Info:

   ```bash
   curl -X GET http://localhost:8080/info
   ```

#### Respons:
   ```json
   { "appName": "SimpleAPI", "appPort": "8080", "appEnv": "development", "appVersion": "1.0.0" }
   ```

### **GET** Request: Health:

   ```bash
   curl -X GET http://localhost:8080/health
   ```

#### Respons:
   ```json
   { "status": "ok", "currentTime": "2024-10-23T12:34:56Z" }
   ```
