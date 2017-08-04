package bootstrap

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Parser must implement Parser Json
type parser interface {
	ParseJSON([]byte) error
}

// LoadJSON config file
func LoadJSON(configFile string, p parser) {
	log.Println("call loadJSON func")
	var err error
	var absPath string
	var input = io.ReadCloser(os.Stdin)
	if absPath, err = filepath.Abs(configFile); err != nil {
		log.Fatalln(err)
	}

	if input, err = os.Open(absPath); err != nil {
		log.Fatalln(err)
	}

	//Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	//Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Println("json config error")
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
