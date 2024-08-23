/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Data struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command adds task to the To Do list",
	Long: `You will use this command to add tast to the end of the To Do list, it will automatically marked uncomplete,
	example usage tasks add <desctription>. It only takes one argument which is the desctription`,
	Run: add,
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

func add(cmd *cobra.Command, args []string) {
	lastID := 1
	file, err := os.OpenFile("output.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)

	if errors.Is(err, os.ErrNotExist) {
		headers := []string{"ID", "Description", "CreatedAt", "IsComplete"}
		if err := writer.Write(headers); err != nil {
			log.Fatal("Failed to write Headers", err)
		}

	} else {
		lines, err := csv.NewReader(file).ReadAll()
		if err != nil {
			log.Fatal("Couldn't Read the file")
		}
		lastID = len(lines)
	}

	defer file.Close()

	data := Data{
		ID:          lastID,
		Description: args[0],
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	record := []string{
		strconv.FormatInt(int64(data.ID), 10),
		data.Description,
		data.CreatedAt.Format(time.RFC3339),
		strconv.FormatBool(data.IsComplete),
	}

	if err := writer.Write(record); err != nil {
		log.Fatal("Failed to write record", err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal("Failed to flush writer", err)
	}

}
