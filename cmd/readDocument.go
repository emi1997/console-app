package cmd

import (
	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(readDocument)
}

//Defining new command
var readDocument = &cobra.Command{
	Use:   "readdoc",
	Short: "read document from a given index",
	Long:  `Use this command to read a document from the index of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.ReadDocument()
	},
}
