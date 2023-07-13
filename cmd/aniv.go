/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mikuta0407/mltd-pat/calc"
	"github.com/mikuta0407/mltd-pat/config"
	"github.com/spf13/cobra"
)

// anivCmd represents the aniv command
var anivCmd = &cobra.Command{
	Use:   "aniv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aniv called")
		pointlist, err := config.ParseToml("pointlist.toml")
		if err != nil {
			fmt.Println(err)
			return
		}

		currentPoint, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		targetPoint, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		err = calc.Main(pointlist, "aniv", currentPoint, targetPoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(anivCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// anivCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// anivCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
