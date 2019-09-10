package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init (){
  rootCmd.AddCommand(addDocument)
}
//Defining new command
var addDocument = &cobra.Command{
  Use: "adddocument",
  Short: "Add new document to a given index",
  Long:  `Use this command to add a new document to the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}