package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/kydance/ziwi/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigDir  = "etc"
	defaultConfigName = "ziwi"

	envPrefix = "ZIWI"
)

var cfg string

func init() {
	flag.StringVar(&cfg, "config", "",
		"The path to the ziwi configuration file. Empty string for no configuration file.")
	flag.Parse()

	if cfg != "" {
		// Read config file from cfgFile.
		viper.SetConfigFile(cfg)
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
	log.NewLogger(&log.Options{
		Prefix:    viper.GetString("log.prefix"),
		Directory: viper.GetString("log.directory"),

		TimeLayout: viper.GetString("log.time-layout"),
		Level:      viper.GetString("log.level"),
		Format:     viper.GetString("log.format"),

		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		DisableSplitError: viper.GetBool("log.disable-split-error"),

		MaxSize:    viper.GetInt("log.max-size"),
		MaxBackups: viper.GetInt("log.max-backups"),
		Compress:   viper.GetBool("log.compress"),
	})

	log.Info("Test ziwi log")
}
