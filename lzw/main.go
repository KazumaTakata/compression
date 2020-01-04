package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func encode(input_data []byte) []uint16 {

	m := make(map[string]uint16)

	for i := 0; i < 256; i++ {
		key := string([]byte{byte(i)})
		m[key] = uint16(i)
	}

	output := []uint16{}
	var ent_index uint16 = 256

	cur_string := []byte{}

	for len(input_data) != 0 {
		ch := input_data[0]
		input_data = input_data[1:]
		key := append(cur_string, ch)

		if _, ok := m[string(key)]; ok {
			cur_string = key
		} else {
			output = append(output, m[string(cur_string)])
			m[string(key)] = ent_index
			ent_index += 1
			cur_string = []byte{ch}

		}

	}

	for k, v := range m {
		if v > 255 {
			fmt.Printf("%s:%d\n", k, v)
		}
	}

	return output

}

func decode(input_data []uint16) []byte {

	m := make(map[uint16][]byte)
	var ent_index uint16 = 256

	for i := 0; i < 256; i++ {
		m[uint16(i)] = []byte{byte(i)}
	}
	output := []byte{}
	entry := []byte{}
	var prevcode uint16
	var curcode uint16

	prevcode = input_data[0]
	input_data = input_data[1:]

	decoded := m[prevcode]
	output = append(output, decoded...)

	for len(input_data) != 0 {
		curcode = input_data[0]
		input_data = input_data[1:]
		entry = m[curcode]
		output = append(output, entry...)
		ch := entry[0]
		m[ent_index] = append(m[prevcode], ch)
		ent_index += 1
		prevcode = curcode
		for k, v := range m {
			if k > 255 {
				fmt.Printf("%d:%s\n", k, v)
			}
		}

		fmt.Printf("%v%d\n", output, ent_index)

	}

	return output

}

func main() {

	input_data, _ := ioutil.ReadFile("sample.txt")
	output := encode(input_data)
    fmt.Printf("=======\n")
	/* m := make(map[string]uint16)*/

	//for i := 0; i < 256; i++ {
	//key := string([]byte{byte(i)})
	//m[key] = uint16(i)
	//}

	//output := []uint16{}
	//var ent_index uint16 = 256
	//check(err)

	//cur_string := []byte{}

	//for len(input_data) != 0 {
	//ch := input_data[0]
	//input_data = input_data[1:]
	//key := append(cur_string, ch)

	//if _, ok := m[string(key)]; ok {
	//cur_string = key
	//} else {
	//output = append(output, m[string(cur_string)])
	//m[string(key)] = ent_index
	//ent_index += 1
	//cur_string = []byte{ch}

	//}

	//}
	fmt.Printf("%v", output)

	decoded := decode(output)

	fmt.Printf("%v", decoded)

}
