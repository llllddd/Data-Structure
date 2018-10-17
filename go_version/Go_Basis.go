package main

import (
	"fmt"
)

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}

/*
Modifier function, which take receiver as pointer. Go passes actual
instance of the object to this function (Just like pass by pointer.) Any change
done over the object is reflected on the original object.
*/

type MyInt int

func (data MyInt) increment1() {
	data++
}

func (data *MyInt) increment2() {
	*data = *data + 1
}

/*-----------------------------Slice----------------------------------------------
Go Slice is an abstraction over Array. It actually uses arrays as an underlying
structure. The various operations over slice are:
1. Inbuilt append() function is used to add the elements to a slice. If the size of
underlying array is not enough then automatically a new array is created and
content of the old array is copied to it.
2. The len() function returns the number of elements presents in the slice.
3. The cap() function returns the capacity of the underlying array of the slice.
4.
The copy() function, the contents of a source slice are copied to a
destination slice.
5. Re-slices the slice, the syntax <Slice Name>[start : end] , It returns a slice
object containing the elements of base slice from index start to end-1 Length of
the new slice will be (end - start), and capacity will be cap(base slice) - start.
*/

/*-----------------------------Map--------------------------------------------
Various operation on map:
1. Assignment: < variable>[<key>] = <value>
2. Delete: delete(<variable >, <key>)
3. Access: value, ok = < variable >[<key>] , the first value will have the
   value of key in the map. If the key is not present in the map, it will return
   zero value corresponding to the value data-type. The second argument
   returns whether the map contains the key.
*/

func main() {
	i := 10
	j := &i
	IncrementPassByPointer(&j)
	fmt.Println(*j)

	m := make(map[string]int)
	m["Apple"] = 30
	m["Banana"] = 40
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[", key, "->", val, "]")
	}
	delete(m, "Apple")
	value, ok := m["Apple"]
	fmt.Println("Apple", value, "Present", ok)
}
