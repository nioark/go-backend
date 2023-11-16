package repository

import (
	"main/models"
	"main/modules/domain/pedidos"

	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) pedidos.Repository {
	return &repository{conn: conn}
}

func (r repository) GetPedido(ID uint64) (pedido models.Pedido, err error) {
	return pedido, r.conn.First(&pedido, ID).Error
}

func (r repository) GetUsuario(usuarioID uint64) (usuario models.Usuario, err error) {
	return usuario, r.conn.First(&usuario, usuarioID).Error
}

func (r repository) GetUsuarioByPedido(pedidoId uint64) (usuario models.Usuario, err error) {
	pedido, err_j := r.GetPedido(pedidoId)

	if err_j != nil {
		return models.Usuario{}, err
	}

	return usuario, r.conn.Where("id = ?", pedido.UsuarioID).First(&usuario).Error

}

func (r repository) FetchPedidos() (pedido []models.Pedido, err error) {
	return pedido, r.conn.Find(&pedido).Error
}

func (r repository) Add(pedido models.Pedido) (models.Pedido, error) {
	return pedido, r.conn.Create(&pedido).Error
}

func (r repository) Update(ID uint64, name string, quantidade uint) error {
	return r.conn.Model(&models.Pedido{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"name":       name,
		"quantidade": quantidade}).Error
}

func (r repository) Remove(ID uint64) error {
	return r.conn.Delete(&models.Pedido{}, ID).Error
}
