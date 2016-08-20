package main

import (
	"fmt"
	"os"
	"github.com/containous/flaeg"
	"github.com/containous/staert"
)



type DatabaseInfo struct {
	Host string `description:"Hostname or IP address of postgresql server"`
	Port string `description:"Port number"`
}

type Configuration struct {
	Db DatabaseInfo `description:"Database"`
}

func main() {
	var config *Configuration = &Configuration{
		DatabaseInfo{
			Host: "localhost",
			Port: "5432",
		},
	}

	var pointersConfig *Configuration = &Configuration{}

	rootCmd := &flaeg.Command{
		Name: "mdmappsvc",
		Description: "MDM app service stores information about apps that may be installed on an MDM client.",
		Config: config,
		DefaultPointersConfig: pointersConfig,
		Run: func() error {
			fmt.Printf("Run flaegtest command with config : %+v\n", config)
			return nil
		},
	}

	st := staert.NewStaert(rootCmd)
	toml := staert.NewTomlSource("mdmappsvc", []string{"./"})
	fl := flaeg.New(rootCmd, os.Args[1:])

	st.AddSource(toml)
	st.AddSource(fl)
	loadedConfig, err := st.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	fmt.Printf("%#v\n", loadedConfig)
}
