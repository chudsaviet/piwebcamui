package main

import (
	"fmt"
	pb "github.com/chudsaviet/piwebcamui/v2/build/gen/proto"
)

func main() {
	fmt.Println("Hello world!")

	x := pb.WebcamConfig{
		InputResX: 800,
		InputResY: 600,
		InputFps:  29.97,
	}
	fmt.Println(x.String())
}
