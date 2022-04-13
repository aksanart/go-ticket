// running test
// go test .\test\userController_test.go -v
package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aksanart/go-ticket/controllers"
	"github.com/aksanart/go-ticket/entity"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockData struct {
	mock.Mock
}

func (m MockData) FindAll() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m MockData) FindUserByEmail(x string) (user1 entity.UserEntity, err error) {
	args := m.Called()
	return user1, args.Error(1)
}

func (m MockData) Save(entity.UserEntity) error {
	args := m.Called()
	return args.Error(0)
}

func TestUserControllerCreateUser(t *testing.T) {
	testCases := []struct {
		httpCode int
		nameFunc string
		request  interface{}
		expected string
	}{
		{
			httpCode: 200,
			nameFunc: "positive test",
			request:  `{"username": "aksan", "email": "aksan@gmail.com", "password": "aksan", "name": "aksan", "gender": "male"}`,
			expected: `{"code": 1000,"message": "success register user","status": "SUCCESS","value": null}`,
		},
		{
			httpCode: 400,
			nameFunc: "negative test 1",
			request:  `{"username": "aksan", "email": "aksan@gmail.com", "password": "aksan", "name": "aksan", "gender": ""}`,
			expected: `{"code":2001,"message":"Key: 'UserEntity.Gender' Error:Field validation for 'Gender' failed on the 'required' tag","status":"FAILED"}`,
		},
	}
	for _, v := range testCases {
		t.Run(v.nameFunc, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			userDBMOck := new(MockData)
			userDBMOck.On("Save", mock.Anything).Return(nil)
			var user entity.UserEntity
			json.Unmarshal([]byte(v.request.(string)), &user)
			data, err := json.Marshal(user)
			if err != nil {
				t.Fatal(err)
			}
			payload := bytes.NewBuffer(data)
			req, err := http.NewRequest("POST", "/register", payload)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			r := gin.Default()
			r.Use(controllers.CreateUser(userDBMOck))
			r.ServeHTTP(rr, req)
			log.Println(rr.Body.String())
			assert.Equal(t, v.httpCode, rr.Code)
			assert.JSONEq(t, v.expected, rr.Body.String())
		})
	}
}
