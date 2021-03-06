package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/datakube/datakube/configuration"
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/server"
	"github.com/datakube/datakube/store"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string
var logLevel string
var config configuration.ServerConfiguration

// RootCmd represents the base command when called without any subcommands
var serverCommand = &cobra.Command{
	Use:   "datakube/",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.Unmarshal(&config)

		if err != nil {
			fmt.Printf("unable to decode into struct, %v", err)
		}

		level, err := logrus.ParseLevel(strings.ToLower(config.LogLevel))
		if err != nil {
			log.Error("Error getting level", err)
		}
		log.SetLevel(level)

		store := initStore(config.Datastore.Path)

		var stopChan = make(<-chan struct{})

		Server := server.NewServer(config, store)
		Server.Start(stopChan)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := serverCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initServerConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	serverCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.toml)")
	serverCommand.PersistentFlags().StringVar(&logLevel, "logLevel", "error", "Dumper Application Log Level")

	viper.BindPFlag("config", serverCommand.PersistentFlags().Lookup("config"))
	viper.BindPFlag("logLevel", serverCommand.PersistentFlags().Lookup("logLevel"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	serverCommand.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initServerConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".dodo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file %s", err.Error())
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func initStore(dataStorePath string) *store.DataStore {
	store, err := store.NewStore(dataStorePath)
	if err != nil {
		log.Fatal(err)
	}

	err = store.Open()

	if err != nil {
		log.Fatal(err)
		os.Exit(15)
	}

	return store
}
