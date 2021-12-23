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

// ScanFile takes filepath, reads file content, returns []string
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

// WriteInFile takes []string and filepath, creates new file, writes in a new file
func WriteInFile(data []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, datum := range data {
		_, err := file.WriteString(datum + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// SeparateNames separate full name to first name and last name
func SeparateNames(names []string) (firstNames, lastNames []string) {
	for _, name := range names {
		separateHelper := strings.Split(name, " ")
		firstNames = append(firstNames, separateHelper[0])
		lastNames = append(lastNames, separateHelper[1])
	}
	return firstNames, lastNames
}

//DeleteDubs takes []string and removes duplicates
func DeleteDubs(names []string) []string{
	keys := make(map[string]bool)
	var namesWithoutDubs []string
	for _, name := range names {
		if !keys[name]{
			keys[name] = true
			namesWithoutDubs = append(namesWithoutDubs, name)
		}
	}
	return namesWithoutDubs
}

// PageInput scans file by filepath, writes content on webpage
func PageInput(w http.ResponseWriter, filePath string ) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		_, err := fmt.Fprintf(w, scanner.Text() + "\n")
		if err != nil {
			return err
		}
	}
	if scanner.Err() == err {
		return err
	}
	return nil
}

// GenerateNames takes two []string and quantity how many times new full name should be generated
// returns []string with new full names
func GenerateNames(firstNames, lastNames []string, numOfNames int) (generatedNames []string){
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numOfNames; i++ {
		firstNamesNumRand:= rand.Intn(len(firstNames)-1)
		lastNamesNumRand:= rand.Intn(len(lastNames)-1)
		fullName := firstNames[firstNamesNumRand] + " "+  lastNames[lastNamesNumRand]
		generatedNames = append(generatedNames, fullName)
	}
	return generatedNames
}