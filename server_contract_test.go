package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGraphQLValidationContracts(t *testing.T) {
	server := httptest.NewServer(newHandler())
	t.Cleanup(server.Close)
	client := server.Client()
	client.Timeout = 5 * time.Second

	cases := []struct {
		name string
		body string
	}{
		{"unknown field", `{"query":"{ unknownField }"}`},
		{"invalid syntax", `{"query":"{ heroes {"}`},
		{"missing required variable", `{"query":"query($show: Boolean!) { heroes @include(if: $show) { name } }"}`},
		{"wrong variable type", `{"query":"query($show: Boolean!) { heroes @include(if: $show) { name } }", "variables":{"show":"not-a-boolean"}}`},
		{"unknown operation", `{"query":"query Heroes { heroes { name } }", "operationName":"Missing"}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.Post(server.URL+"/query", "application/json", bytes.NewBufferString(tc.body))
			require.NoError(t, err)
			defer response.Body.Close()
			require.Less(t, response.StatusCode, http.StatusInternalServerError)
			var payload struct {
				Data   json.RawMessage `json:"data"`
				Errors []struct {
					Message string `json:"message"`
				} `json:"errors"`
			}
			require.NoError(t, json.NewDecoder(response.Body).Decode(&payload))
			require.NotEmpty(t, payload.Errors)
			require.NotEmpty(t, payload.Errors[0].Message)
			require.True(t, len(payload.Data) == 0 || string(payload.Data) == "null")
		})
	}
}

func TestGraphQLBooleanVariableControlsSelection(t *testing.T) {
	server := httptest.NewServer(newHandler())
	t.Cleanup(server.Close)
	client := server.Client()
	client.Timeout = 5 * time.Second
	response, err := client.Post(server.URL+"/query", "application/json", bytes.NewBufferString(
		`{"query":"query($show: Boolean!) { heroes @include(if: $show) { name } types }", "variables":{"show":false}}`))
	require.NoError(t, err)
	defer response.Body.Close()
	require.Equal(t, http.StatusOK, response.StatusCode)
	var payload struct {
		Data   map[string]json.RawMessage `json:"data"`
		Errors []json.RawMessage          `json:"errors"`
	}
	require.NoError(t, json.NewDecoder(response.Body).Decode(&payload))
	require.Empty(t, payload.Errors)
	require.NotContains(t, payload.Data, "heroes")
	require.Contains(t, payload.Data, "types")
}
