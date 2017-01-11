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
	slug, _ := idToSlug(125)
	assertEqual(t, slug, "cb", "Wrong slug")

	slug, _ = idToSlug(19158)
	assertEqual(t, slug, "e9a", "Wrong slug")
}

func TestIdToSlug_InvalidInput(t *testing.T) {
	slug, err := idToSlug(0)
	assertEqual(t, slug, "", "slug should be empty string on error")
	assertEqual(t, err.Error(), "id must be positive integer", "Wrong error message")
}

func TestSlugToId(t *testing.T) {
	id, _ := slugToId("cb")
	assertEqual(t, id, 125, "Wrong id")

	id, _ = slugToId("e9a")
	assertEqual(t, id, 19158, "Wrong id")
}

func TestSlugToId_InvalidInput(t *testing.T) {
	id, err := slugToId("男裝")
	assertEqual(t, id, 0, "id should be 0 on error")
	assertEqual(t, err.Error(), "Invalid character found in slug", "Wrong error message")
}

