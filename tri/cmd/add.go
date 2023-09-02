/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
        "fmt"
//
	"github.com/spf13/cobra"
        "github.com/tchioubak/learngo/tri/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Long: `Add will add a new todo item to the list`,
        /*
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
        */
        Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
    //items []todo.Item{}
    items, err := todo.ReadItems(dataFile)
    if err != nil {
        fmt.Errorf("%v", err)
    }
    for _, x := range args {
        items = append(items, todo.Item{x})
    }
//  fmt.Println("add : ", items)
//  fmt.Printf("add : %#v\n", items)
    err = todo.SaveItems(dataFile, items)
    if err != nil {
        fmt.Errorf("%v", err)
    }
}
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
