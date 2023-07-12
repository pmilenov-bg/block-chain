package	main

import (
	"bufio"
	"os"
	"fmt"
	hashGenerator "hash-chain/src/services"
)

func main()  {

	reader := bufio.NewScanner(os.Stdin)
	paswordMap := make(map[string]string)
	reader.Scan()
	inp := reader.Text()
	// if err {
	// 	fmt.Println(err)
	// }
	paswordMap[inp] =  hashGenerator.HashString1(inp)
	passwords := []string{}
	passwords = append(passwords, inp)
	count := 0
	for {
		fmt.Println("password:")
		reader.Scan()
		input := reader.Text()
		// if err {
		// 	fmt.Println(err)
		// }
		// input =strings.TrimSpace(input)

		if input =="exit" {
			break
		}
		passwords = append(passwords, input)
		paswordMap[input] = hashGenerator.HashString2(input, paswordMap[passwords[count]]) 
		count++
	}

	fmt.Println(paswordMap)
	// reader.Scan()
	// line := reader.Text()
	// // if err != nil {
	// // fmt.Println("Error: %s", err)
	// // }
	// hashString := hashGenerator.HashString1(line)
	// fmt.Println(hashString)
	
	// reader.Scan()
	// line2 := reader.Text()
	// hashString2 := hashGenerator.HashString2(line2, hashString)
	// fmt.Println(hashString2)
}