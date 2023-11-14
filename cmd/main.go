package main

import (
	api "main/modules/domain/usuarios/delivery/http"
	"main/modules/domain/usuarios/repository"
	"main/modules/domain/usuarios/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Db connection
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Echo instance
	e := echo.New()
	e.Static("/", "public")

	usuariosRepo := repository.New(db)
	usuariosUsecase := usecase.New(usuariosRepo)
	api.New(e, usuariosUsecase)

	e.Logger.Fatal(e.Start(":1323"))

}
