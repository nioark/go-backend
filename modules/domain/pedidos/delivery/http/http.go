package api

import (
	"errors"
	"fmt"
	"log"
	"main/modules/domain/pedidos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type result struct {
	Message string `json:"message,omitempty"` // omit se vazio
	Data    any    `json:"data,omitempty"`    //omit se vario
	Error   string `json:"error,omitempty"`   //omit se vario
}

type handler struct {
	useCase pedidos.UseCase
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

func New(e *echo.Echo, u pedidos.UseCase) {
	h := &handler{useCase: u}
	e.GET("/pedidos/:id/usuario", h.GetUsuarioByPedido)
	e.GET("/pedidos/:id", h.GetPedido)
	e.GET("/pedidos", h.FetchPedidos)
	e.POST("/pedidos", h.AddPedido)
	e.PUT("/pedidos/:id", h.UpdatePedido)
	e.DELETE("/pedidos/:id", h.RemovePedido)
}

func (h handler) GetUsuarioByPedido(c echo.Context) error {
	idform := c.Param("id")

	idint, err := strconv.ParseUint(idform, 10, 64)

	if err != nil {
		return c.JSON(result{}.New("id invalido", nil, err))
	}

	usuario, err := h.useCase.GetUsuarioByPedido(idint)
	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar o usuario", nil, err))
	}

	return c.JSON(result{}.New("usuario encontrado com sucesso", usuario, nil))
}
func (h handler) GetPedido(c echo.Context) error {
	idform := c.Param("id")
	idint, err := strconv.ParseUint(idform, 10, 64)

	if err != nil {
		return c.JSON(result{}.New("id invalido", nil, err))
	}

	pedidos, err := h.useCase.GetPedido(idint)
	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar o pedido", nil, err))
	}

	return c.JSON(result{}.New("pedido encontrado com sucesso", pedidos, nil))
}

func (h handler) FetchPedidos(c echo.Context) error {
	pedidos, err := h.useCase.FetchPedidos()
	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar os pedidos", nil, err))
	}

	return c.JSON(result{}.New("lista de pedidos", pedidos, nil))
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

func parseUint(s string, fieldName string, err_list *[]error) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
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

func (h handler) AddPedido(c echo.Context) error {
	err_list := []error{}

	name := parseStr(c.FormValue("name"), "name", &err_list)
	quantidade := parseUint(c.FormValue("quantidade"), "quantidade", &err_list)
	usuario_id := parseId(c.FormValue("usuarioid"), "usuarioid", &err_list)

	if len(err_list) > 0 {
		return c.JSON(result{}.New(errToString(err_list), nil, errors.New("parametros invalidos")))
	}

	pedido, err := h.useCase.AddPedido(usuario_id, name, uint(quantidade))

	if err != nil {
		return c.JSON(result{}.New("não foi possivel adicionar o pedido", nil, err))
	}

	log.Printf("Pedido adicionado, Name: %s Quantidade: %d Usuario_ID: %d", name, quantidade, usuario_id)

	return c.JSON(result{}.New("pedido adicionado com sucesso", pedido, nil))
}

func (h handler) UpdatePedido(c echo.Context) error {
	err_list := []error{}

	id := parseId(c.Param("id"), "id", &err_list)
	name := parseStr(c.FormValue("name"), "name", &err_list)
	quantidade := parseUint(c.FormValue("quantidade"), "quantidade", &err_list)

	if len(err_list) > 0 {
		return c.JSON(result{}.New(errToString(err_list), nil, errors.New("parametros invalidos")))
	}

	pedido, err := h.useCase.UpdatePedido(id, name, uint(quantidade))
	if err != nil {
		return c.JSON(result{}.New("pedido não atualizado", nil, err))
	}

	log.Printf("Pedido atualizado, Pedido ID: %d", id)

	return c.JSON(result{}.New("pedido atualizado com sucesso", pedido, nil))
}

func (h handler) RemovePedido(c echo.Context) error {
	idform := c.Param("id")
	idint, err := strconv.ParseUint(idform, 10, 64)

	if err != nil {
		return c.JSON(result{}.New("id invalido", nil, err))
	}

	err = h.useCase.RemovePedido(idint)

	if err != nil {
		return c.JSON(result{}.New("pedido não removido", nil, err))
	}

	log.Printf("Pedido deletado, Pedido ID: %s", idform)

	return c.JSON(result{}.New("pedido removido com sucesso", map[string]interface{}{"id": idint}, nil))

}
