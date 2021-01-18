package cassandra

import (
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

const (
	cassandraHostURLsVar = "CASSANDRA_HOSTS"
)

var (
	session        *gocql.Session
	cassandraHosts = goDotEnvVariable(cassandraHostURLsVar)
)

func init() {
	// Connect to Cassandra cluster
	cluster := gocql.NewCluster(cassandraHosts)
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

// GetSession returns a new cassandra session and/or an error
func GetSession() *gocql.Session {
	return session
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
