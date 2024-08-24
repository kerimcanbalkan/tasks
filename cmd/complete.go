/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "This command marks given task as complete",
	Long:  `Example usage tasks complete <task ID>`,
	Run:   complete,
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func complete(cmd *cobra.Command, args []string) {
	id := args[0]
	var keepers [][]string
	taskUpdated := false

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

	for _, line := range lines {
		if line[0] != id {
			keepers = append(keepers, line)
		} else {
			updated := []string{line[0], line[1], line[2], "true"}
			keepers = append(keepers, updated)
			taskUpdated = true
		}
	}

	if !taskUpdated {
		fmt.Printf("Task with given id %s not found", id)
		return
	}

	f, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Could not create the file", err)
	}
	writer := csv.NewWriter(f)

	for _, el := range keepers {
		if err := writer.Write(el); err != nil {
			log.Fatal("Failed to write record:", err)
		}

	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal("Failed to flush writer:", err)
	}

}
