package main

import (
	"errors"
	"fmt"

	"github.com/josestg/dsa/x/dependency_injection/fx"
)

type Config struct {
	Env string
	DSN string
}

type Logger struct {
	c *Config
}

type Database struct {
	c *Config
	l *Logger
}

func NewConfig() *Config {
	return &Config{
		Env: "stg",
		DSN: "postgres://user:password@localhost:5432/my_db",
	}
}

func NewLogger(c *Config) *Logger {
	return &Logger{c: c}
}

func NewDatabase(c *Config, l *Logger) *Database {
	return &Database{c: c, l: l}
}

type Service struct {
	db  *Database
	log *Logger
}

func NewService(db *Database, log *Logger) *Service {
	return &Service{db: db, log: log}
}

func main() {
	c := fx.NewContainer()

	err := errors.Join(
		c.Provide(NewDatabase),
		c.Provide(NewConfig),
		c.Provide(NewLogger),
		c.Provide(NewService),
	)
	if err != nil {
		panic(err)
	}

	if err = c.Build(); err != nil {
		panic(err)
	}

	serv, err := fx.Resolve[*Service](c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", serv)
	fmt.Println(serv.db.c)
}
