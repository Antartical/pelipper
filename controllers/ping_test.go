package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	assert := require.New(t)
	var response gin.H
	expected_response := gin.H{
		"data": "pong",
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	Ping(c)

	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	if err != nil {
		assert.Fail("Payload does not match with the expected one")
	}

	assert.Equal(recorder.Result().StatusCode, 200)
	assert.Equal(expected_response, response)
}
