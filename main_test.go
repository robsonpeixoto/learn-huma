package main

import (
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	_, api := humatest.New(t)
	addRoutes(api)

	resp := api.Get("/greeting/world")
	assert.Contains(t, resp.Body.String(), "Hello, world")
}

func TestPutReview(t *testing.T) {
	_, api := humatest.New(t)
	addRoutes(api)

	resp := api.Post("/reviews", map[string]any{
		"author": "daniel",
		"rating": 5,
	})

	assert.Equal(t, 201, resp.Code)
}

func TestPutReviewError(t *testing.T) {
	_, api := humatest.New(t)
	addRoutes(api)

	resp := api.Post("/reviews", map[string]any{
		"rating": 10,
	})

	assert.Equal(t, 422, resp.Code)
}
