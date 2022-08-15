/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"folderGen/internals/folder"

	"github.com/spf13/cobra"
)

var (
	pathDirectory  string
	nameTemplate   string
	createLocation string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get template of a directory",
	Run: func(cmd *cobra.Command, args []string) {
		folder.CreateTemplateOfDirectory(pathDirectory, nameTemplate, createLocation)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&pathDirectory, "directory", "d", "", "directory for scan")
	getCmd.Flags().StringVarP(&nameTemplate, "name", "o", "template", "name out json")
	getCmd.Flags().StringVarP(&createLocation, "location", "l", "", "location for template")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
