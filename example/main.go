package main

import (
	"boyzgenk/playground/csvwriter/yoso"
	"fmt"
	"log"
	"os"
)

func main() {
	header := []string{
		"id",
		"name",
		"gender",
	}

	cfg := yoso.Config{
		Header:       header,
		Separator:    ';',
		FileName:     "test",
		UsePart:      true,
		LimitPerPart: 50,
	}

	dep, err := yoso.NewWriter(cfg)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}

	for i := 0; i < 100; i++ {
		data := []string{
			fmt.Sprintf("%d", i),
			"mboh",
			"lanang",
		}

		err = dep.Write(data)
		if err != nil {
			log.Println(err.Error())
			os.Exit(0)
		}
	}

	defer dep.Close()
}
