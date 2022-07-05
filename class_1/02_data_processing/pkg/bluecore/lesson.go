package bluecore

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/afero"
	"io/ioutil"
	"os"
)

func OpenFile(cfg Config) (f afero.File, err error) {
	// Open File os.O_RDONLY 0644
	file, err := cfg.Fs.OpenFile(cfg.FilePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("problem opening file(%s) err(%v)", cfg.FilePath, err)
	}

	return file, nil
}

func ReadFileData(cfg Config) (data []byte, err error) {
	file, err := OpenFile(cfg)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file data
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Blank return returns the named return values
	return
}

func Lesson(cfg Config) error {
	fmt.Println("Lesson 02 Data Processing")

	// Read file data
	data, err := ReadFileData(cfg)
	if err != nil {
		return err
	}

	//return ProcessJsonRaw(data)
	_, err = ProcessJsonJC(data)
	return err
}

func ProcessJsonRaw(data []byte) error {
	// interface{}
	// map[string]obj {}
	// make(map[string]interface{}, 100)
	// make([]byte, 100, 0)
	m := map[string]interface{}{}

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	fmt.Printf("classification: %s\n", m["classification"])

	return nil
}

func ProcessJsonJC(data []byte) (Journey, error) {
	jc := Journey{}

	if err := json.Unmarshal(data, &jc); err != nil {
		return jc, err
	}

	fmt.Printf("obj: %+v\n", jc)

	if n, ok := jc.Name.(int); ok {
		fmt.Printf("obj name as int: %d\n", n)
	} else if n, ok := jc.Name.(string); ok {
		fmt.Printf("obj name as string: %s\n", n)
	}

	out, err := OutputJSON(jc)
	if err != nil {
		return jc, err
	}

	fmt.Printf("remarshalled: %s\n", out)

	return jc, nil
}

func OutputJSON(jc Journey) (data string, err error) {
	b, err := json.Marshal(&jc)
	if err != nil {
		return "", err
	}

	data = string(b)
	return data, nil

}

type Draft struct{}
type Published struct{}

func CastToString(v interface{}) (string, error) {
	switch v.(type) {
	case Draft:
	case Published:
	default:
		return "", fmt.Errorf("asdfasdf")
	}
	return "", nil
}

// Reference JSON
// ../../data/journey_campaign.json
