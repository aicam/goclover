/*
Warning: to test functions, you need to specify all the initial variables
*/

package goclover

import (
	"github.com/aicam/goclover/structures"
	"log"
	"testing"
)

var CLIENT_SECRET = "6e0849b6-8530-dddf-06a6-c61e889d236d"
var CLIENT_ID = "PSV9BBR6N34X2"
var CODE = "e51a1825-61c8-6bc2-ef4e-8be058ad8ed2"          // changes every minute
var ACCESS_SECRET = "559216cc-201c-1032-c16b-7d9520eee2b0" // changes every auth
var MID = "Q2SW4KM0D6851"

var SAMPLE_ITEM_ID = "1ZCQ912HY9M90"
var SAMPLE_ATOMIC_ORDER = structures.CreateAtomicOrderReq{OrderCart: struct {
	OrderType struct {
		Id string `json:"id"`
	} `json:"orderType"`
	GroupLineItems string                `json:"groupLineItems"`
	LineItems      []structures.LineItem `json:"lineItems"`
}(struct {
	OrderType struct {
		Id string `json:"id"`
	}
	GroupLineItems string
	LineItems      []structures.LineItem
}{OrderType: struct {
	Id string `json:"id"`
}(struct{ Id string }{Id: "GVEJ3ZXWTWFAA"}), GroupLineItems: "false", LineItems: []structures.LineItem{
	{Item: struct {
		Id string `json:"id"`
	}(struct{ Id string }{Id: "1ZCQ912HY9M90"})},
}})}
var SAMPLE_ORDER_ID = "SP4ZQ47CTRHBR"
var SAMPLE_CARD_TOKEN = "clv_1TSTS1Tavwzh9ihRkVBTQtox"
var SAMPLE_MODIFIER_GROUP = "C0F3B0Z07G5A4"
var SAMPLE_MODIFIER_ID = "3XC334KW1VJM6"

func TestService_GetAccessToken(t *testing.T) {
	s := InitApp(CLIENT_SECRET)
	secret, err := s.GetAccessToken(CLIENT_ID, CODE)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("Access token is ", secret)
}

func TestInventorySession_GetListOfItems(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	items, err := session.GetListOfItems()

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("first item ", items.Elements[0])
}

func TestInventorySession_GetItemCategories(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	items, err := session.GetItemCategories(SAMPLE_ITEM_ID)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("first category ", items.Elements[0])
}

func TestInventorySession_CreateAtomicOrder(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	order, err := session.CreateAtomicOrder(SAMPLE_ATOMIC_ORDER)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("order id ", order.Id)
}

func TestInventorySession_GetOrderInfo(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	order, err := session.GetOrderInfo(SAMPLE_ORDER_ID)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("order type id ", order.OrderType.Id)
}

func TestInventorySession_GetOrderLineItems(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	order, err := session.GetOrderLineItems(SAMPLE_ORDER_ID)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("order first line item id ", order.Elements[0].Id)
}

func TestInventorySession_ChargeCard(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	order, err := session.ChargeCard(200, "USD", SAMPLE_CARD_TOKEN)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("charge id ", order.Id)
}

func TestInventorySession_PayOrder(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	order, err := session.PayOrder(structures.PayOrderReq{
		Source: SAMPLE_CARD_TOKEN,
		Amount: 200,
	}, SAMPLE_ORDER_ID)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("order id ", order.Id)
}

func TestInventorySession_GetAllItemModifierGroups(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	modifierGroups, err := session.GetAllItemModifierGroups()

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("modifierGroups id ", modifierGroups.Elements[0])
}

func TestInventorySession_GetAllItemsOfModifierGroup(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	itemsModifierGroup, err := session.GetAllItemsOfModifierGroup(SAMPLE_MODIFIER_GROUP)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("itemsModifierGroup id ", itemsModifierGroup.Elements[0])
}

func TestInventorySession_GetModifier(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	modifier, err := session.GetModifier(SAMPLE_MODIFIER_GROUP, SAMPLE_MODIFIER_ID)

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("modifier id ", modifier)
}

func TestInventorySession_GetOrderTypes(t *testing.T) {
	s := InitApp(CLIENT_SECRET)

	session := InitSession(*s, ACCESS_SECRET, MID)
	modifier, err := session.GetOrderTypes()

	if err != nil {
		t.Fatal("Error raised: ", err)
		return
	}

	log.Println("first order type id ", modifier)
}
