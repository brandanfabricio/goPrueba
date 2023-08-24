package database

import (
	"context"
	"ejercicio3/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Setting) (*sqlx.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)

	return sqlx.ConnectContext(ctx, "mysql", connectionString)

}
