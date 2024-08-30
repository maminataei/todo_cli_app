package utilities

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type IO struct {
	scanner *bufio.Scanner
}

func NewIO() IO{
	return IO{ 
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (io IO) ReadStr() (string, error) {
	for io.scanner.Scan() {
		return io.scanner.Text(), nil
	}
	return "", errors.New("invalid input")
}

func (io IO) ReadNumber() (int, error) {
	str, inputStrErr := io.ReadStr()
	if inputStrErr != nil {
		return -1, inputStrErr
	}
	number, castingErr := strconv.Atoi(str)
	if castingErr != nil {
		return -1, castingErr
	}
	return number, nil
}

func (io IO) ReadBool() (bool, error) {
	str, inputStrErr := io.ReadStr()
	if inputStrErr != nil {
		return false, inputStrErr
	}
	bool, castingErr := strconv.ParseBool(str)
	if castingErr != nil {
		return false, castingErr
	}
	return bool, nil
}