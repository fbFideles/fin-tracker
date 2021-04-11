package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fbFideles/fin-tracker/routers/userRouter"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func UserTest(t *testing.T) {

	r := gin.Default()

	n := r.Group("api/v1/user")
	userRouter.UserRouter(n)

	var (
		singUpData = map[string]interface{}{
			"first_name": "asdfsdfa",
			"last_name":  "asdfasdfkj",
			"email":      "asdasd",
			"password":   "askdjaskd",
		}
	)

	t.Run("SingUpTest", func(t *testing.T) {
		m, err := json.Marshal(&singUpData)
		assert.Equal(t, err, nil)

		w := httptest.NewRecorder()
		http.NewRequest("POST", "/v1/public/user/singup", bytes.NewBuffer(m))

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
