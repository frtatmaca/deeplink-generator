package repository

import (
	"github.com/deeplink/adapter"
	"github.com/deeplink/data"
)

type Repository struct {
	dbClient data.IDbClient
}

func NewRepository() Repository {
	repo := adapter.DbClient()

	return Repository{repo}
}

func (p *Repository) Add(query interface{}, args ...interface{}) (err error) {
	return p.dbClient.Add(query, args...)
}

func (p *Repository) GetAll() (interface{}, error) {

	return p.dbClient.GetAll()
}

func (p *Repository) GetById(id int32) interface{} {

	return ""
}

func (p *Repository) Update(interface{}) bool {

	return true
}

func (p *Repository) Delete(id int32) bool {

	return true
}

func (p *Repository) Find(key string) interface{} {

	return ""
}
