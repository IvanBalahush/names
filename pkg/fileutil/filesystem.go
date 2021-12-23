package fileutil

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// ScanFile takes filepath, reads file content
func ScanFile(name string) (data []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return data, nil
}

// WriteInFile creates new file, writes content in a new file
func WriteInFile(data []string, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, item := range data {
		_, err := file.WriteString(item + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Separator separates every string in data on two parts, each part inserts in turn into a new slice (separatedSlice1, separatedSlice2)
func Separator(data []string, separator string) (separatedSlice1, separatedSlice2 []string) {
	for _, item := range data {
		separateHelper := strings.Split(item, separator)
		separatedSlice1 = append(separatedSlice1, separateHelper[0])
		separatedSlice2 = append(separatedSlice2, separateHelper[1])
	}

	return separatedSlice1, separatedSlice2
}

// DeleteDubs creates new slice and via map[string]bool appends new item to the slice
// repeating element doesn't write in the slice
func DeleteDubs(data []string) []string {
	keys := make(map[string]bool)
	var dataWithoutDubs []string

	for _, item := range data {
		if !keys[item] {
			keys[item] = true
			dataWithoutDubs = append(dataWithoutDubs, item)
		}
	}

	return dataWithoutDubs
}

// PageInput scans file by filepath
func PageInput(w http.ResponseWriter, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err := fmt.Fprintf(w, scanner.Text()+"\n")
		if err != nil {
			return err
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

// Generator quantity means how many times new generations should be made
func Generator(slice1, slice2 []string, numOfGenerations int) (generatedSlice []string){
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numOfGenerations; i++ {
		NumRandSlice1 := rand.Intn(len(slice1) - 1)
		NumRandSlice2 := rand.Intn(len(slice2) - 1)
		fullGeneration := slice1[NumRandSlice1] + " " + slice2[NumRandSlice2]
		generatedSlice = append(generatedSlice, fullGeneration)
	}

	return generatedSlice
}
