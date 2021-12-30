package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

	var image string
	image = "https://source.unsplash.com/random/300x200?sig="
	for i := 0; i < 50; i++ {
		str := fmt.Sprintf("%s%d", image, rand.Intn(100))

		resp, err := http.Head(str)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("Response status:", resp.Status)
	}

}
