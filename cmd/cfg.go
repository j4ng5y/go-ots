package cmd

import "github.com/spf13/viper"

type cli struct {
	cfg *viper.Viper
}

func (c *cli) init(cfgFile string) error {
	c.cfg.SetConfigFile(cfgFile)

	if err := c.cfg.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func newCLI() *cli {
	var c cli
	c.cfg = viper.New()
	return &c
}
