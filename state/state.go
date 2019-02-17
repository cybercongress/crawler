package state

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type IndexState struct {
	path              string `json:"-"`
	SubmitLinksOffset int64  `json:"submitLinksOffset"`
}

func Load(path string) (s IndexState, err error) {

	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = createEmptyConfigFile(path); err != nil {
			return
		}
	}

	configAsBytes, _ := ioutil.ReadFile(path)
	err = json.Unmarshal(configAsBytes, &s)
	s.path = path
	return
}

func (s IndexState) Save() error {

	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}

	configFile, err := os.OpenFile(s.path, os.O_RDWR, 0666)
	defer configFile.Close()
	if err != nil {
		return err
	}

	if err = configFile.Truncate(0); err != nil {
		return err
	}

	if _, err = configFile.Seek(0, 0); err != nil {
		return err
	}

	if _, err = configFile.Write(bytes); err != nil {
		return err
	}
	return configFile.Sync()
}

func createEmptyConfigFile(fileName string) error {

	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	s := IndexState{
		path:              fileName,
		SubmitLinksOffset: 0,
	}
	return s.Save()
}
