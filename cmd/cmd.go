package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

// Execute runs the CLI
func Execute() {
	var (
		cfgFile string
		rootCmd = &cobra.Command{
			Use:     "go-ots",
			Version: "0.1.0",
			Short:   "",
			Long:    "",
			Example: "",
			Run:     func(cmd *cobra.Command, args []string) {},
		}

		runCmd = &cobra.Command{
			Use:     "run",
			Short:   "Run the Go OTS Server",
			Long:    "",
			Example: "",
			Run:     runFunc,
		}
	)

	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&cfgFile, "config-file", "f", "", "the configuration file to use")
	runCmd.MarkFlagRequired("config-file")

	if err := rootCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
