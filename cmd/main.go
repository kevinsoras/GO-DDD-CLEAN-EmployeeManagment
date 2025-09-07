// main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinsoras/employee-management/contexts/employee/interfaces"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	"github.com/kevinsoras/employee-management/shared/infrastructure/logger"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using system env vars")
	}

	// Inicializar logger
	appLogger := logger.New()

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		appLogger.Error("DB_URL not set in environment")
		os.Exit(1)
	}
	db.TestPostgresConnection(dsn)
	dbConn := db.NewPostgresConnection(dsn)

	// Inicializar controller con inyección de dependencias
	employeeController := interfaces.NewEmployeeController(dbConn, appLogger)

	// Inicializar API
	http.HandleFunc("/employee", employeeController.HandleRegister)

	// Iniciar servidor
	appLogger.Info("Server started", "port", 3000)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		appLogger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
