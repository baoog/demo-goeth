package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Get(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}

func init() {
	log.Println("Loading config from .env")
	err := Load(".env")
	if err != nil {
		log.Println("Load config failed with error: %v", err.Error())
		return
	}
	log.Println("Load config success !")
}

func Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		text := strings.SplitAfterN(scanner.Text(), "=", 2)
		if len(text) > 1 {
			key := strings.ReplaceAll(text[0], " ", "")[:len(text[0])-1]
			value := strings.ReplaceAll(text[1], " ", "")
			os.Setenv(key, value)
		}
	}

	return nil
}
