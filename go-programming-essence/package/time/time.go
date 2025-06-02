package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	s := "2022/12/25 07:42:38"
	d, err := time.Parse("2006/01/02 15:04:05", s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(d)

	var du time.Duration
	du, err = time.ParseDuration("2h30m")
	if err != nil {
		log.Fatalln("error")
		return
	}
	fmt.Println(du)

	t := d.Add(du)
	fmt.Println(t.Format("2006/01/02 15:04:05"))
}
