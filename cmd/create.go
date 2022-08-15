/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"folderGen/internals/folder"

	"github.com/spf13/cobra"
)

var creationFilePath string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create directory from creation File",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.ValidateArgs(args)
		if err != nil {
			fmt.Println(err)
		}
		folder.CreateDirectoryOfCreationFile(creationFilePath)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&creationFilePath, "file", "f", "", "Path of creation file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
