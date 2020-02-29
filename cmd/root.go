package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "orthanc-drop",
		Short: "orthanc-drop",
		Long:  `orthanc-drop`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("Error: %v, ", err)
		os.Exit(1)
	}
}
