// main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Item struct untuk representasi objek di API
type Item struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}


// HealthResponse struct untuk representasi respons dari endpoint /health
type HealthResponse struct {
	Status      string    `json:"status"`
	CurrentTime time.Time `json:"currentTime"`
}

// InfoResponse struct untuk representasi respons dari endpoint /info
type InfoResponse struct {
	AppName string `json:"appName"`
	AppPort string `json:"appPort"`
	AppEnv  string `json:"appEnv"`
	AppVersion  string `json:"appVersion"`
}


// Simulasikan database sederhana
var items []Item

// GetItemsHandler mengembalikan daftar semua item
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetItemHandler mengembalikan item berdasarkan ID
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

// CreateItemHandler membuat item baru
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	newItem.ID = fmt.Sprintf("%d", len(items)+1)
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


// HealthCheckHandler menangani permintaan ke endpoint /health
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status:      "ok",
		CurrentTime: time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	infoResponse := InfoResponse{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppVersion:  os.Getenv("APP_VERSION"),
	}
	json.NewEncoder(w).Encode(infoResponse)
}

func main() {

	// Read environment variables
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")
	appEnv := os.Getenv("APP_ENV")
//	appVersion := os.Getenv("APP_VERSION")
	// Use default values if environment variables are not set
	if appName == "" {
		appName = "SimpleAPI"
	}
	if appPort == "" {
		appPort = "8080"
	}
	if appEnv == "" {
		appEnv = "development"
	}


	// Inisialisasi router menggunakan Mux
	router := mux.NewRouter()

	// Menambahkan route untuk endpoint-endpoint yang diinginkan
	// Menambahkan route untuk endpoint /health
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	// Menambahkan route untuk endpoint /info
	router.HandleFunc("/info", InfoHandler).Methods("GET")

	router.HandleFunc("/items", GetItemsHandler).Methods("GET")
	router.HandleFunc("/items/{id}", GetItemHandler).Methods("GET")
	router.HandleFunc("/items", CreateItemHandler).Methods("POST")

	// Mulai server pada port yang diatur dari variabel lingkungan
	fmt.Printf("%s is running in %s mode on http://localhost:%s\n", appName, appEnv, appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, router))

	// Mulai server pada port 8080
//	fmt.Println("Server berjalan di http://localhost:8080")
//	log.Fatal(http.ListenAndServe(":8080", router))
}

