package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/raflynagachi/kopi-santai-backend/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const GormDBPointerType = "*gorm.DB"

func ServeReq(opts *server.RouterConfig, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
	router := server.NewRouter(opts)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return router, rec
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func MockDB() *gorm.DB {
	db, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	return gormDB
}
