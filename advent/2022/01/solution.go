package solution

import (
    "bufio"
    "fmt"
    "os"
)

const INPUT_FILE string = "input.txt"

func main() {
    readfile, err := os.Open(INPUT_FILE)

    if (err != nil) {
        fmt.Println(err)
    }
    scanner := bufio.NewScanner(readfile)
    scanner.Split(bufio.ScanLines)
    
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    readfile.Close()
}
