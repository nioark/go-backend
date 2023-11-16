package main

import (
	api_usuarios "main/modules/domain/usuarios/delivery/http"
	usuarios_repo "main/modules/domain/usuarios/repository"
	usuarios_usecase "main/modules/domain/usuarios/usecase"

	api_pedidos "main/modules/domain/pedidos/delivery/http"
	pedidos_repo "main/modules/domain/pedidos/repository"
	pedidos_usecase "main/modules/domain/pedidos/usecase"

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

	usuariosRepo := usuarios_repo.New(db)
	usuariosUsecase := usuarios_usecase.New(usuariosRepo)
	api_usuarios.New(e, usuariosUsecase)

	pedidosRepo := pedidos_repo.New(db)
	pedidosUsecase := pedidos_usecase.New(pedidosRepo)
	api_pedidos.New(e, pedidosUsecase)

	e.Logger.Fatal(e.Start(":1323"))

}
