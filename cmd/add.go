package cmd

import (
	"encoding/csv"
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
	example usage tasks add <description>. It only takes one argument which is the desctription`,
	Run: add,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	// Open the file in append mode
	file, err := os.OpenFile("output.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Failed to close file:", err)
		}
	}()

	// Check if the file is empty to decide if we need to write headers
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get file stats:", err)
	}

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// If the file is empty, write headers
	if stat.Size() == 0 {
		headers := []string{"ID", "Description", "CreatedAt", "IsComplete"}
		if err := writer.Write(headers); err != nil {
			log.Fatal("Failed to write headers:", err)
		}
		writer.Flush()
	}

	// Read existing records to determine the next ID
	file.Seek(0, 0) // Rewind to the beginning of the file to read all records
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}

	lastID := len(lines)

	// Create a new data entry
	data := Data{
		ID:          lastID, // ID will be between 1 and 9
		Description: args[0],
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	// Write the new record to the CSV
	record := []string{
		strconv.FormatInt(int64(data.ID), 10),
		data.Description,
		data.CreatedAt.Format(time.RFC3339),
		strconv.FormatBool(data.IsComplete),
	}

	if err := writer.Write(record); err != nil {
		log.Fatal("Failed to write record:", err)
	}

	// Ensure the data is flushed to the file
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal("Failed to flush writer:", err)
	}
}
