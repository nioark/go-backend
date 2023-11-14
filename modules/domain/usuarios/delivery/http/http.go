package api

import (
	"log"
	"main/modules/domain/usuarios"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type result struct {
	Message string `json:"message,omitempty"` // omit se vazio
	Data    any    `json:"data,omitempty"`    //omit se vario
	Error   error  `json:"error,omitempty"`   //omit se vario
}

func (result) New(message string, data any, err error) (int, result) {
	if err != nil {
		return http.StatusInternalServerError, result{
			Message: message,
			Error:   err,
		}
	}

	return http.StatusOK, result{
		Message: message,
		Error:   err,
		Data:    data,
	}

}

type handler struct {
	useCase usuarios.UseCase
}

func New(e *echo.Echo, u usuarios.UseCase) {
	h := &handler{useCase: u}
	e.GET("/", h.home)
	e.POST("/add", h.add_user)
	e.POST("/update", h.update_user)
	e.POST("/remove", h.remove_user)
}

// Handler
func (h handler) home(c echo.Context) error {
	usuarios, err := h.useCase.Fetch()
	if err != nil {
		return c.JSON(result{}.New("não foi possivel recuperar os usuarios", nil, err))
	}

	return c.JSON(result{}.New("Lista de usuários", usuarios, nil))
}

func (h handler) add_user(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	usuario, err := h.useCase.AddUser(username, password)

	if err != nil || usuario.Id == 0 {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	log.Printf("Usuario adicionado, Username: %s Password: %s", username, password)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (h handler) update_user(c echo.Context) error {
	idform := c.FormValue("id")
	username := c.FormValue("username")
	password := c.FormValue("password")

	idint, err := strconv.ParseInt(idform, 10, 64)

	//Id invalido
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	usuario, err := h.useCase.UpdateUser(idint, username, password)

	//Usuario não removido
	//		      || todo mover erro pra usecase
	if err != nil || usuario.Id == 0 {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	log.Printf("Usuario atualizado, User ID: %s", idform)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

func (h handler) remove_user(c echo.Context) error {
	idform := c.FormValue("id")

	idint, err := strconv.ParseInt(idform, 10, 64)

	//Id invalido
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	err = h.useCase.RemoveUser(idint)

	//Usuario não removido
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	log.Printf("Usuario deletado, User ID: %s", idform)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

/*
func (h handler) gerar_tabela(id, username, password string) string {
	tab := fmt.Sprintf(`
	<tr>
		<th scope='row'>%s</th>
		<td>%s</td>
		<td>%s</td>
	</tr>
	`, id, username, password)

	return tab
}
*/

//var results []map[string]interface{}

/*
	var tabs string

	for _, j := range results {
		id := fmt.Sprint(j["id"])
		username := j["name"].(string)
		password := j["password"].(string)

		tab := gerar_tabela(id, username, password)
		tabs += tab
	}

	return c.Render(http.StatusOK, "index.html", tabs)*/
