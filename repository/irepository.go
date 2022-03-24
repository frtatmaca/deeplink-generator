package repository

type IRepository interface {
	GetAll() interface{}
	GetById(int32) interface{}
	Add(interface{}) (err error)
	Update(interface{}) bool
	Delete(int) bool
	Find(string) interface{}
}
