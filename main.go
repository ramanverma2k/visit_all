package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)


func main() {
	main_url := "https://www.moneycontrol.com/"
	keyword := "Gayatri"
	
	homepage := page_request(&main_url)

	re := regexp.MustCompile(`\s*(?i)https://www[.]moneycontrol[.]com(\"([^"]*\")|'[^']*'|([^'">\s]+))`)
	links_to_visit := re.FindAllString(homepage, -1)
	filtered_links := removeDuplicatesUnordered(links_to_visit)
	fmt.Println(visit_links(&keyword, filtered_links))
}

func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}

func page_request(url *string) string {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func visit_links(keyword *string,links []string) int {
	var sum int
	for _, l := range links {
		count := find_keyword(keyword, page_request(&l))
		sum = sum + count
	}
	return sum
}

func find_keyword(keyword *string, source string) int {
	re := regexp.MustCompile(*keyword)
	match := re.FindAllString(source, -1)
	if match != nil {
		fmt.Println(keyword, match)
	}
	return len(match)
}