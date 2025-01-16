package controllers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goblogart/controllers"
	"goblogart/inits"
	"goblogart/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	user := models.User{}
	inits.DB.FirstOrCreate(&user, models.User{Email: "goblogart.test@gmail.com"})

	router.Use(func(c *gin.Context) {
		c.Set("user", user)
		c.Next()
	})

	router.POST("/", controllers.CreatePost)

	body := gin.H{
		"title":  "Test Title",
		"body":   "This is a test body",
		"likes":  10,
		"draft":  false,
		"author": "Test Author",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	data, exists := responseBody["data"].(map[string]interface{})
	assert.True(t, exists)
	assert.Equal(t, body["title"], data["Title"])
	assert.Equal(t, body["body"], data["Body"])
	assert.Equal(t, float64(body["likes"].(int)), data["Likes"])
	assert.Equal(t, body["draft"], data["Draft"])
	assert.Equal(t, body["author"], data["Author"])

	inits.DB.Exec("DELETE FROM posts")
}
