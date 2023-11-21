package models

type Pedido struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Quantidade uint   `json:"quantidade"`
	UsuarioID  uint64 `json:"usuarioid"`
}
