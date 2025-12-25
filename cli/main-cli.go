package main

import (
	"fmt"
	"log"
	"os"

	"example.com/soda_counter-app/cli/pkg/cli"
	"example.com/soda_counter-app/cli/pkg/fileops"
	ascii "github.com/common-nighthawk/go-figure"
)

// main function
func main() {

	ascii.NewFigure("TBHome - Soda  counter", "doom", true).Print()

	// Get init path
	jsonFilePath, err := fileops.GetInitPath()

	if err != nil {
		log.Fatal("Error retrieving init-path: ", err)
	}

	// Create sodaCounter data variable
	sodaCounter_data := make(map[string]uint)

	// Print the path to the JSON file
	//fmt.Println("JSON file path is:", jsonFilePath)
	fmt.Printf("\nWelcome to the Soda Counter App!\n")

	// Validate if JSON file exists, if not create it with dummy data
	_, err = os.Stat(jsonFilePath)
	if os.IsNotExist(err) {
		fmt.Println("Storage file does not exist. Creating a new one.")
		emptyData := make(map[string]uint)
		err := fileops.ExportDataToFile(jsonFilePath, emptyData)
		if err != nil {
			log.Fatal("Failed to create storage file:", err)
		}
	} else if err != nil {
		log.Fatal("Error checking storage file:", err)
	} else {
		fmt.Println("Storage file found.\n")
	}

	sodaCounter_data, err = fileops.ImportDataFromFile(jsonFilePath)
	if err != nil {
		log.Fatal("Couldn't import data from file: ", jsonFilePath)
	}

	// Start menu
	options := []string{
		"Check storage",
		"Withdraw soda",
		"Deposit soda",
		"InitDummyData",
		"Exit",
	}

	// user choice variable
	var choice int

	for {
		cli.ClearScreen()
		cli.PresentOptions(options)

		// Get user choice
		fmt.Print("\nYour choice: ")
		fmt.Scanln(&choice)

		if choice > 0 || choice < len(options) {

			switch options[choice-1] {
			case "Check storage":
				cli.ClearScreen()
				// Update dataset from datasource
				checkStorage(sodaCounter_data)
			case "Withdraw soda":
				cli.ClearScreen()
				withdrawSoda(sodaCounter_data, jsonFilePath)
			case "Deposit soda":
				cli.ClearScreen()
				fmt.Println("Deposit soda - Feature coming soon!")
				err := depositSoda(sodaCounter_data, jsonFilePath)
				if err != nil {
					fmt.Errorf("An error occurred while depositing soda")
				}
			case "InitDummyData":
				cli.ClearScreen()
				sodaCounter_data = initDummyData()
			default:
				cli.ClearScreen()
				fmt.Println("Exiting the Soda Counter App. Goodbye!")
				return
			}

			cli.ClearScreen()
		} else {
			fmt.Errorf("Invalid choice: %d", choice)
		}
	}
}

// initDummyData initializes dummy soda data for testing purposes
func initDummyData() (dummyData map[string]uint) {
	sodaData := map[string]uint{
		"Coke":        10,
		"Pepsi":       8,
		"Sprite":      15,
		"Fanta":       5,
		"MountainDew": 12,
	}

	return sodaData
}

func checkStorage(data map[string]uint) {
	for soda, quantity := range data {
		fmt.Printf("%s: %d\n", soda, quantity)
	}

	waitForEnter()
}

func withdrawSoda(data map[string]uint, path string) {
	fmt.Println("Withdraw Soda Feature - Coming Soon!\n")
	waitForEnter()
}

func depositSoda(data map[string]uint, path string) (err error) { //TODO Add support for writing to variable instead of directly to datafile (Move to own function?)
	fmt.Println("Deposit Soda Feature\n")

	var sodaName string
	var quantity uint

	fmt.Print("Enter soda name: ")
	fmt.Scanln(&sodaName)
	if err != nil {
		return fmt.Errorf("Failed to read soda name: ", err)
	}

	fmt.Print("Enter quantity to deposit: ")
	fmt.Scanln(&quantity)

	dataFromSource, err := fileops.ImportDataFromFile(path)
	if err != nil {
		log.Fatal("Failed to import data:", err)
	}

	dataFromSource[sodaName] += quantity

	err = fileops.ExportDataToFile(path, dataFromSource)
	if err != nil {
		return fmt.Errorf("Failed to export to datafile: ", err)
	}

	fmt.Printf("Successfully deposited %d of %s.\n", quantity, sodaName)

	waitForEnter()

	return nil
}

// func waitForEnter waits for user to press enter before continuing
func waitForEnter() {
	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln()
}
