/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Gabriel-Macedogmc/hexagonal-architecture/adapters/cli"
	"github.com/spf13/cobra"
)

var (
	action       string
	productId    string
	productName  string
	productPrice float64
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(service, action, productId, productName, productPrice)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	// Here you will define your flags and configuration settings.
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable / Disable")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
