package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type FlagConfig struct {
	Name       string
	Alias      string
	Default    interface{}
	Usage      string
	EnvVar     string
	Persistent bool
}

type customValue struct {
	value interface{}
}

func (c *customValue) String() string {
	return fmt.Sprintf("%v", c.value)
}

func (c *customValue) Set(value string) error {
	c.value = value
	return nil
}

func (c *customValue) Type() string {
	return fmt.Sprintf("%T", c.value)
}

func BindFlag(cmd *cobra.Command, config FlagConfig) {
	value := &customValue{value: config.Default}
	viper.RegisterAlias(config.Name, config.Name)

	flag := &pflag.Flag{
		Name:      config.Name,
		Shorthand: config.Alias,
		Usage:     config.Usage,
		Value:     value,
	}

	cmd.Flags().VarP(value, config.Name, config.Alias, config.Usage)

	if config.EnvVar != "" {
		_ = viper.BindEnv(config.Name, config.EnvVar)
	}

	if config.Persistent {
		cmd.PersistentFlags().AddFlag(flag)
	} else {
		cmd.Flags().AddFlag(flag)
	}
}
