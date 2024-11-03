package cli

import (
	"github.com/crypto-grill/app/internal/config"
	"github.com/crypto-grill/app/internal/config/viper"
	"github.com/crypto-grill/app/internal/runner"
	"github.com/crypto-grill/app/internal/server"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func Execute(args []string) error {
	cfg, err := viper.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	if err := configureLogger(cfg); err != nil {
		return err
	}

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "App CLI",
	}

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run HTTP server",
		RunE: func(_ *cobra.Command, _ []string) error {
			go func() {
				if err := runner.SyncMessages(); err != nil {
					log.Fatal(err)
				}
			}()

			return server.Start(cfg)
		},
	}

	rootCmd.AddCommand(serveCmd)
	rootCmd.SetArgs(args[1:])

	return rootCmd.Execute()
}

func configureLogger(cfg *config.Config) error {
	logConfig := zap.NewProductionConfig()
	level, err := zapcore.ParseLevel(cfg.Log.Level)
	if err != nil {
		return err
	}

	logConfig.Level.SetLevel(level)
	logConfig.DisableStacktrace = true
	logConfig.Encoding = "console"
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := logConfig.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}
