package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Client that can have the responsibility of read and format an output
type Client struct {
}

// ArrayInt parse a every line of a file to a slice of int
func (c *Client) ArrayInt(path string) ([]int, error) {
	elements := make([]int, 0)
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error passing to int. %w", err)
		}

		elements = append(elements, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elements, nil
}
