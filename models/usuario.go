package models

type Usuario struct {
	Id       int64
	Name     string
	Password string
	Pedidos  []Pedido
}

type Pedido struct {
	ID         uint
	Name       string
	Quantidade uint
	UsuarioID  uint
}
