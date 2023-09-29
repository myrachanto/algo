package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetUserById(t *testing.T) {
	server := New()
	test := httptest.NewServer(http.HandlerFunc(server.HandleGetUserById))
	capacity := 1000

	for i := 1; i <= capacity; i++ {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", test.URL, id)
		resp, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}
		user := &User{}
		if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
			t.Error(err)
		}
		fmt.Println(user)
	}
	fmt.Println("database hit:", server.dbhit)

}
