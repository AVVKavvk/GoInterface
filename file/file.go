package file

import (
	"fmt"
	"io"
	"os"
)

func Dummy() {
	content, err := os.Open("/home/vipin/Go/1/file/1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer content.Close()

	info, err := content.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println(info)
	b := make([]byte, 5)
	for {
		n, err := content.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		fmt.Println(string(b[:n]))
	}
	fmt.Println()
}
