package main

import "github.com/urfave/cli/v2"

func main() {
	cfg := struct {
		taskCoordinatorURL string
	}{}
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "task-coordinator-url",
				Usage:       "URL of the task coordinator",
				EnvVars:     []string{"TASK_COORDINATOR_URL"},
				Destination: &cfg.taskCoordinatorURL,
				Required:    true,
			},
		},
	}
	app.RunAndExitOnError()
}
