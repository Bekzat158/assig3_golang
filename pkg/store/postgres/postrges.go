// postgres.go
package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Store struct {
	Pool *pgxpool.Pool
}

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}

func (s Settings) toDSN() string {
	var args []string

	if s.Host != "" {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	} else {
		args = append(args, "host=localhost")
	}

	if s.Port != 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	} else {
		args = append(args, "port=5432")
	}

	if s.Database != "" {
		args = append(args, fmt.Sprintf("dbname=%s", s.Database))
	} else {
		args = append(args, "dbname=postgres")
	}

	if s.User != "" {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	} else {
		args = append(args, "user=postgres")
	}

	if s.Password != "" {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	} else {
		args = append(args, "password=password")
	}

	if s.SSLMode != "" {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	} else {
		args = append(args, "sslmode=disable")
	}

	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {
	config, err := pgxpool.ParseConfig(settings.toDSN())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{Pool: conn}, nil
}
