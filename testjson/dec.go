package main

import "encoding/json"
import "fmt"
import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	outfile, err := os.OpenFile("b.json", os.O_WRONLY, os.ModePerm)
	check(err)

	enc := json.NewEncoder(outfile)
	d := map[string]int{"first": 1, "second": 2}
	enc.Encode(d)

	infile, err := os.Open("b.json")
	check(err)
	dec  := json.NewDecoder(infile)
	dec.Decode(&d)
	fmt.Println(d)

	/*
		file, err := os.Open("b.json")
		check(err)

		b := make([]byte, 2^10)
		n, err := file.Read(b)
		check(err)

		fmt.Println("Read ", n, " bytes")

		var dat string
		err = json.Unmarshal(b, &dat)
		check(err)

		fmt.Println(string(dat))

	*/
}
