package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("could not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	//Scan() returns true until no lines are left
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check for file scan errors
	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("reading the file failed")

	}

	// file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to json")
	}

	// file.Close()
	return nil
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
