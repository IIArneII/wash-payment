package dal

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/dal"
	"wash-payment/internal/pkg/db"
	"wash-payment/internal/pkg/logger"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/powerman/gotest/testinit"
	"go.uber.org/zap"
)

var ctx = context.Background()

var (
	user          = "postgres"
	password      = "postgres"
	dbname        = "test"
	migrationsDir = "../../internal/migrations"

	l            *zap.SugaredLogger
	repositories *app.Repositories
)

func TestMain(m *testing.M) { testinit.Main(m) }

func init() { initRepositories() }

func initRepositories() {
	port, err := getFreePort()
	if err != nil {
		testinit.Fatal(err)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		testinit.Fatal(err)
	}

	opts := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.7",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbname,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: strconv.Itoa(port)},
			},
		},
	}

	resource, err := pool.RunWithOptions(opts)
	if err != nil {
		testinit.Fatal(err)
	}

	testinit.Teardown(func() {
		if err := pool.Purge(resource); err != nil {
			testinit.Fatal(err)
		}
	})

	dbConn, err := db.InitDB(dns(user, password, dbname, "localhost", port))
	if err != nil {
		testinit.Fatal(err)
	}

	testinit.Teardown(func() {
		if err := dbConn.Close(); err != nil {
			testinit.Fatal(err)
		}
	})

	err = db.UpMigrations(dbConn.DB, migrationsDir)
	if err != nil {
		testinit.Fatal(err)
	}

	l, err := logger.NewLogger(logger.LevelInfo)
	if err != nil {
		testinit.Fatal(err)
	}

	repositories = dal.NewRepositories(l, dbConn)
}

func getFreePort() (port int, err error) {
	if a, err := net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		if l, err := net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}

func dns(user, password, dbname, host string, port int) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, password, dbname, host, port)
}
