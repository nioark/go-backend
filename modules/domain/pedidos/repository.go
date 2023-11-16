package pedidos

import "main/models"

type Repository interface {
	Fetch() ([]models.Pedido, error)
	GetPedido(ID uint64) (models.Pedido, error)
	GetUsuarioByPedido(pedidoId uint64) (models.Usuario, error)
	GetUsuario(usuarioID uint64) (models.Usuario, error)
	Add(pedido models.Pedido) (models.Pedido, error)
	Update(id uint64, name string, quantidade uint) error
	Remove(id uint64) error
}
