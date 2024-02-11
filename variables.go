package main

import "fmt"

const LoginToken string = "hhhh" //public = L (capital)

func main(){
	fmt.Println("Variables")

	var username  string = "itesh"
	println(username)
	fmt.Printf("variables is of type %T \n",username) //placeholder 

	var isLoggedin  bool = true
	println(isLoggedin)
	fmt.Printf("variables is of type %T \n",username) //placeholder 

	var smallvalue  uint8 = 255
	println(smallvalue)
	fmt.Printf("variables is of type %T \n",smallvalue) //placeholder 

	var smallvaluefloat  float32 = 255.333348565748
	println(smallvaluefloat)
	fmt.Printf("variables is of type %T \n",smallvaluefloat) //placeholder 

	numberofuserss := 3000
	fmt.Println(numberofuserss ) //inside the method

	fmt.Println(LoginToken)
}
