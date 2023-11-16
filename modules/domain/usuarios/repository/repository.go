package repository

import (
	"main/models"
	"main/modules/domain/usuarios"

	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) usuarios.Repository {
	return &repository{conn: conn}
}

func (r repository) Get(ID uint64) (usuario models.Usuario, err error) {
	return usuario, r.conn.First(&usuario, ID).Error
}

func (r repository) FetchUsuarios() (usuarios []models.Usuario, err error) {
	return usuarios, r.conn.Find(&usuarios).Error
}

func (r repository) FetchPedidos(ID uint64) (pedidos []models.Pedido, err error) {
	return pedidos, r.conn.Where("usuario_id = ?", ID).Find(&pedidos).Error
}

func (r repository) Add(usuario models.Usuario) (models.Usuario, error) {
	return usuario, r.conn.Create(&usuario).Error
}

func (r repository) Update(ID uint64, name, password string) error {
	return r.conn.Model(&models.Usuario{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"name":     name,
		"password": password}).Error
}

func (r repository) Remove(ID uint64) error {
	return r.conn.Delete(&models.Usuario{}, ID).Error
}
