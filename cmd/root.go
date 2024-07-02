package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/sincerefly/capybara/base"
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/cmd/cmdutils"
	"github.com/sincerefly/capybara/global"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
	"gitlab.com/avarf/getenvs"
	"strings"
)

type cobraFunc func(cmd *cobra.Command, args []string)
type pythonFunc func(cmd *cobra.Command, args []string, data Data)

type Data struct {
	Name string
}

func python(fn pythonFunc, envName string) cobraFunc {
	return func(cmd *cobra.Command, args []string) {
		name := getenvs.GetEnvString(envName, "dong")
		data := Data{
			Name: name,
		}
		fn(cmd, args, data)
	}
}

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	flags := rootCmd.Flags()
	flags.BoolP("version", "v", false, "print version")

	rootCmd.PersistentFlags().BoolVar(&global.ParamDebug, "debug", false, "enables detailed logging for debugging.")
	rootCmd.PersistentFlags().BoolVar(&global.ParamNoParallelism, "no-parallelism", false, "disables parallel processing, without goroutine.")
}

var rootCmd = &cobra.Command{

	Use:   "border",
	Short: "Add Borders to Images in Bulk",
	Long:  "A command-line tool to add simple borders and annotate photo information such as camera model, manufacturer, focal length, lens, etc.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) { // all subcommand
		debug, _ := cmd.Flags().GetBool("debug")
		if debug {
			log.SetLoggerDebugLevel() // show debug log
		}
	},
	Run: python(func(cmd *cobra.Command, args []string, data Data) {

		if cmd.Flags().Lookup("version").Changed {
			base.PrintVersion()
			return
		}

		log.Info(cfgFile)

	}, "NAME"),
}

func initConfig() {
	if cfgFile == "" {
		home, err := homedir.Dir()

		cmdutils.CheckErr(err)
		v.AddConfigPath(".")
		v.AddConfigPath(home)
		v.AddConfigPath("/etc/border/")
		v.SetConfigName("config")
		v.SetConfigType("toml")
	} else {
		v.SetConfigFile(cfgFile)
	}

	v.SetEnvPrefix("border")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigParseError); ok {
			panic(err)
		}
		cfgFile = "No config file used"
	} else {
		cfgFile = "Using config file: " + v.ConfigFileUsed()
	}
}
