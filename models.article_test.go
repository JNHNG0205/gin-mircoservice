package main

import "testing"

// test the GetAllArticles function
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check if the length of the list is the same
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check if the element is identical
	for i, v := range alist {
		if v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title ||
			v.Content != articleList[i].Content {
			t.Fail()
			break
		}

	}
}
