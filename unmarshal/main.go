package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var pr []uint8

type bonus struct {
	Name   string
	Primes string
}
type bon []bonus

var s string

func main() {
	pr := read("primes.json")
	var data bon
	//marshal
	err := json.Unmarshal(pr, &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pr)
	fmt.Println(data)

}
func read(filePath string) []uint8 {
	//open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	/*read file : Golang Specification says :"Read reads up to len(b) bytes from the File."
	  but how do the program know the length of the file ?
	  so set a length that should be greater than it mains contains (here at 1024)
	*/
	d := make([]uint8, 1024)
	n, err := f.Read(d) //n = 1024
	if err != nil {
		log.Fatalln(err)
	}
	/*with a setting that is bigger thanin reality an error may occur
	in connection with 0 refering to ascii table as NULL when unmarshalling.
	To avoid it, we create a counter, range over slice d
	and check when appending to pr if each value (each byte)satisfy
	the condition of not being 0 NULL  */
	var count int
	for count, value := range d {
		if count < n {
			count++
			if value != 0 {
				pr = append(pr, value)
			} else {
				break
			}
		}
	}
	fmt.Println(count)
	return pr
}
