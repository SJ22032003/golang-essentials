package main

import "fmt"

// func main() {
// 	websites := map[string]string{
// 		"google": "http://google.com", "facebook": "http://facebook.com", "twitter": "http://twitter.com",
// 	}

// 	allowedWebsites := []string{"google", "facebook", "twitter", "linkedin"}

// 	var i = 0
// 	for i < len(allowedWebsites) {
// 		website, ok := websites[allowedWebsites[i]]
// 		if ok {
// 			fmt.Println(website)
// 		} else {
// 			fmt.Println("Website not found", allowedWebsites[i])
// 			// delete(websites, allowedWebsites[i]) // delete the website from the map
// 		}
// 		i++
// 	}

// }

type AllowedWebsites map[string]interface{} // Add closing curly brace here | any

type AllowedUsage map[string]AllowedWebsites

func main() {
	allowedUsage := AllowedUsage{
		"google": {
			"url":   []string{"http://google.com", "http://google.co.in"},
			"usage": "search engine",
		},
	}

	allowedWebsites := []string{"google", "facebook", "twitter", "linkedin"}

	allowedUsage["google"]["url"] = append(allowedUsage["google"]["url"].([]string), "http://google.co.uk")

	var i = 0
	for i < len(allowedWebsites) {
		website, ok := allowedUsage[allowedWebsites[i]]
		if ok {
			fmt.Println(website["url"], website["usage"])
		} else {
			fmt.Println("Website not found", allowedWebsites[i])
			// delete(websites, allowedWebsites[i]) // delete the website from the map
		}
		i++
	}

}
