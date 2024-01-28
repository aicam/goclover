/*
Warning: to test functions, you need to specify all the initial variables
*/

package goclover

import (
	"log"
	"testing"
)

var CLIENT_SECRET = ""
var CLIENT_ID = ""
var CODE = ""          // changes every minute
var ACCESS_SECRET = "" // changes every auth
var MID = ""

var SAMPLE_ITEM_ID = "1ZCQ912HY9M90"

func TestService_GetAccessToken(t *testing.T) {
	s := InitApp(CLIENT_SECRET)
	secret, err := s.GetAccessToken(CLIENT_ID, CODE)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("Access token is ", secret)
}

func TestSession_GetListOfItems(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	items, err := session.GetListOfItems()

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("first item ", items.Elements[0])
}

func TestSession_GetItemCategories(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	items, err := session.GetItemCategories(SAMPLE_ITEM_ID)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("first category ", items.Elements[0])
}
