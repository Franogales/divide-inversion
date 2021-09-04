package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testdata struct {
	Investment int32 `json:"investment"`
}

func routetest() {

}

func TestApiCobranza(t *testing.T) {

	setTests := []struct {
		Name        string
		Input       int32
		StatusCode  int
		Body        string
		ExpectedErr string
	}{
		{
			Name:       "Igual que 50",
			Input:      50,
			StatusCode: 400,
			Body:       `{"error":"does not meet minimum requirements"}`,
		},
		{
			Name:       "Igual que 200",
			Input:      200,
			StatusCode: 400,
			Body:       `{"error":"does not meet minimum requirements"}`,
		},
		{
			Name:       "Igual que 300",
			Input:      300,
			StatusCode: 200,
			Body:       `{"credit_type_300":1,"credit_type_500":0,"credit_type_700":0}`,
		},
		{
			Name:       "Igual que 400",
			Input:      400,
			StatusCode: 400,
			Body:       `{"error":"the investment could not be assigned in any credit"}`,
		},
		{
			Name:        "Igual que 500",
			Input:       500,
			StatusCode:  200,
			Body:        `{"credit_type_300":0,"credit_type_500":1,"credit_type_700":0}`,
			ExpectedErr: "",
		},
		{
			Name:       "Igual que 600",
			Input:      600,
			StatusCode: 200,
			Body:       `{"credit_type_300":2,"credit_type_500":0,"credit_type_700":0}`,
		},
		{
			Name:       "Igual que 700",
			Input:      700,
			StatusCode: 200,
			Body:       `{"credit_type_300":0,"credit_type_500":0,"credit_type_700":1}`,
		},
		{
			Name:       "Igual que 800",
			Input:      800,
			StatusCode: 200,
			Body:       `{"credit_type_300":1,"credit_type_500":1,"credit_type_700":0}`,
		},
		{
			Name:       "Igual que 900",
			Input:      900,
			StatusCode: 200,
			Body:       `{"credit_type_300":3,"credit_type_500":0,"credit_type_700":0}`,
		},
		{
			Name:       "Igual que 1000",
			Input:      1000,
			StatusCode: 200,
			Body:       `{"credit_type_300":0,"credit_type_500":2,"credit_type_700":0}`,
		},
		{
			Name:       "Igual que 1100",
			Input:      1100,
			StatusCode: 200,
			Body:       `{"credit_type_300":2,"credit_type_500":1,"credit_type_700":0}`,
		},
		{
			Name:       "Igual que 1200",
			Input:      1200,
			StatusCode: 200,
			Body:       `{"credit_type_300":0,"credit_type_500":1,"credit_type_700":1}`,
		},
		{
			Name:       "Igual que 1300",
			Input:      1300,
			StatusCode: 200,
			Body:       `{"credit_type_300":2,"credit_type_500":0,"credit_type_700":1}`,
		},
		{
			Name:       "Igual que 3000",
			Input:      3000,
			StatusCode: 200,
			Body:       `{"credit_type_300":2,"credit_type_500":2,"credit_type_700":2}`,
		},
		{
			Name:       "Igual que 6700",
			Input:      6700,
			StatusCode: 200,
			Body:       `{"credit_type_300":2,"credit_type_500":1,"credit_type_700":8}`,
		},
		{
			Name:       "Igual que 0",
			Input:      0,
			StatusCode: 400,
			Body:       `{"error":"Bad parameter, expected 'investment' type integer"}`,
		},
	}
	for _, tt := range setTests {
		t.Run(tt.Name, func(t *testing.T) {
			router := setupRouter()
			a := assert.New(t)
			const url string = "/"

			td := testdata{
				Investment: tt.Input,
			}
			fmt.Println(td)
			reqBody, err := json.Marshal(td)
			if err != nil {
				a.Error(err)
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.StatusCode, w.Code)
			assert.Equal(t, tt.Body, w.Body.String())
		})
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/", Assign)
	return r
}
