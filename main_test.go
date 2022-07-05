package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllHandler(t *testing.T){
	ts := httptest.NewServer(Server())
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/getAll", ts.URL))

	assert.NoError(t,err,"An error occurred while sending the request.")
	assert.Equal(t,200,resp.StatusCode,"Status code should be 200")
	body,_ := ioutil.ReadAll(resp.Body)
	result := string(body)
	assert.Greater(t,len(result),5,"Response body should not be empty")
	assert.Equal(t,[]string{"application/json; charset=utf-8"},resp.Header["Content-Type"],"Response should be json")
}

func TestConvertExchangeRatesHandler(t *testing.T){

	ts := httptest.NewServer(Server())
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/Try/Usd/1/convertRate", ts.URL))

	assert.NoError(t,err,"An error occurred while sending the request.")
	assert.Equal(t,200,resp.StatusCode,"Status code should be 200")
	body,_ := ioutil.ReadAll(resp.Body)
	result := string(body)
	assert.Greater(t,len(result),5,"Response body should not be empty")
	assert.Equal(t,[]string{"application/json; charset=utf-8"},resp.Header["Content-Type"],"Response should be json")
}


func TestPullExchangeRateHandler(t *testing.T){

	ts := httptest.NewServer(Server())
    // Shut down the server and block until all requests have gone through
    defer ts.Close()

	resp, err := http.Post(fmt.Sprintf("%s/pullRate", ts.URL),"application/json",nil)

	assert.NoError(t,err,"An error occurred while sending the request.")
	assert.Equal(t,200,resp.StatusCode,"Status code should be 200")
	body,_ := ioutil.ReadAll(resp.Body)
	result := string(body)
	assert.Greater(t,len(result),5,"Response body should not be empty")
	assert.Equal(t,[]string{"application/json"},resp.Header["Content-Type"],"Response should be json")
}