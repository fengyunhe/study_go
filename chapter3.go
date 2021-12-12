package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Printf("%.2f liters needed \n", 4.2*3.0/10.0)

	fmt.Printf("The %s cost %d certs each.\n", "test", 23) // The test cost 23 certs each.
	fmt.Printf("A float: %f\n", 3.1415)                    // A float: 3.141500
	fmt.Printf("An integer: %d\n", 15)                     // An integer: 15
	fmt.Printf("A string: %s\n", "hello")                  // A string: hello
	fmt.Printf("A boolean: %t\n", false)                   // A boolean: false
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)      // Values: 1.2      true
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)   // Values: 1.2 "\t" true
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)       // Types: float64 string bool
	fmt.Printf("Percent sign: %%\n")                       // Percent sign: %
	fmt.Printf("%#v %#v %#v\n", "", "\n", "\t")            // "" "\n" "\t"

	//////

	fmt.Printf("%5s | %-10s \n", "Book", "Auther")
	fmt.Printf("%5s | %10s \n", "Javascript权威指南", "不知道")
	fmt.Printf("%5s | %10s \n", "Java权威指南", "还是不知道")
	fmt.Printf("%5s | %10s \n", "阿里巴巴java语法", "仍然还是不知道啊")

	fmt.Printf("%5.3f \n", 1234.14151451) //1234.142
	fmt.Printf("%.2f \n", 1234.14151451)  //1234.14

	repeatLine("hello \n", 4)

	dozen := double(6)

	fmt.Println(dozen)

	needed, err := paintNeeded(10, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed \n", needed)

	//point

	x := 1

	fmt.Println(&x) //print point address

	y := &x //use a variable ref point

	fmt.Println(*&x) //print value at point

	*y = 4 //change value at point

	fmt.Println(x)

	//////、
	z := 3
	double2(&z)
	fmt.Println(z) //6
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf(line)
	}
}

func double(number float64) float64 {
	return number * 2
}

func double2(number *int) {
	*number *= 2
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	return width * height, nil

}
