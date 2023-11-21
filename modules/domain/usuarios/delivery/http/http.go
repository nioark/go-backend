package api

import (
	"errors"
	"fmt"
	"log"
	"main/modules/domain/usuarios"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type result struct {
	Message string `json:"message,omitempty"` // omit se vazio
	Data    any    `json:"data,omitempty"`    //omit se vario
	Error   string `json:"error,omitempty"`   //omit se vario
}

func (result) New(message string, data any, err error) (int, result) {
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, result{
			Message: message,
			Error:   err.Error(),
		}
	}

	return http.StatusOK, result{
		Message: message,
		Data:    data,
	}

}

type handler struct {
	useCase usuarios.UseCase
}

func New(e *echo.Echo, u usuarios.UseCase) {
	h := &handler{useCase: u}
	e.GET("/usuarios/:id/pedidos", h.GetUsuarioPedidos)
	e.GET("/usuarios/:id", h.GetUsuario)
	e.GET("/usuarios", h.FetchUsuarios)
	e.POST("/usuarios", h.AddUser)
	e.PUT("/usuarios/:id", h.UpdateUser)
	e.DELETE("/usuarios/:id", h.RemoveUser)
}

func (h handler) GetUsuarioPedidos(c echo.Context) error {

	idform := c.Param("id")

	idint, err := strconv.ParseUint(idform, 10, 64)

	//Id invalido
	if err != nil {
		return c.JSON(result{}.New("id de usuario invalido", nil, err))
	}

	pedidos, err := h.useCase.FetchPedidos(idint)

	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar os pedidos", nil, err))
	}

	return c.JSON(result{}.New("Lista de pedidos do usuario", pedidos, nil))
}

func (h handler) GetUsuario(c echo.Context) error {

	idform := c.Param("id")

	idint, err := strconv.ParseUint(idform, 10, 64)

	//Id invalido
	if err != nil {
		return c.JSON(result{}.New("id de usuario invalido", nil, err))
	}

	pedidos, err := h.useCase.GetUsuario(idint)

	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar o usuario", nil, err))
	}

	return c.JSON(result{}.New("Usuario encontrado com sucesso", pedidos, nil))
}

// Handler
func (h handler) FetchUsuarios(c echo.Context) error {
	usuarios, err := h.useCase.FetchUsuarios()
	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar os usuarios", nil, err))
	}

	return c.JSON(result{}.New("Lista de usuários", usuarios, nil))
}

func parseId(s string, fieldName string, err_list *[]error) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		*err_list = append(*err_list, fmt.Errorf("%s must be a number: %w", fieldName, err))
		return 0
	}

	if value == 0 {
		*err_list = append(*err_list, fmt.Errorf("%s cannot be zero", fieldName))
		return 0
	}
	return value
}

func parseUint(s string, fieldName string, err_list *[]error) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		*err_list = append(*err_list, fmt.Errorf("%s must be a number: %w", fieldName, err))
		return 0
	}

	return value
}

func parseStr(s string, fieldName string, err_list *[]error) string {
	if s == "" {
		*err_list = append(*err_list, fmt.Errorf("%s cannot be empty", fieldName))
		return ""
	}

	return s
}

func errToString(err_list []error) string {
	var errMsg string
	for _, err := range err_list {
		errMsg += err.Error() + "; "
	}
	return errMsg
}

func (h handler) AddUser(c echo.Context) error {
	err_list := []error{}

	username := parseStr(c.FormValue("name"), "name", &err_list)
	password := parseStr(c.FormValue("password"), "password", &err_list)

	if len(err_list) > 0 {
		return c.JSON(result{}.New(errToString(err_list), nil, errors.New("parametros invalidos")))
	}

	usuario, err := h.useCase.AddUser(username, password)

	if err != nil {
		return c.JSON(result{}.New("não foi possivel adicionar o usuario", nil, err))
	}

	log.Printf("Usuario adicionado, Username: %s Password: %s", username, password)

	return c.JSON(result{}.New("usuario adicionado com sucesso", usuario, nil))
}

func (h handler) UpdateUser(c echo.Context) error {
	err_list := []error{}

	id := parseId(c.Param("id"), "id", &err_list)
	username := parseStr(c.FormValue("name"), "name", &err_list)
	password := parseStr(c.FormValue("password"), "password", &err_list)

	if len(err_list) > 0 {
		return c.JSON(result{}.New(errToString(err_list), nil, errors.New("parametros invalidos")))
	}

	usuario, err := h.useCase.UpdateUser(id, username, password)

	if err != nil {
		return c.JSON(result{}.New("usuario não atualizado", nil, err))
	}

	log.Printf("Usuario atualizado, User ID: %d", id)

	return c.JSON(result{}.New("usuario atualizado com sucesso", usuario, nil))

}

func (h handler) RemoveUser(c echo.Context) error {
	idform := c.Param("id")

	idint, err := strconv.ParseUint(idform, 10, 64)

	//Id invalido
	if err != nil {
		return c.JSON(result{}.New("id invalido", nil, err))
	}

	err = h.useCase.RemoveUser(idint)

	//Usuario não removido
	if err != nil {
		return c.JSON(result{}.New("usuario não removido", nil, err))

	}

	log.Printf("Usuario deletado, User ID: %s", idform)

	return c.JSON(result{}.New("usuario removido com sucesso", nil, nil))

}
