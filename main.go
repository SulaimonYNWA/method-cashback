package main

	import (
		"fmt"
	)

	type User struct{
		id int64
		name string
		surname string
		login string
		password string
		sex string
		age int64
		userInfo string
	}

	// func ChangeInfo(person User, info string){
	// 	person.userInfo = info
	// 	fmt.Println(person)
	// }
	func (person *User) changeUserInfo(text string) {		//slovo PERSON pridumyvaem samy. 
		person.userInfo = text
		// newUserSecond.userInfo = ....
		fmt.Println( "in function ",person)		//new value. it must be string as well
	}

func maiasn() {
	// var a int64 = 10 
	// a:= int64(10)

	// newUser := User{
	// 	id: 0,
	// 	name: "",
	// 	surname: "",
	// 	login: "",
	// 	password: "",
	// 	sex: "",
	// 	age : 0,
	// 	userInfo: "",
	// }
	 var newUserSecond User

		newUserSecond.id = 1
		newUserSecond.name = "Name"
		newUserSecond.surname = "surname"
		newUserSecond.login = "login"
		newUserSecond.password = "pass"
		newUserSecond.sex = "male"
		newUserSecond.age = 18
		newUserSecond.userInfo = "info"
		fmt.Println(newUserSecond)
		newUserSecond.changeUserInfo("changed ingormation")
		// changeUserInfo(&newUserSecond, "changed value and info")
		fmt.Println(newUserSecond)
	
}