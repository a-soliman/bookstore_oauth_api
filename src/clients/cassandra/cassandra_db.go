package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to Cassandra cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("cassandra connection successfully created")
	// defer session.Close()
}

// GetSession returns a new cassandra session and/or an error
func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
