package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry"

	encodeStd := "ABCDEFGHIJ*LMNOPQRS#UVWXY+abcdefghijkl=nopqrstuvwxyz0123456789-("

	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	std64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s)
	fmt.Println(s64)
	fmt.Println(std64)

	bs, err := base64.StdEncoding.DecodeString(std64)
	if err != nil {
		log.Fatalln("Cannot decode", err)
	}
	fmt.Println(string(bs))
}
