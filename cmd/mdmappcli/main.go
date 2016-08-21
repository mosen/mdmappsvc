package main

import (
	"fmt"
	"os"
	"github.com/containous/flaeg"
	_ "github.com/BurntSushi/toml"
	"github.com/containous/staert"
	_ "github.com/lib/pq"
	"github.com/go-kit/kit/log"
	"github.com/mosen/mdmappsvc/client"
	"golang.org/x/net/context"
)




type Configuration struct {
	URL string `description:"URL of the mdmappsvc"`
}

type SourcesConfiguration struct {
	URL string `description:"URL of the mdmappsvc"`
	List bool `description:"List sources"`
}

func main() {
	var config *Configuration = &Configuration{
		URL: "http://localhost:8080/",
	}

	var pointersConfig *Configuration = &Configuration{}

	rootCmd := &flaeg.Command{
		Name: "mdmappcli",
		Description: "MDM app service client",
		Config: config,
		DefaultPointersConfig: pointersConfig,
		Run: func() error {
			run(config)
			return nil
		},
	}

	var srcConfig *SourcesConfiguration = &SourcesConfiguration{
		URL: "http://localhost:8080/",
		List: false,
	}
	var srcPointersConfig *SourcesConfiguration = &SourcesConfiguration{}
	sourcesCmd := &flaeg.Command{
		Name: "sources",
		Description: "Read and write application sources",
		Config: srcConfig,
		DefaultPointersConfig: srcPointersConfig,
		Run: func() error {
			runSources(srcConfig)
			return nil
		},
	}

	fl := flaeg.New(rootCmd, os.Args[1:])
	fl.AddCommand(sourcesCmd)

	if _, err := fl.Parse(rootCmd); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	st := staert.NewStaert(rootCmd)
	toml := staert.NewTomlSource("mdmappcli", []string{"."})

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
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	_, err := client.New(config.URL, logger)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}


}

func runSources(config *SourcesConfiguration) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	}

	logger.Log("level", "debug", "msg", "Fetching sources")

	svc, err := client.New(config.URL, logger)
	if err != nil {
		logger.Log("level", "error", "msg", "Cannot create client for URL")
	}

	sources, err := svc.GetSources(context.Background())
	if err != nil {
		logger.Log("level", "error", "msg", err)
		os.Exit(-1)
	}

	for _, source := range sources {
		fmt.Println(source)
	}
}
