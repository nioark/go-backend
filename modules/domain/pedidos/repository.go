package pedidos

import "main/models"

type Repository interface {
	FetchPedidos() ([]models.Pedido, error)
	GetPedido(ID uint64) (models.Pedido, error)
	GetUsuarioByPedido(pedidoId uint64) (models.Usuario, error)
	Add(pedido models.Pedido) (ID uint64, err error)
	Update(id uint64, name string, quantidade uint) error
	Remove(id uint64) error
}
