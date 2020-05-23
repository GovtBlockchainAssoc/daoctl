package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var yamlDefault = []byte(`
EosioEndpoint: https://api.telos.kitchen
AssetsAsFloat: true
DAOContract: dao.hypha
Treasury:
  TokenContract: token.hypha
  Symbol: HUSD
  Contract: bank.hypha
  EthUSDTContract: 0xdac17f958d2ee523a2206206994597c13d831ec7
  EthUSDTAddress: 0xC20f453a4B4995CA032570f212988F4978B35dDd
  BtcAddress: 35hfgfaUouzYixTUDV6CFqMiTSZcuNtTf9
TelosDecideContract: trailservice
`)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "daoctl",
	Short: "Decentralized Autonomous Organization (DAO) control application for Hypha DAO query and actions",
	Long: `Decentralized Autonomous Organization (DAO) control application for Hypha DAO query and actions.
Query and manage roles, assignments, periods, payouts, and proposals.

Example use:
	daoctl get assignments --include-proposals
	daoctl get treasury

Hypha - Dapps for a New World - visit online @ hypha.earth`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./daoctl.yaml)")
	// RootCmd.Flags().BoolP("assets-as-floats", "f", false, "Format assets objects as floats (helpful for CSV export)")
	//RootCmd.Flags().BoolP("include-proposals", "p", false, "Include proposals when retrieving objects")
	RootCmd.PersistentFlags().StringP("vault-file", "", "./eosc-vault.json", "Wallet file that contains encrypted key material")
	RootCmd.PersistentFlags().IntP("delay-sec", "", 0, "Set time to wait before transaction is executed, in seconds. Defaults to 0 second.")
	RootCmd.PersistentFlags().IntP("expiration", "", 30, "Set time before transaction expires, in seconds. Defaults to 30 seconds.")
  RootCmd.PersistentFlags().BoolP("include-archive", "o", false, "include a table with the archive objects")
  RootCmd.PersistentFlags().BoolP("include-proposals", "i", false, "include a table with proposals in the output")
  RootCmd.PersistentFlags().BoolP("active", "a", true, "show active objects")
  RootCmd.PersistentFlags().BoolP("failed-proposals", "", false, "include a table with failed proposals")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

		// Search config in home directory with name ".daoctl" (without extension).
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(home)
		viper.SetConfigName("daoctl")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		viper.ReadConfig(bytes.NewBuffer(yamlDefault))
	}

	viper.SetEnvPrefix("daoctl")
	viper.AutomaticEnv() // read in environment variables that match
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	recurseViperCommands(RootCmd, nil)

	if viper.GetBool("global-debug") {
		zlog, err := zap.NewDevelopment()
		if err == nil {
			SetLogger(zlog)
		}
	}
}

func recurseViperCommands(root *cobra.Command, segments []string) {
	// Stolen from: github.com/abourget/viperbind
	var segmentPrefix string
	if len(segments) > 0 {
		segmentPrefix = strings.Join(segments, "-") + "-"
	}

	root.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		newVar := segmentPrefix + "global-" + f.Name
		viper.BindPFlag(newVar, f)
	})
	root.Flags().VisitAll(func(f *pflag.Flag) {
		newVar := segmentPrefix + "cmd-" + f.Name
		viper.BindPFlag(newVar, f)
	})

	for _, cmd := range root.Commands() {
		recurseViperCommands(cmd, append(segments, cmd.Name()))
	}
}
