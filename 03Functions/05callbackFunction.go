package main

//CallbackFunc sds, I could not call this function is because arguments need to be passed while callingcd 03
func CallbackFunc(num int) int {

	return num
}

//SendAdder sds
func SendAdder(intSlice []int, callbackFunc func(i int)) {

	for _, num := range intSlice {

		callbackFunc(num)
	}
}
