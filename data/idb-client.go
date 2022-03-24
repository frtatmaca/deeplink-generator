package data

type IDbClient interface {
	GetAll() (response interface{}, err error)
	//GetById(int32) interface{}
	Add(interface{}, ...interface{}) (err error)
	//Update(interface{}) bool
	//Delete(int) bool
	//Find(string) interface{}
}
