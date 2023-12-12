package main

import "fmt"

type Student struct {
	Name string
	Age  int8
}

//func (s Student) Introduce(msg string) string {
//	return fmt.Sprintf("%v, my name is %s, my age is %d", msg, s.Name, s.Age)
//}

func (s Student) changeName() {
	s.Name = "Mailo"
}

func (s *Student) changeName2() {
	s.Name = "Mailo"
}

func (s *Student) Greet() {
	fmt.Println("Hello")
	fmt.Println(s.Name)
	fmt.Println(s.Age)
}

func main() {
	// fruits := []string{"apple", "grape", "banana"}

	//var student = Student{
	//	Name: "Agus",
	//	Age:  17,
	//}
	//var student2 = &Student{
	//	Name: "Mailo",
	//	Age:  17,
	//}
	//
	//student.Greet()
	//student2.Greet()

	//student.changeName()
	//fmt.Println(student.Name)
	//fmt.Println("===")
	//student.changeName2()
	//fmt.Println(student.Name)

	//result, err := namaFunc("Agus", 20)
	//if err != nil {
	//	fmt.Println("Error")
	//}
	//fmt.Println(result)

	result := print("Agus", "Budi", "Caca")
	fmt.Println(result)

	//numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//result := sum(numberList...)
	//fmt.Println(result)

	//profile("Agus", "Pasta", "Mie Goreng", "Bihun")
	fmt.Println("masuk")
}

func profile(name string, favFoods ...string) {
	var result string
	result = "My name is " + name + ", my favorite foods are "
	for _, hobby := range favFoods {
		result += hobby + ", "
	}
	println(result)
}

func sum(numberList ...int) int {
	total := 0
	for _, number := range numberList {
		total += number
	}
	return total
}

//func namaFunc(name string, age int8) (string, error) {
//
//	result := fmt.Sprintf("Hello %s, my Age is %v", name, age)
//	return result, nil
//}

func print(names ...string) []map[string]string {
	var result []map[string]string
	for i, name := range names {
		key := fmt.Sprintf("Student %d", i+1)
		tmp := map[string]string{
			key: name,
		}
		result = append(result, tmp)
	}
	return result
}
