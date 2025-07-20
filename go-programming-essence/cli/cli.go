package main

import "flag"

func main() {
	max := flag.Int("max", 10, "maximum number of items to process")
	name := flag.String("name", "default", "name of the process")

	var name1 string
	var max1 int
	flag.StringVar(&name1, "name1", "default1", "name of the first process")
	flag.IntVar(&max1, "max1", 20, "maximum number of items for the first process")
	flag.Parse()

	println("Max:", *max)
	println("Name:", *name)
	for _, arg := range flag.Args() {
		println("Argument:", arg)
	}
}
