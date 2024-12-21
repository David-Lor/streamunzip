package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/zhyee/zipstream"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ZipExtractor struct {
	ZipStream  io.Reader
	OutputPath string

	zipReader *zipstream.Reader
}

func (e *ZipExtractor) ExtractInteractively() error {
	// https://pkg.go.dev/github.com/zhyee/zipstream#section-readme
	var saveAll bool
	e.zipReader = zipstream.NewReader(e.ZipStream)

fileLoop:
	for {
		file, err := e.zipReader.GetNextEntry()
		if err == io.EOF {
			log.Println("Zip file read complete!")
			break fileLoop
		}
		if err != nil {
			return fmt.Errorf("unable to get next entry: %v", err)
		}

		if file.IsDir() {
			continue fileLoop
		}

	optionLoop:
		for {
			option := ""
			if !saveAll {
				fmt.Println("What to do with this file?")
				fmt.Println("\tName:", file.Name)
				fmt.Println("\tSize:", humanize.Bytes(file.UncompressedSize64))
				fmt.Printf("[%s]Save [%s]Skip [%s]Save all [%s]Quit\n", OptionSave, OptionSkip, OptionSaveAll, OptionQuit)
				option = ReadInput("")
			}

			option = strings.TrimSpace(option)
			option = strings.ToLower(option)
			if option == "a" {
				saveAll = true
				log.Println("Save all")
			}

			switch {
			case option == "s" || saveAll:
				log.Println("Saving file", file.Name)
				if err := e.saveFile(file); err != nil {
					log.Println("ERROR saving file", file.Name, ":", err)
				}

			case option == "k":
				log.Println("Skip file", file.Name)

			case option == "q":
				log.Println("Quit")
				return nil

			default:
				continue optionLoop
			}

			break optionLoop
		}
	}

	return nil
}

func (e *ZipExtractor) saveFile(fileEntry *zipstream.Entry) error {
	inputFile, err := fileEntry.Open()
	if err != nil {
		return fmt.Errorf("failed opening compressed file: %v", err)
	}
	defer inputFile.Close()

	outputFilePath := filepath.Join(e.OutputPath, fileEntry.Name)
	if err := createParentDir(outputFilePath); err != nil {
		return fmt.Errorf("failed creating parent directory for output file %s: %v", outputFilePath, err)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed creating output file %s: %v", outputFilePath, err)
	}
	defer outputFile.Close()

	outputSize, err := io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("failed writing output file %s: %v", outputFilePath, err)
	}

	uintOutputSize := uint64(outputSize)
	if uintOutputSize != fileEntry.UncompressedSize64 {
		return fmt.Errorf("mismatch between sizes of input file (%s) and written file in %s (%s)", humanize.Bytes(uintOutputSize), outputFilePath, humanize.Bytes(fileEntry.UncompressedSize64))
	}

	if err := os.Chtimes(outputFilePath, time.Time{}, fileEntry.Modified); err != nil {
		return fmt.Errorf("failed setting modification time for file %s: %v", outputFilePath, err)
	}

	return nil
}

func createParentDir(filePath string) error {
	dirPath := filepath.Dir(filePath)
	return os.MkdirAll(dirPath, os.ModePerm)
}
