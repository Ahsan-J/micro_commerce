package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// FailOnError is a general helper to deal with error logging
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ImportJSON will import string path from JSON and map it into type given
func ImportJSON(importPath string, typeInterface interface{}) {
	jsonFile, err := os.Open(importPath)
	defer jsonFile.Close()

	// if we os.Open returns an error then handle it
	FailOnError(err, "Cannot import JSON "+importPath)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &typeInterface)
}
