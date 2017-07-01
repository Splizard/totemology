package main

import "../"
import "fmt"

func main() {
	t := totemology.NewTotem()
	for i:=0; i < 50; i++ {
		t.WriteImage(fmt.Sprint(i)+".png")
		t.GrowMissile()
	}
}
