package main

import (
	"context"
	"ejercicio3/database"
	"ejercicio3/settings"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("Select * from users")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
