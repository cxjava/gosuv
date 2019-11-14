//go:generate go run github.com/UnnoTed/fileb0x b0x.yaml

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cxjava/gosuv/log"
	"github.com/urfave/cli"
)

var (
	cfg Configuration
)

func checkServerStatus() error {
	resp, err := http.Get(cfg.Client.ServerURL + "/api/status")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var ret JSONResponse
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return errors.New("json loads error: " + string(body))
	}
	if ret.Status != 0 {
		return fmt.Errorf("%v", ret.Value)
	}
	return nil
}

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		showVersion()
	}
}

func main() {
	var defaultConfigPath = filepath.Join(defaultGosuvDir, "conf/config.yml")

	app := cli.NewApp()
	app.Name = "gosuv"
	app.Usage = "golang port of python-supervisor"
	app.Before = func(c *cli.Context) error {
		var err error
		cfgPath := c.String("conf")
		cfg, err = readConf(cfgPath)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "cxjava",
			Email: "cxjava@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "conf, c",
			Usage: "config file",
			Value: defaultConfigPath,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "start-server",
			Usage: "Start supervisor and run in background",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "foreground, f",
					Usage: "start in foreground",
				},
				&cli.StringFlag{
					Name:  "conf, c",
					Usage: "config file",
					Value: defaultConfigPath,
				},
			},
			Action: actionStartServer,
		},
		{
			Name:    "status",
			Aliases: []string{"st"},
			Usage:   "Show program status",
			Action:  actionStatus,
		},
		{
			Name:   "start",
			Usage:  "Start program",
			Action: actionStart,
		},
		{
			Name:   "stop",
			Usage:  "Stop program",
			Action: actionStop,
		},
		{
			Name:   "reload",
			Usage:  "Reload config file",
			Action: actionReload,
		},
		{
			Name:  "shutdown",
			Usage: "Shutdown server",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "restart, r",
					Usage: "restart server(todo)",
				},
			},
			Action: actionShutdown,
		},
		{
			Name:    "conftest",
			Aliases: []string{"t"},
			Usage:   "Test if config file is valid",
			Action:  actionConfigTest,
		},
		{
			Name:   "edit",
			Usage:  "Edit config file",
			Action: actionEdit,
		},
	}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
