package main

import "fmt"

func main() {
	// websites := []string{"https://google.com", "https://aws.com"}

	websites := map[string]string{
		"google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	// fmt.Println(websites["Amazon Web Services"])

	websites["udemy"] = "https://udemy.com"
	// fmt.Println(websites)

	delete(websites, "google")
	// fmt.Println(websites)

	userNames := make([]string, 2, 5)
	userNames[0] = "julie"

	userNames = append(userNames, "Tom")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Chad")
	userNames = append(userNames, "Manuel")

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range websites {
		fmt.Println(key)
		fmt.Println(value)
	}
}
