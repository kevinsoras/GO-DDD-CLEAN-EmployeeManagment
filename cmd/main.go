// main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "github.com/kevinsoras/employee-management/docs" // Importa los docs generados por Swag
	"github.com/joho/godotenv"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	"github.com/kevinsoras/employee-management/shared/infrastructure/logger"
	httpSwagger "github.com/swaggo/http-swagger" // Importa el manejador de Swagger UI

	"github.com/kevinsoras/employee-management/app" // Importa el nuevo paquete 'app'
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

	// Ensamblar toda la aplicación
	application := app.NewApplication(dbConn, appLogger)

	// Inicializar API
	http.HandleFunc("/employee", application.EmployeeController.HandleRegister)

	// Ruta para la documentación de Swagger
	http.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/swagger/doc.json")))

	// Iniciar servidor
	appLogger.Info("Server started", "port", 3000)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		appLogger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
