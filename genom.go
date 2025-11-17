package main

import (
	"fmt"
)

func genome() {
	var str1, str2 string
	genom1 := NewStringSet()
	genom2 := NewStringSet()

	fmt.Println("Введите геномы")
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	for i := 0; i < len(str1)-1; i++ {
		addElem := string(str1[i]) + string(str1[i+1])
		genom1.Add(addElem)
	}

	for i := 0; i < len(str2)-1; i++ {
		addElem := string(str2[i]) + string(str2[i+1])
		genom2.Add(addElem)
	}

	intersectionGenom := genom1.IntersectionWith(genom2)
	fmt.Printf("Similarity = %d\n", intersectionGenom.Size())
}
