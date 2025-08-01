package cmd

import (
	"fmt"
	"os" // Needed for os.Hostname()

	"github.com/spf13/cobra"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Retrieves the hostname of the machine",
	Long: `This command fetches and displays the hostname of the system it is running on.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := os.Hostname()
		if err != nil {
			fmt.Printf("Error getting hostname: %v\n", err)
			return
		}
		fmt.Println(name)
	},
}

func init() {
	// Add the hostname command to the root command.
	// Assuming rootCmd is defined in cmd/root.go
	rootCmd.AddCommand(hostnameCmd)
}
