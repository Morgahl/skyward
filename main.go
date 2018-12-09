package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	gafhoc, err := os.Create("The Greater Argument for Human Originated Chaos.txt")
	if err != nil {
		fmt.Printf("error writing to writer: %s\n", err)
		return
	}
	defer gafhoc.Close()

	gafhocThesis := bufio.NewWriterSize(gafhoc, 64*1024)
	defer gafhocThesis.Flush()

	startThesis := time.Now()
	defer func() {
		fmt.Println(time.Since(startThesis))
	}()

	for i := 0; i <= hawCount; i++ {
		if err := writeHumansAreWeird(gafhocThesis); err != nil {
			fmt.Printf("error writing to writer: %s\n", err)
			return
		}
	}
}

const (
	haw      = "Humans are weird."
	hawCount = 3756932
)

var hawBytes = []byte(haw)

func writeHumansAreWeird(w io.Writer) error {
	if _, err := w.Write(hawBytes); err != nil {
		return err
	}

	if _, err := w.Write([]byte(" ")); err != nil {
		return err
	}

	return nil
}
