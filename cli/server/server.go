package main

import (
  "errors"
  "io/ioutil"
  "log"
  "os"

  "github.com/wearp/news/service"
  "github.com/codegangsta/cli"
  "gopkg.in/yaml.v1"
)

func getConfig(c *cli.Context) (service.Config, error) {
  yamlPath := c.GlobalString("config")
  config := service.Config{}

  if _, err := os.Stat(yamlPath); err != nil {
    return config, errors.New("config path not valid")
  }

  ymlData, err := ioutil.ReadFile(yamlPath)
  if err != nil {
    return config, err
  }

  err = yaml.Unmarshal([]byte(ymlData), &config)
  return config, err
}

func main() {

  app := cli.NewApp()
  app.Name = "NEWS"
  app.Usage = "National Early Warning Score (NEWS) Microservice"
  app.Version = "0.0.1"

  app.Flags = []cli.Flag{
    cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
  }

  app.Commands = []cli.Command{
    {
      Name: "server",
      Usage: "Run the http server",
      Action: func(c *cli.Context) {
              cfg, err := getConfig(c)
              if err != nil {
                log.Fatal(err)
                return
              }
              
              svc := service.NewService{}

              if err = svc.Run(cfg); err != nil {
                log.Fatal(err)
              }
      },
    },
    { 
      Name: "migratedb",
      Usage: "Perform database migration",
      Action: func(c *cli.Context) {
              cfg, err := getConfig(c)
              if err != nil {
                log.Fatal(err)
                return
              }

              svc := service.NewService{}

              if err = svc.Migrate(cfg); err != nil {
                log.Fatal(err)
              }
      },
    },
  }
  app.Run(os.Args)
}
