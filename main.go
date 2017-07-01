/*
	This program generates totem poles which contain recognisable patterns.
	
	Based off Eulers Number Triangle.
	http://mathworld.wolfram.com/EulersNumberTriangle.html
*/
package totemology


import "fmt"
import "math/big"
import "image"
import "image/png"
import "os"
import "image/color"

//A totem is an object which can be grown in different ways.
//See the respective grow methods.
type Totem struct {
	count int64
	row []*big.Int
}

//These totems have shapes like missiles.
//This is the original euler's triangle.
func (totem *Totem) GrowMissile() {
	next := make([]*big.Int, len(totem.row)+1)
	next[0] = big.NewInt(1)
	for x := int64(1); x < int64(len(totem.row)); x++ {
		next[x] = big.NewInt(0)
		left := big.NewInt(0)
		left.Mul( totem.row[x-1], big.NewInt(totem.count-x+1))
		
		right := big.NewInt(0)
		right.Mul(totem.row[x], big.NewInt(1+x))
		
		
		next[x].Add(left, right)
	} 
	next[len(next)-1] = big.NewInt(1)
	totem.row = next
	totem.count ++
}

//These totems have shapes which look like arrows.
func (totem *Totem) GrowArrow() {
	next := make([]*big.Int, len(totem.row)+1)
	next[0] = big.NewInt(1)
	for x := int64(1); x < int64(len(totem.row)); x++ {
		next[x] = big.NewInt(0)
		left := big.NewInt(0)
		left.Mul( totem.row[x-1], big.NewInt(totem.count))
		
		right := big.NewInt(0)
		right.Mul(totem.row[x], big.NewInt(totem.count))
		
		
		next[x].Add(left, right)
	} 
	next[len(next)-1] = big.NewInt(1)
	totem.row = next
	totem.count ++
}

//You can find faces in this totem.
func (totem *Totem) GrowFace() {
	next := make([]*big.Int, len(totem.row)+1)
	next[0] = big.NewInt(1)
	for x := int64(1); x < int64(len(totem.row)); x++ {
		next[x] = big.NewInt(0)
		left := big.NewInt(0)
		left.Mul( totem.row[x-1],  big.NewInt(1))
		
		right := big.NewInt(0)
		right.Mul(totem.row[x],  big.NewInt(1))
		
		
		next[x].Add(left, right)
	} 
	next[len(next)-1] = big.NewInt(1)
	totem.row = next
	totem.count ++
}

//This totem can have patterns which resemble rockets.
func (totem *Totem) GrowRocket(mod int64) {
	next := make([]*big.Int, len(totem.row)+1)
	next[0] = big.NewInt(1)
	for x := int64(1); x < int64(len(totem.row)); x++ {
		next[x] = big.NewInt(0)
		left := big.NewInt(0)
		left.Mul( totem.row[x-1],  big.NewInt(totem.count-x+mod))
		
		right := big.NewInt(0)
		right.Mul(totem.row[x],  big.NewInt(x+mod))
		
		
		next[x].Add(left, right)
	} 
	next[len(next)-1] = big.NewInt(1)
	totem.row = next
	totem.count ++
}

//This totem can have shapes which can resemble bodies.
//You can provide a value to modify the images.
func (totem *Totem) GrowBody(mod int64) {
	next := make([]*big.Int, len(totem.row)+1)
	next[0] = big.NewInt(1)
	for x := int64(1); x < int64(len(totem.row)); x++ {
		next[x] = big.NewInt(0)
		left := big.NewInt(0)
		left.Mul( totem.row[x-1],  big.NewInt(mod))
		
		right := big.NewInt(0)
		right.Mul(totem.row[x],  big.NewInt(mod))
		
		
		next[x].Add(left, right)
	} 
	next[len(next)-1] = big.NewInt(1)
	totem.row = next
	totem.count ++
}

//This totem has shapes which resemble a being.
//You can provide values a and b to modify the growth.
func (totem *Totem) GrowBeing(a, b int64) {
	next := make([]*big.Int, len(totem.row)+1)
	next[0] = big.NewInt(a)
	for x := int64(1); x < int64(len(totem.row)); x++ {
		next[x] = big.NewInt(0)
		left := big.NewInt(0)
		left.Mul( totem.row[x-1],  big.NewInt(b))
		
		right := big.NewInt(0)
		right.Mul(totem.row[x],  big.NewInt(b))
		
		
		next[x].Add(left, right)
	} 
	next[len(next)-1] = big.NewInt(a)
	totem.row = next
	totem.count ++
}

//Returns the current iteration of the totem as a string.
func (totem *Totem) String() string {
	var r string
	for _, value := range totem.row {
		r += fmt.Sprint(" ",value)
	}
	return r
}

//Write the current iteration of the totem to disc at path "path".
func (totem *Totem) WriteImage(path string) {
	width, height := len(totem.row), totem.row[(len(totem.row)-1)/2].BitLen()
	img := image.NewGray(image.Rect(0, 0, width, height))
	
	for x, number := range totem.row {
		for y := 0; y <= number.BitLen(); y++ {
			if number.Bit(y) == 1 {
				img.Set(x, height-y-1, color.Gray{255})
			} else {
				img.Set(x, height-y-1, color.Gray{0})
			}
			
		}
	}
	
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, img)
}

//Return a totem object.
func NewTotem() *Totem {
	totem := new(Totem)
	totem.row = make([]*big.Int, 1)
	totem.row[0] = big.NewInt(1)
	totem.count++
	return totem
}
