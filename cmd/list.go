/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: list,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func list(cmd *cobra.Command, args []string) {
	file, err := os.Open("output.csv")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("There are no tasks added!")
		return
	} else if err != nil {
		log.Fatal("Couldn't open the file", err)
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("Couldn't read the csv")
	}
	writer := tabwriter.NewWriter(
		os.Stdout, 0, 2, 4, ' ', 0,
	)

	writer.Write([]byte("ID\tDescription\tCreatedAt\tIsComplete\n"))
	for _, line := range lines {
		if line[0] == "ID" {
			continue
		}
		parsedTime, err := time.Parse(time.RFC3339, line[2])
		if err != nil {
			log.Fatal("Failed to parse time:", err)
		}
		timeDifference := timediff.TimeDiff(parsedTime)
		formattedString := fmt.Sprintf("%s\t%s\t%s\t%s\n", line[0], line[1], timeDifference, line[3])
		writer.Write([]byte(formattedString))
	}
	writer.Flush()
	defer file.Close()
}
