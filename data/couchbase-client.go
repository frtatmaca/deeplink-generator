package data

import (
	"errors"
	"flag"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/deeplink/models"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

type CouchbaseClient struct {
}

func NewCouchbaseClient() *CouchbaseClient {
	connectCB()
	//createLinkTable(Db)
	return &CouchbaseClient{}
}

var cbDb *TravelSampleApp

var (
	defaultCBHost          = "localhost"
	defaultCBScheme        = "couchbase://" // Set to couchbase:// if using Couchbase Server Community Edition
	travelSampleBucketName = "hello-world"
	defaultCBUsername      = "Administrator"
	defaultCBPassword      = "123456"
	jwtSecret              = []byte("IAMSOSECRETIVE!")
)

type TravelSampleApp struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

func envFlagString(envName, name, value, usage string) *string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		value = envValue
	}
	return flag.String(name, value, usage)
}

func connectCB() {
	var err error
	connStr := envFlagString("CB_HOST", "host", defaultCBHost, "The connection string to use for connecting to the server")
	username := envFlagString("CB_USER", "user", defaultCBUsername, "The username to use for authentication against the server")
	password := envFlagString("CB_PASS", "password", defaultCBPassword, "The password to use for authentication against the server")
	scheme := envFlagString("CB_SCHEME", "scheme", defaultCBScheme,
		"The scheme to use for connecting to couchbase. Default to couchbases - set to couchbase:// when using couchbase community edition")

	flag.Parse()

	// Connect the SDK to Couchbase Server.
	clusterOpts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: *username,
			Password: *password,
		},
	}
	cluster, err := gocb.Connect(*scheme+*connStr, clusterOpts)
	if err != nil {
		panic(err)
	}

	// Create a bucket instance, which we'll need for access to scopes and collections.
	bucket := cluster.Bucket(travelSampleBucketName)

	cbDb = &TravelSampleApp{
		cluster: cluster,
		bucket:  bucket,
	}
}

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

func (p *CouchbaseClient) Add(query interface{}, args ...interface{}) (err error) {
	fmt.Println("frt frt add head")
	helloWorld := strings.ToLower("user-info")
	scope := cbDb.bucket.Scope(helloWorld)
	users := scope.Collection("users")

	userKey := strings.ToLower("frtatmaca2")

	user := models.JsonUser{
		Name:     "frt",
		Password: "12345678",
	}

	fmt.Println("frt frt add 1:")

	asd, _ := users.Upsert(userKey, user, nil)

	fmt.Println("frt frt add: %+v", asd)

	if errors.Is(err, gocb.ErrDocumentExists) {
		//writeJsonFailure(w, 409, ErrUserExists)
		return
	} else if err != nil {
		//app.logger.Printf("Error occurred inserting user: %s", err)
		//writeJsonFailure(w, 500, err)
		return
	}

	return nil
}

func (p *CouchbaseClient) delete() {
	fmt.Println("asd asd frt 3 delete")

	r, err := cbDb.cluster.Query("delete from `hello-world` as users where meta(users).id=`frtatmaca`", nil)

	//helloWorld := strings.ToLower("user-info")
	//scope := cbDb.bucket.Scope(helloWorld)
	//users := scope.Collection("users")

	//r, err := users.Remove("frtatmaca2", nil)

	fmt.Println("asd asd frt 3 delete err: %+v", err)
	fmt.Println("asd asd frt 3 delete r: %+v", r)

}

func (p *CouchbaseClient) GetAll() (response interface{}, err error) {
	p.delete()

	_ = p.Add(nil, nil)

	queryParams := make([]interface{}, 1)
	//fmt.Println("asd asd frt 1")
	queryStr := "SELECT name FROM `beer-sample`"
	//fmt.Println("asd asd frt 2: %+v", queryStr)
	fmt.Println("asd asd frt 3: %+v", queryParams)

	var respData models.JsonAirportSearchResp
	respData.Context.Add(fmt.Sprintf("N1QL query - scoped to inventory: %s", queryStr))

	//fmt.Println("asd asd frt 444: %+v", respData)

	/*rows, err := cbDb.cluster.Query(queryStr, nil)
	//fmt.Println(rows)

	if err != nil {
		fmt.Println(err)
		return
	}

	respData.Data = []models.JsonAirport{}
	for rows.Next() {
		var airport models.JsonAirport
		if err = rows.Row(&airport); err != nil {
			//app.logger.Printf("Error occurred during airport search result parsing: %s", err)
			//writeJsonFailure(w, 500, err)
			return
		}

		respData.Data = append(respData.Data, airport)
	}

	//fmt.Println("asd asd frt 3: %+v", respData)  */

	return respData, nil
}
