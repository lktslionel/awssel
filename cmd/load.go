package cmd

import (
	"fmt"

	"github.com/lktslionel/awssel/env"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load env vars for a service",
	Long: `Load all environment variables stored in 
AWS SSM Parameter Store for a given service`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		fmt.Println(args)
		fmt.Println("serviceName", serviceName)
		fmt.Println("prefixPath", prefixPath)
		fmt.Println("awsRegion", awsRegion)
		fmt.Println("filterPattern", filterPattern)
		fmt.Println("outputFormat", outputFormat)
		fmt.Println("exportable", exportable)
		fmt.Println("awsRole", awsRole)
		fmt.Println("awsProfile", awsProfile)

		return
	},
}

// Flags variables
var (
	serviceName   string
	prefixPath    string
	awsRegion     string
	filterPattern string
	outputFormat  string
	exportable    bool
	awsRole       string
	awsProfile    string
)

func init() {
	rootCmd.AddCommand(loadCmd)

	// Flags
	//loadCmd.Flags().BoolP("help", "h", true, "Show help")
	loadCmd.Flags().StringVar(&serviceName, "service-name", "", "service name (required)")
	loadCmd.Flags().StringVar(&prefixPath, "prefix-path", "", "Prefix output string with 'export' statement (required)")
	loadCmd.Flags().StringVar(&awsRegion, "aws-region", "", "AWS region name (required)")
	loadCmd.Flags().StringVar(&filterPattern, "filter-pattern", "", "regex that every env var name must match")
	loadCmd.Flags().StringVar(&outputFormat, "output-format", env.DefaultFormat, "env vars's output format")
	loadCmd.Flags().BoolVar(&exportable, "export", false, "Prefix output string with export statement")

	loadCmd.MarkFlagRequired("service-name")
	loadCmd.MarkFlagRequired("prefix-path")
	loadCmd.MarkFlagRequired("aws-region")

	loadCmd.Flags().SortFlags = false
	//loadCmd.Flags().StringVar(&awsRole, 			"awsRole"
	//loadCmd.Flags().StringVar(&awsProfile, 		"awsProfile"
}
