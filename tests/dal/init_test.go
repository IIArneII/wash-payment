package dal

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"testing"
	"wash-payment/internal/app"
	"wash-payment/internal/dal"
	"wash-payment/internal/pkg/db"
	"wash-payment/internal/pkg/logger"

	"github.com/gocraft/dbr/v2"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.uber.org/zap"
)

const (
	user          = "postgres"
	password      = "postgres"
	dbname        = "wash_payment"
	migrationsDir = "../../migrations"
)

var (
	repositories *app.Repositories
	ctx          = context.Background()
	service      *app.Services
)

func TestMain(m *testing.M) {
	l, err := logger.NewLogger(logger.LevelInfo)
	if err != nil {
		log.Fatal(err)
	}

	closeContainer, port, err := createContainer(l)
	if err != nil {
		l.Fatal(err)
	}
	defer func() {
		closeContainer()
		l.Info("close container")
	}()

	dbConn, err := initDB(port)
	if err != nil {
		l.Fatal(err)
	}
	defer func() {
		dbConn.Close()
		l.Info("close db connection")
	}()

	repositories = dal.NewRepositories(l, dbConn)

	m.Run()
}

func createContainer(l *zap.SugaredLogger) (func(), int, error) {
	port, err := getFreePort()
	if err != nil {
		return nil, 0, nil
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, 0, nil
	}

	opts := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
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
		return nil, 0, nil
	}

	return func() {
		if err := pool.Purge(resource); err != nil {
			l.Fatal(err)
		}
	}, port, nil
}

func initDB(port int) (*dbr.Connection, error) {
	dbConn, err := db.InitDB(dns(user, password, dbname, "localhost", port))
	if err != nil {
		return nil, err
	}

	err = db.UpMigrations(dbConn.DB, migrationsDir)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
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
