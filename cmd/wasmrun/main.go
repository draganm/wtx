package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
	"github.com/urfave/cli/v2"
)

func main() {

	cfg := struct {
		wasmFile string
		env      *cli.StringSlice
	}{
		env: &cli.StringSlice{},
	}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "wasm-file",
				Aliases:     []string{"w"},
				Usage:       "path to the wasm file",
				EnvVars:     []string{"WASM_FILE"},
				Destination: &cfg.wasmFile,
				Required:    true,
			},
			&cli.StringSliceFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				EnvVars:     []string{"ENV"},
				Usage:       "environment variables",
				Destination: cfg.env,
			},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			runtime := wazero.NewRuntime(ctx)

			cl, err := wasi_snapshot_preview1.Instantiate(ctx, runtime)
			if err != nil {
				return fmt.Errorf("failed to instantiate WASI: %w", err)
			}

			defer cl.Close(ctx)

			wasmBytes, err := os.ReadFile(cfg.wasmFile)
			if err != nil {
				return fmt.Errorf("failed to read wasm file: %w", err)
			}

			module, err := runtime.CompileModule(ctx, wasmBytes)
			if err != nil {
				return fmt.Errorf("failed to compile WASM module: %w", err)
			}

			config := wazero.NewModuleConfig().
				WithStdin(os.Stdin).
				WithStdout(os.Stdout).
				WithStderr(os.Stderr).
				WithArgs(c.Args().Slice()...)

			for _, env := range cfg.env.Value() {
				k, v, ok := strings.Cut(env, "=")
				if !ok {
					return fmt.Errorf("invalid env variable: %s", env)
				}
				config = config.WithEnv(k, v)
			}

			startTime := time.Now()

			defer func() {
				fmt.Printf("elapsed time: %v\n", time.Since(startTime))
			}()

			instance, err := runtime.InstantiateModule(ctx, module, config)
			if err != nil {

				ee := &sys.ExitError{}

				if errors.As(err, &ee) {
					return cli.Exit("", int(ee.ExitCode()))
				}

				return fmt.Errorf("failed to instantiate module: %w", err)
			}

			defer instance.Close(context.Background())

			// err = instance.Close(ctx)
			// if err != nil {
			// 	return fmt.Errorf("failed to close instance: %w", err)
			// }

			// res, err := instance.ExportedFunction("_start").Call(ctx)

			// if err != nil {
			// 	return fmt.Errorf("failed to call _start: %w", err)
			// }

			// if len(res) == 0 {
			// 	return nil
			// }

			// if res[0] != 0 {
			// 	return cli.Exit(fmt.Sprintf("exit-code: %v", res), 1)
			// }

			return nil
		},
	}

	app.RunAndExitOnError()

}
