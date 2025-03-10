package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// readfile with error handling
// read from file
func GetFloatFromFile(fileName string) (float64, error) {
	fmt.Println("Reading balance from file")
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0.0, errors.New("error reading file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return 0.0, errors.New("failed to parse stored value")
	}
	return value, nil
	}

// write to file
func WriteFloatToFile(value float64, fileName string) {
	fmt.Println("Writing balance to file")
	valueWrite := fmt.Sprintf("%.2f", value)
	// 0644 permisssions
	os.WriteFile(fileName, []byte(valueWrite), 0644)
}