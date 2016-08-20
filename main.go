package main

import (
	"fmt"
	"os"
	"github.com/containous/flaeg"
	_ "github.com/BurntSushi/toml"
	"github.com/containous/staert"
)



type DatabaseInfo struct {
	Host string `description:"Hostname or IP address of postgresql server"`
	Port string `description:"database port number"`
	Name string `description:"database name"`
	Username string `description:"database username"`
	Password string `description:"database password"`
}

type ListenInfo struct {
	IP string `description:"IP Address to listen on"`
	Port string `description:"listen on port number"`
}

type Configuration struct {
	Db *DatabaseInfo `description:"Database"`
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
	fmt.Printf("%#v\n", config.Db)
}
