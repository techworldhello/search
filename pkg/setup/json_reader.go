package setup

import (
	"io/ioutil"
	"os"
)

// JSONReader holds the file to be read
type JSONReader struct {
	FileName string
}

// ReadFile opens the given file and reads all data from a io.Reader until EOF
func (j JSONReader) ReadFile() ([]byte, error) {
	file, err := os.Open(j.FileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			println(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
