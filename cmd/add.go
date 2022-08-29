package cmd

import (
	"github.com/dilly3/task-manager/functiondiary"
	"github.com/spf13/cobra"
)

// addCmd represents the add command

var addCmd = &cobra.Command{

	Use:   "add",
	Short: " Add Item to Your Todo List",
	Long: ` This command followed by a string , eg : [add Wash clothes] ; adds Wash Cloth to your list of tasks:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		functiondiary.Add(args...)

	},
}

//func Add(args ...string) []diary.Item {
//	defer diary.Cleanup()
//	if len(args) > 8 {
//		fmt.Println("task wont be Added")
//		panic("String too long")
//
//	}
//	itemsArr, readErr := diary.Reader() // read from csv file
//
//	if readErr != nil { // check for error
//		log.Fatal("reader failed to read ", readErr)
//	}
//	if len(itemsArr) < 1 {
//		itemsArr = append(itemsArr, diary.Item{1, diary.CreateString(args), "undone",
//			diary.TimeFormat(), " ", " "})
//	} else {
//		itemsArr = append(itemsArr, diary.Item{len(itemsArr) + 1, diary.CreateString(args), "undone",
//			diary.TimeFormat(), " ", " "})
//	}
//
//	writerErr := diary.Writer(itemsArr)
//	if writerErr != nil {
//		log.Fatal(writerErr)
//	}
//	fmt.Printf("Task Added \n")
//	return itemsArr
//}

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
