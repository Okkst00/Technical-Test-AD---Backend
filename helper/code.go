package helper

import (
	"fmt"
	"math/rand"
	"time"
)


func GenerateOrderCode() string {

	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(999999)


	return fmt.Sprintf(
		"ORD-%d-%06d",
		time.Now().Year(),
		number,
	)

}