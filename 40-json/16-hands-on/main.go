package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type codes struct {
	Code int
	Description string
}

func main() {
	var data []codes
	rcvd := `
	[{"Code":200,"Description":"StatusOK"},
	{"Code":301,"Description":"StatusMovedPermanently"},
	{"Code":302,"Description":"StatusFound"},
	{"Code":303,"Description":"StatusSeeOther"},
	{"Code":307,"Description":"StatusTemporaryRedirect"},
	{"Code":400,"Description":"StatusBadRequest"},
	{"Code":401,"Description":"StatusUnauthorized"},
	{"Code":402,"Description":"StatusPaymentRequired"},
	{"Code":403,"Description":"StatusForbidden"},
	{"Code":404,"Description":"StatusNotFound"},
	{"Code":405,"Description":"StatusMethodNotAllowed"},
	{"Code":418,"Description":"StatusTeapot"},
	{"Code":500,"Description":"StatusInternalServerError"}]
	`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Description)
	}
}
