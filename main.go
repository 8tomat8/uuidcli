package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no value.\nuse 'b2d7a013-86df-4aa7-8245-8f3715c87ae2' or 'KI+VUsKMSqGvNJW6jORsWg=='")
	}

	val := os.Args[1]
	id, err := uuid.Parse(val)
	if err == nil {
		fmt.Println(base64.StdEncoding.EncodeToString(id[:]))
		return
	}

	fmt.Println(decodeUUID(val))
}

func decodeUUID(raw string) string {
	rawB, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode base64"))
	}

	id, err := uuid.FromBytes(rawB)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse uuid"))
	}

	return id.String()
}

func encodeUUID(raw string) string {
	id, err := uuid.Parse(raw)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse uuid"))
	}

	return base64.StdEncoding.EncodeToString(id[:])
}
