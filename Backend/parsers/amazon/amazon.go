package amazon

import (
	"fmt"
	"math/rand"
	"strings"

	"../../reviews"
)

func GetReviews(pattern string) []reviews.Review {
	var revs []reviews.Review

	for i := 0; i < 100; i++ {
		review := reviews.Review{
			Rating: rand.Float32() + 4,
			Origin: "amazon",
			Author: "Tester",
			Avatar: fmt.Sprintf("https://www.gravatar.com/avatar/%d?d=identicon", rand.Intn(10000)),
			Text:   fmt.Sprintf("Thats what I think about %s", strings.Replace(pattern, "+", " ", -1)),
			Tags: []reviews.Tag{
				{"keyboard", 0.3},
				{"mouse", -0.7},
			},
		}

		revs = append(revs, review)
	}

	return revs
}
