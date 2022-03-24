package data

type DbClientTest struct {
}

func NewDbClientTest() *DbClientTest {

	return &DbClientTest{}
}

func (p *DbClientTest) Add(query interface{}, args ...interface{}) (err error) {

	return nil
}

func (p *DbClientTest) GetAll() (response interface{}, err error) {

	return response, nil
}
