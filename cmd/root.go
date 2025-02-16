package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const TIMEFORM = "02/01/2006 - 15:04"

type Entry struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Tel        string `json:"telephone"`
	LastAccess string `json:"lastaccess"`
}

type phoneBook []Entry

var JSONFILE = "./storage/contacts.json"
var index map[string]int
var data = phoneBook{}

func DeSerialize(slice any, data io.Reader) error {
	d := json.NewDecoder(data)
	return d.Decode(slice)
}

func Serialize(slice any, data io.Writer) error {
	e := json.NewEncoder(data)
	return e.Encode(slice)
}

func readJSONFile(filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		return err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := DeSerialize(&data, f); err != nil {
		return err
	}
	return nil
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := Serialize(&data, f); err != nil {
		return err
	}
	return nil
}

func createIndex() {
	index = make(map[string]int, len(data))
	for i, entry := range data {
		index[entry.Tel] = i
	}
}

func makeTime() string {
	return time.Now().Format(TIMEFORM)
}

func NewEntry(n, s, t string) *Entry {
	if n == "" || s == "" {
		return nil
	}
	la := makeTime()
	return &Entry{
		Name:       n,
		Surname:    s,
		Tel:        t,
		LastAccess: la,
	}
}

func setJSONFILE() error {
	path := os.Getenv("PHONEBOOK")
	if path != "" {
		JSONFILE = path
	}
	if _, err := os.Stat(JSONFILE); err != nil {
		fmt.Printf("Creating: %s\n", JSONFILE)
		f, err := os.Create(JSONFILE)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	fileInfo, err := os.Stat(JSONFILE)
	if err != nil {
		return err
	}
	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", JSONFILE)
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "phonebook",
	Short: "A phonebook application.",
	Long:  `This is a phonebook application that uses JSON file to record.`,
}

func Execute() {
	if err := setJSONFILE(); err != nil {
		fmt.Println(err)
		return
	}
	if err := readJSONFile(JSONFILE); err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	createIndex()

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
