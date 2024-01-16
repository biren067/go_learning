package variable

import ( "fmt"

)

//  we cannot use outside of function with shorthand := to declare variable
// var surname ="kumar" // work
// surname:="kumar" // error
func variable(){
	// // ========== one way to declare variable ==========
	// var name string = "bond"
	// var name1 string
	// var name2 string = "james"
	// fmt.Println("How are you, Mr. ",name1,name,name2)
	// name="james"
	// fmt.Println(name)

	// ================= 2nd way =============
	// name :="bond"
	// fmt.Println("itmes::",name)

	// // =========== testing global variable =====
	// fmt.Println("Global variable::",surname)
	// log.Println("bond")

}
