package cmd

import (
	"github.com/j4ng5y/go-ots/pkg/srv"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

func runFunc(cmd *cobra.Command, args []string) {
	cfgFile, err := cmd.Flags().GetString("config-file")
	if err != nil {
		klog.Fatal(err)
	}

	cfg := newCLI()
	if err := cfg.init(cfgFile); err != nil {
		klog.Fatal(err)
	}

	S, err := srv.New(cfg.cfg)
	if err != nil {
		klog.Fatal(err)
	}

	S.Run()
}
