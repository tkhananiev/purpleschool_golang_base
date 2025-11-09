package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("./files/data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
func WriteFile(content []byte, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file has been created")

}
