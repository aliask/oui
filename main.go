package main

import (
	"fmt"
	"os"

	"github.com/aliask/oui/ouidb"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "oui",
		Short: "OIU lookup tool",
	}

	var lookupCmd = &cobra.Command{
		Use:   "lookup [MAC Address]",
		Short: "Perform lookup",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				ouidb.Lookup(args[0])
			} else {
				fmt.Println("Please provide a MAC address or use the 'update' command.")
			}
		},
	}
	rootCmd.AddCommand(lookupCmd)

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update the OUI database from IEEE",
		Run: func(cmd *cobra.Command, args []string) {
			ouidb.UpdateDatabase()
		},
	}
	rootCmd.AddCommand(updateCmd)

	// default to performing lookup if no valid command is given
	_, _, err := rootCmd.Find(os.Args[1:])
	if err != nil {
		args := append([]string{"lookup"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
