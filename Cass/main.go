package Cass

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session
var err error

func init() {

	cluster := gocql.NewCluster("127.0.0.1")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter the Keyspace name or press enter for default:")
	scanner.Scan()
	if scanner.Text() == "" {
		cluster.Keyspace = "usersmessages"
	} else {
		cluster.Keyspace = scanner.Text()
		scan := scanner.Text()
		if scan != "usersmessages" {
			fmt.Println("Entered keyspace: ", scan, "does not exist")
			// for now like this
			fmt.Println("Relaunch the program and enter proper Keyspace or create it before init")
			os.Exit(0)
		}
	}

	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Cluster", cluster.Keyspace, "doesn't exist")
		panic(err)
	}
	fmt.Println(LogTime() + ", Cassandra has been initialized")
	fmt.Println(LogTime()+", Cluster initialized:", cluster.Keyspace)
}

func LogTime() string {
	msTime := strconv.FormatUint(uint64(time.Now().Unix())*1000, 10)
	return msTime
}

func ChkErrs(err error) {
	if err != nil {
		panic(err)
	}
}
