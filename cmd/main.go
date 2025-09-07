// main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/kevinsoras/employee-management/docs" // Importa los docs generados por Swag
	"github.com/joho/godotenv"
	"github.com/kevinsoras/employee-management/contexts/employee/interfaces"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	"github.com/kevinsoras/employee-management/shared/infrastructure/logger"
	httpSwagger "github.com/swaggo/http-swagger" // Importa el manejador de Swagger UI
)

// @title Employee Management API
// @version 1.0
// @description This is the API for managing employees.
// @host localhost:3000
// @BasePath /
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

	// Ruta para la documentación de Swagger
	http.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/swagger/doc.json")))

	// Iniciar servidor
	appLogger.Info("Server started", "port", 3000)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		appLogger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
