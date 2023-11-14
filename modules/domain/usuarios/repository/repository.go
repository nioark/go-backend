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

func (r repository) Get(ID int64) (usuario models.Usuario, err error) {
	return usuario, r.conn.First(&usuario).Error
}

func (r repository) Fetch() (usuarios []models.Usuario, err error) {
	return usuarios, r.conn.Find(&usuarios).Error
}

func (r repository) Add(usuario models.Usuario) (usuariof models.Usuario, err error) {
	return models.Usuario{}, r.conn.Create(&usuario).Error
}

func (r repository) Update(ID int64, name, password string) error {
	return r.conn.Model(&models.Usuario{}).Where("id = ?", ID).Updates(map[string]interface{}{
		"name":     name,
		"password": password}).Error
}

func (r repository) Remove(ID int64) error {
	return r.conn.Delete(&models.Usuario{}, ID).Error
}
