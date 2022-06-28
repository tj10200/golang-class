package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tj10200/golang-class/class_1/00_base_lesson_setup/pkg/bluecore"
	"os"
	"strings"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

// EnvPrefix All flags will be available in the OS environment as "LESSON_<flag_name>"
const EnvPrefix = "LESSON"

// Our two dummy env vars
var hello string
var world string

func NewRootCommand() *cobra.Command {
	// Define our command
	rootCmd := &cobra.Command{
		Use:   "00_base_lesson_setup",
		Short: "boiler plate for golang lessons",
		Long:  `boiler plate for golang lessons`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateConfig(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunCommand(bluecore.Config{
				Hello: hello,
				World: world,
			})
		},
	}

	setupRootFlags(rootCmd)

	return rootCmd
}

func setupRootFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("config_file", "example_toml_config", "the configuration file, no extension")
	cmd.PersistentFlags().StringSlice("config_search_path", []string{"./config"}, "the list of paths to search for the config file")

	cmd.Flags().StringVar(&hello, "hello", "hello", "the hello part of hello world")
	cmd.Flags().StringVar(&world, "world", "world", "the world part of hello world")
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	cfgFile, err := cmd.Flags().GetString("config_file")
	if err != nil {
		return err
	}

	if f, ok := os.LookupEnv(fmt.Sprintf("%s_CONFIG_FILE", EnvPrefix)); ok {
		cfgFile = f
	}
	// Set the base name of the config file, without the file extension.
	v.SetConfigName(cfgFile)

	// Set the search path from the command line args
	cfgSearchPath, err := cmd.Flags().GetStringSlice("config_search_path")
	if err != nil {
		return err
	}

	// Override the search path if set in the environment
	if p, ok := os.LookupEnv(fmt.Sprintf("%s_CONFIG_SEARCH_PATH", EnvPrefix)); ok {
		cfgSearchPath = strings.FieldsFunc(p, func(r rune) bool {
			return r == ',' || r == ' '
		})
	}
	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	for _, path := range cfgSearchPath {
		v.AddConfigPath(path)
	}

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable RBAC_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(EnvPrefix)

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	v.BindPFlags(cmd.PersistentFlags())
	v.BindPFlags(cmd.Flags())

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --pg-host to RBAC_PG_HOST
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", EnvPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func validateConfig(cmd *cobra.Command) error {
	if len(hello) == 0 {
		return fmt.Errorf("hello var not set")
	} else if len(hello) > 255 {
		return fmt.Errorf("hello var too long")
	}

	if len(world) == 0 {
		return fmt.Errorf("world var not set")
	} else if len(world) > 255 {
		return fmt.Errorf("world var too long")
	}

	// Check for valid flag configurations here
	return nil
}
