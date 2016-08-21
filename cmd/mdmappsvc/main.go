package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/containous/flaeg"
	_ "github.com/BurntSushi/toml"
	"github.com/containous/staert"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/go-kit/kit/log"
	"github.com/DavidHuie/gomigrate"
	"github.com/mosen/mdmappsvc/source"
	"golang.org/x/net/context"
	"github.com/jmoiron/sqlx"
	"os/signal"
	"syscall"
)



type DatabaseInfo struct {
	Host string `description:"Hostname or IP address of postgresql server"`
	Port string `description:"database port number"`
	Name string `description:"database name"`
	Username string `description:"database username"`
	Password string `description:"database password"`
	SSLMode string `description:"postgres SSL mode"`
}

type ListenInfo struct {
	IP string `description:"IP Address to listen on"`
	Port string `description:"listen on port number"`
}

type Configuration struct {
	Db *DatabaseInfo `description:"Database connection options"`
	Listen *ListenInfo `description:"Listen"`
}

func main() {
	var config *Configuration = &Configuration{
		&DatabaseInfo{
			Host: "localhost",
			Port: "5432",
			Name: "mdmappsvc",
			Username: "mdmappsvc",
			Password: "mdmappsvc",
		},
		&ListenInfo{
			IP: "0.0.0.0",
			Port: "8080",
		},
	}

	var pointersConfig *Configuration = &Configuration{}

	rootCmd := &flaeg.Command{
		Name: "mdmappsvc",
		Description: "MDM app service stores information about apps that may be installed on an MDM client.",
		Config: config,
		DefaultPointersConfig: pointersConfig,
		Run: func() error {
			run(config)
			return nil
		},
	}

	fl := flaeg.New(rootCmd, os.Args[1:])
	if _, err := fl.Parse(rootCmd); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	st := staert.NewStaert(rootCmd)
	toml := staert.NewTomlSource("mdmappsvc", []string{"."})

	st.AddSource(toml)
	st.AddSource(fl)

	if _, err := st.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := st.Run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	os.Exit(0)
}

func run(config *Configuration) {
	var err error
	var db *sql.DB
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	}

	db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Db.Host,
		config.Db.Port,
		config.Db.Username,
		config.Db.Password,
		config.Db.Name,
	))
	if err != nil {
		logger.Log("level", "error", "msg", err)
		os.Exit(-1)
	}

	err = db.Ping()
	if err != nil {
		logger.Log("level", "error", "msg", err)
		os.Exit(-1)
	}

	var dbx *sqlx.DB = sqlx.NewDb(db, "postgres")

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./migrations")
	migrationErr := migrator.Migrate()

	if migrationErr != nil {
		logger.Log("level", "error", "msg", err)
		os.Exit(-1)
	}

	ctx := context.Background()

	sourceRepo := source.NewRepository(dbx, logger)
	sourceSvc := source.NewService(sourceRepo, logger)
	sourceSvc = source.LoggingMiddleware(log.NewContext(logger).With("component", "source.Service"))(sourceSvc)
	sourceHandler := source.MakeHTTPHandler(ctx, sourceSvc, log.NewContext(logger).With("component", "HTTP"))

	mux := http.NewServeMux()
	mux.Handle("/v1/sources/", sourceHandler)

	portStr := fmt.Sprintf("%v:%v", config.Listen.IP, config.Listen.Port)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("level", "info", "msg", "Listening on " + portStr)
		errs <- http.ListenAndServe(portStr, mux)
	}()

	logger.Log("exit", <-errs)
}
