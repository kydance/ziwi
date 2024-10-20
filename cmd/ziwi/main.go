package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kydance/ziwi/pkg/log"
	"github.com/kydance/ziwi/pkg/strutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigDir  = "etc"
	defaultConfigName = "ziwi"

	envPrefix = "ZIWI"
)

var cfgFile string

// run is the real main entry point.
func run() error {
	settings, err := json.Marshal(viper.AllSettings())
	if err != nil {
		log.Errorw("Failed to marshal viper settings", "err", err)
		return err
	}

	log.Infow(string(settings))

	log.Debugf("strutil.CamelCase(\"hello world\"): %v\n", strutil.CamelCase("hello world"))

	str := "*kyden*"
	log.Debugf("strutil.UnWarp(str, \"*\"): %v\n", strutil.UnWarp(str, "*"))
	log.Infoln(str)

	return nil
}

// NewZiwiCommand creates *cobra.Command object. Then, call Execute to run application.
func NewZiwiCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ziwi", // Command name
		Short: "A Go tools lib",
		Long: `A Go tools lib.

Find more ziwi information at:
	https://github.com/kydance/ziwi#readme`,
		// Commands that fail to print the usage.
		SilenceUsage: false,

		// When running cmd.Execute(), it will be called.
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init log
			log.Init(&log.Options{
				DisableCaller:     viper.GetBool("log.disable-caller"),
				DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
				Level:             viper.GetString("log.level"),
				Format:            viper.GetString("log.format"),
				OutputPaths:       viper.GetStringSlice("log.output-paths"),
			})

			// Sync write buffer log to file
			defer log.Sync()

			return run()
		},

		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q",
						cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	// Other command flags
	// ...

	// 持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c",
		"", "The path to the ziwi configuration file. Empty string for no configuration file.")
	// 本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

func initConfig() {
	if cfgFile != "" {
		// Read config file from cfgFile.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(home, defaultConfigDir)) // $HOME/defaultConfigDir
		viper.AddConfigPath(filepath.Join(".", defaultConfigDir))  // ./defaultConfigDir

		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}

	// Read matched environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read config file.
	// If a config file is specified, use it. Otherwise, search in defaultConfigDir.
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

func main() {
	cmd := NewZiwiCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
