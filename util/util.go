package util

import "io/ioutil"

func ReadInput(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	return string(content), err
}
