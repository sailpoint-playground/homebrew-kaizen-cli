/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kaizen",
	Short: "A powerful CLI for streamlined CI/CD operations.",
	Long: `Kaizen CLI is your comprehensive command-line tool designed to automate and
simplify various Continuous Integration and Continuous Delivery tasks.

With Kaizen, you can effortlessly manage your CI/CD pipelines, including:

- **Spacelift Stack Management:** Create and configure Spacelift stacks directly.
- **Jira Integration:** Generate and update Jira tickets for workflow automation.
- **ArgoCD Application Deployments:** Deploy and manage your applications within ArgoCD.
- **Image Tag Updates:** Automate pull requests for image tag changes in your repositories.
- **Multi-Architecture Docker Builds:** Build and publish multi-architecture Docker images.
- **ECR Publishing:** Push your Docker images to Amazon Elastic Container Registry (ECR).

Kaizen aims to enhance your development workflow by providing a single, intuitive
interface for common CI/CD actions, making your processes more efficient and reliable.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kaizen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

