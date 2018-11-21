package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

//check whether the bytes is LittleEndian or BigEndian
//mostly it is the LittleEndian
func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func main() {
	//convert float to bytes
	bytes := Float64bytes(math.Pi)
	fmt.Println(bytes)

	//read float from bytes
	float := Float64frombytes(bytes)
	fmt.Println(float)
}
