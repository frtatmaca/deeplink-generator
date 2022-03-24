package adapter

import (
	"github.com/deeplink/config"
	"github.com/deeplink/data"
)

func DbClient() (repo data.IDbClient) {
	clientType := config.Config("DB_TYPE")
	switch clientType {
	case "MYSQL":
		repo = data.NewDbClient()
	case "COUCHBASE":
		repo = data.NewCouchbaseClient()
	default:
		repo = data.NewDbClientTest()
	}

	return
}
