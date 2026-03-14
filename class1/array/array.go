package main
import ("fmt")

func main () {
	var arr1 = [3] int {1,4,9}
	arr2 := [5] string {"Banana","Arroz","Sal","Alho","Pera"}

	fmt.Println("arr1: ",arr1)
	fmt.Println("arr2: ",arr2)

	var arr3 = [...]int {1,4,9}
	arr4 := [...] string {"Banana","Arroz","Sal","Alho","Pera"}

	fmt.Println("arr3: ",arr3)
	fmt.Println("arr4: ",arr4)

	//Initialize Only Specific Elements

	var arr5 = [19] int {3:19,16:23}
	var arr6 = [7] string {2:"Arroz",5:"Banana"}

	fmt.Println("arr5: ",arr5)
	fmt.Println("arr6: ",arr6)
	fmt.Println("Lenght of array 6: ",len(arr6))

	fmt.Println("Index sem dados: ",arr5[10])
	fmt.Println("cap() em array é o mesmo que o len: ",cap(arr6))

	//Slice

	slice1 := [] int {}

	fmt.Println("Slice vazio: ",slice1)
	fmt.Println("len() in Slice vazio: ",len(slice1))
	fmt.Println("cap() in Slice vazio: ",cap(slice1))

	slice1 = append(slice1, 12,16,34,8)
	slice1 = append(slice1, 13)

	fmt.Println("len() in Slice depois de append: ",len(slice1))
	fmt.Println("cap() in Slice depois de append: ",cap(slice1))
	fmt.Println("Slice data: ",slice1)

	//slice from array

	var slice2 = arr4[0:3]
	fmt.Println("slice from a arrary: ",slice2)
	fmt.Println("len() in Slice from a array: ",len(slice2))
	fmt.Println("cap() in Slice from a array: ",cap(slice2))

	var slice3 = arr4[0:2]
	fmt.Println("slice from a arrary: ",slice3)
	fmt.Println("len() in Slice from a array: ",len(slice3))
	fmt.Println("cap() in Slice from a array: ",cap(slice3))

	var slice4 = arr4[1:2]
	fmt.Println("slice from a arrary: ",slice4)
	fmt.Println("len() in Slice from a array: ",len(slice4))
	fmt.Println("cap() in Slice from a array: ",cap(slice4))

	slice5 := append(slice2, slice4...)

	fmt.Println("lice5: ",slice5)
	fmt.Println("len of slice5: ",len(slice5))
	fmt.Println("cap of slice5: ",cap(slice5))

	if len(slice2) > len(arr2) {
		fmt.Println("Hello world: ",len(slice2), ">?", len(arr2))
	}else {
		fmt.Println("World Hello: ", len(slice2), ">?", len(arr2))
	}

	day := 4

	switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wendesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Error")
	}

	for i:=0; i<= 1000; i+=100{
		fmt.Println(i)
	}

	for idx, val := range slice5 {
		fmt.Println(idx,val)
	}

	for _,val := range slice5{
		fmt.Println(val)
	}


}