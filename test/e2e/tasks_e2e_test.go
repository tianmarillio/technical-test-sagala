package e2e_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	"github.com/tianmarillio/technical-test-sagala/src/models"
	"github.com/tianmarillio/technical-test-sagala/src/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})

	return db, nil
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	routes.RegisterRoutes(r, db)
	return r
}

func seedDatabase(db *gorm.DB) error {
	query := `
			INSERT INTO tasks (id, title, status, description, due_date) VALUES
			(1, 'Example Task 1', 'waiting_list', 'desc 1', '2024-09-01T10:00:00+07:00'),
			(2, 'Example Task 2', 'waiting_list', 'desc 2', '2024-09-02T10:00:00+07:00'),
			(3, 'Example Task 3', 'in_progress', 'desc 3', '2024-09-03T10:00:00+07:00'),
			(4, 'Example Task 4', 'done', 'desc 4', '2024-09-04T10:00:00+07:00')
	`
	if err := db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}

// POST /tasks
func TestPostTasks(t *testing.T) {
	// Setup db & router
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("could not set up test db: %v", err)
	}
	r := setupRouter(db)

	// Mock data
	jsonBody := make(map[string]interface{})
	jsonBody["title"] = "Test title"
	jsonBody["description"] = "Test description"
	jsonBody["status"] = "in_progress"
	jsonBody["due_date"] = "2024-09-04 10:00:00"

	jsonBodyEncoded, err := json.Marshal(jsonBody)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	jsonBodyBuffer := bytes.NewBuffer(jsonBodyEncoded)

	// HTTP Request
	req, _ := http.NewRequest(http.MethodPost, "/tasks", jsonBodyBuffer)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	// JSON Response
	var jsonResponse struct {
		TaskId uint `json:"task_id"`
	}
	err = json.Unmarshal(resp.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	assert.Equal(t, uint(1), jsonResponse.TaskId)

	// Validate DB record
	var exampleTask models.Task
	db.First(&exampleTask)

	expectedStatus := models.TaskStatusInProgress
	expectedDueDate, _ := now.Parse("2024-09-04 10:00:00")
	expectedDueDateUTC := expectedDueDate.UTC()
	actualDueDateUTC := exampleTask.DueDate.UTC()

	assert.Equal(t, "Test title", exampleTask.Title)
	assert.Equal(t, "Test description", exampleTask.Description)
	assert.Equal(t, expectedStatus, exampleTask.Status)
	assert.Equal(t, &expectedDueDateUTC, &actualDueDateUTC)
}

// GET /tasks
func TestGetTasks(t *testing.T) {
	// Setup db & router
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("could not set up test db: %v", err)
	}
	seedDatabase(db)
	r := setupRouter(db)

	// HTTP Request
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// JSON Response
	var jsonResponse []models.Task
	err = json.Unmarshal(resp.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	assert.Len(t, jsonResponse, 4, "expected result length equals to 4")

	exampleTask := jsonResponse[0]

	expectedStatus := models.TaskStatusWaitingList
	expectedDueDate, _ := now.Parse("2024-09-01 10:00:00")
	expectedDueDateUTC := expectedDueDate.UTC()
	actualDueDateUTC := exampleTask.DueDate.UTC()

	assert.Equal(t, uint(1), exampleTask.ID)
	assert.Equal(t, "Example Task 1", exampleTask.Title)
	assert.Equal(t, "desc 1", exampleTask.Description)
	assert.Equal(t, expectedStatus, exampleTask.Status)
	assert.Equal(t, &expectedDueDateUTC, &actualDueDateUTC)
}

// GET /tasks/:id
func TestGetTask(t *testing.T) {
	// Setup db & router
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("could not set up test db: %v", err)
	}
	seedDatabase(db)
	r := setupRouter(db)

	// HTTP Request
	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// JSON Response
	var jsonResponse models.Task
	err = json.Unmarshal(resp.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	expectedStatus := models.TaskStatusWaitingList
	expectedDueDate, _ := now.Parse("2024-09-01 10:00:00")
	expectedDueDateUTC := expectedDueDate.UTC()
	actualDueDateUTC := jsonResponse.DueDate.UTC()

	assert.Equal(t, uint(1), jsonResponse.ID)
	assert.Equal(t, "Example Task 1", jsonResponse.Title)
	assert.Equal(t, "desc 1", jsonResponse.Description)
	assert.Equal(t, expectedStatus, jsonResponse.Status)
	assert.Equal(t, &expectedDueDateUTC, &actualDueDateUTC)
}

// PATCH /:id
func TestPatchTasks(t *testing.T) {
	// Setup db & router
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("could not set up test db: %v", err)
	}
	seedDatabase(db)
	r := setupRouter(db)

	// Mock data
	jsonBody := make(map[string]interface{})
	jsonBody["title"] = "Test title"
	jsonBody["description"] = "Test description"
	jsonBody["status"] = "in_progress"
	jsonBody["due_date"] = "2024-09-04 10:00:00"

	jsonBodyEncoded, err := json.Marshal(jsonBody)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	jsonBodyBuffer := bytes.NewBuffer(jsonBodyEncoded)

	// HTTP Request
	req, _ := http.NewRequest(http.MethodPatch, "/tasks/1", jsonBodyBuffer)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// JSON Response
	var jsonResponse struct {
		TaskId uint `json:"task_id"`
	}
	err = json.Unmarshal(resp.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	assert.Equal(t, uint(1), jsonResponse.TaskId)

	// Validate DB record
	var exampleTask models.Task
	db.First(&exampleTask, 1)

	expectedStatus := models.TaskStatusInProgress
	expectedDueDate, _ := now.Parse("2024-09-04 10:00:00")
	expectedDueDateUTC := expectedDueDate.UTC()
	actualDueDateUTC := exampleTask.DueDate.UTC()

	assert.Equal(t, "Test title", exampleTask.Title)
	assert.Equal(t, "Test description", exampleTask.Description)
	assert.Equal(t, expectedStatus, exampleTask.Status)
	assert.Equal(t, &expectedDueDateUTC, &actualDueDateUTC)
}

// DELETE /tasks/:id
func TestDeleteTasks(t *testing.T) {
	// Setup db & router
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("could not set up test db: %v", err)
	}
	seedDatabase(db)
	r := setupRouter(db)

	// HTTP Request
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// JSON Response
	var jsonResponse struct {
		TaskId uint `json:"task_id"`
	}
	err = json.Unmarshal(resp.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	assert.Equal(t, uint(1), jsonResponse.TaskId)

	// Validate DB record
	var exampleTask models.Task
	db.First(&exampleTask, 1)

	assert.Zero(t, exampleTask.ID)
}
