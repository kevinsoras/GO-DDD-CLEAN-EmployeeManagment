// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinsoras/employee-management/contexts/employee/interfaces"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL not set in environment")
	}
	db.TestPostgresConnection(dsn)
	dbConn := db.NewPostgresConnection(dsn)

	// Inicializar controller con inyección de dependencias
	employeeController := interfaces.NewEmployeeController(dbConn)

	// Inicializar API
	http.HandleFunc("/employee", employeeController.HandleRegister)

	// Iniciar servidor
	log.Println("Server started on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
