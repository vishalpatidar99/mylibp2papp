package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadString(r *bufio.Reader) (string, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	if str == "\n" {
		return "", nil
	}

	fmt.Printf("received raw data from stream: %s \n", str)

	return str, nil
}

func WriteString(w *bufio.Writer, msg string) (int, error) {

	fmt.Printf("writing raw data to stream: %s \n", msg)

	n, err := w.WriteString(fmt.Sprintf("%s\n", msg))
	if err != nil {
		fmt.Printf("failed to write to buffer: %v \n", err)
		return 0, err
	}
	err = w.Flush()
	if err != nil {
		fmt.Printf("failed to flush buffer: %v \n", err)
		return 0, err
	}
	return n, nil
}

func ReadCipherText(location string) (string, error) {
	file, err := os.Open(location)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(data[:n]), nil
}
