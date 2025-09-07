package e2e_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/kevinsoras/employee-management/app"
	"github.com/kevinsoras/employee-management/shared/infrastructure/db"
	"github.com/kevinsoras/employee-management/shared/infrastructure/logger"
	"github.com/kevinsoras/employee-management/shared/utils"
)

var (
	testDB     *sql.DB
	testServer *httptest.Server
)

// TestMain runs setup and teardown for all E2E tests
func TestMain(m *testing.M) {
	ctx := context.Background()

	// Initialize logger for tests (optional, but good for debugging test setup)
	_ = logger.New() // This will set the default slog logger

	// Setup PostgreSQL container
	pgContainer, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:15-alpine"))
	if err != nil {
		slog.Error("Failed to start postgres container", "error", err)
		os.Exit(1)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		slog.Error("Failed to get connection string", "error", err)
		_ = pgContainer.Terminate(ctx)
		os.Exit(1)
	}

	// Connect to the test database
	testDB = db.NewPostgresConnection(connStr)
	if testDB == nil {
		slog.Error("Failed to connect to test database")
		_ = pgContainer.Terminate(ctx)
		os.Exit(1)
	}

	// Run migrations (assuming you have a way to run them programmatically)
	// For simplicity, we'll just run the up migrations for shared and employee contexts
	// In a real scenario, you'd use your migration tool (e.g., 'migrate' library)
	// This is a placeholder for actual migration execution
	// db.RunMigrations(testDB, "shared/infrastructure/persistence/migrations")
	// db.RunMigrations(testDB, "contexts/employee/infrastructure/persistence/migrations")

	// Assemble the application using the new app.NewApplication function
	appInstance := app.NewApplication(testDB, slog.Default())

	// Create httptest server
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is a simplified handler for the E2E test
		if r.URL.Path == "/employee" && r.Method == http.MethodPost {
			// Call the actual controller handler
			appInstance.EmployeeController.HandleRegister(w, r)
			return
		}
		http.NotFound(w, r)
	}))

	// Run tests
	exitCode := m.Run()

	// Teardown
	testServer.Close()
	_ = testDB.Close()
	_ = pgContainer.Terminate(ctx)

	os.Exit(exitCode)
}

func TestRegisterEmployeeE2E_Success(t *testing.T) {
	// Given
	reqBody := []byte(`{
		"person": {
			"type": "NATURAL",
			"email": "test.e2e@example.com",
			"phone": "+51987654321",
			"address": "Av. Test 123, Lima, Perú",
			"country": "Perú",
			"documentNumber": "12345678",
			"firstName": "Test",
			"lastNamePaternal": "E2E",
			"lastNameMaternal": "User",
			"birthDate": "1990-01-01T00:00:00Z",
			"gender": "M"
		},
		"employment": {
			"salary": 5000.00,
			"contractType": "INDEFINIDO",
			"startDate": "2024-01-01T00:00:00Z",
			"position": "QA Engineer",
			"workSchedule": "Full-time",
			"department": "Testing",
			"workLocation": "Remote",
			"bankAccount": "9876543210",
			"afp": "Habitat",
			"eps": "Pacifico",
			"hasCTS": true,
			"hasGratification": true,
			"hasVacation": true
		}
	}`)

	req, err := http.NewRequest(http.MethodPost, testServer.URL+"/employee", bytes.NewBuffer(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// When
	resp, err := testServer.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Then
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Optionally, read and assert response body
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var apiResp utils.APIResponse
	err = json.Unmarshal(respBody, &apiResp)
	require.NoError(t, err)
	assert.Equal(t, "success", apiResp.Status)
	assert.Contains(t, apiResp.Message, "Empleado registrado exitosamente")

	// Optional: Verify data in DB
	// This would require querying the testDB directly
	// For example: SELECT COUNT(*) FROM persons WHERE document_number = '12345678'
}
