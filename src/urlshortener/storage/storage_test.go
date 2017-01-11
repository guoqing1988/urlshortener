package storage

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestIdToSlug(t *testing.T) {
	slug := idToSlug(125)
	assertEqual(t, slug, "cb", "Wrong slug")

	slug = idToSlug(19158)
	assertEqual(t, slug, "e9a", "Wrong slug")
}

func TestSlugToId(t *testing.T) {
	id := slugToId("cb")
	assertEqual(t, id, 125, "Wrong id")

	id = slugToId("e9a")
	assertEqual(t, id, 19158, "Wrong id")
}
