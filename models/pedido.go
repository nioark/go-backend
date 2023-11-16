package models

type Pedido struct {
	ID         uint64
	Name       string
	Quantidade uint
	UsuarioID  uint64
}
