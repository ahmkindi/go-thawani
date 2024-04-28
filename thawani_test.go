package thawani

import (
	"log"
	"net/url"
	"testing"

	"github.com/ahmkindi/go-thawani/types/mode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestMetadata struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
}

var (
	testURL = "https://uatcheckout.thawani.om"
	testAPI = "rRQ26GcsZzoEhbrP2HZvLYDbn9C9et"
	testKey = "HGvTMLDssJghr9tlN9gr4DVYt0qyBy"
	client  *ThawaniClient
)

func init() {
	host, err := url.Parse(testURL)
	if err != nil {
		log.Fatalf("failed to init thawanit: %s", err.Error())
	}
	client = NewClient(nil, host, testAPI, testKey)
}

func TestCreateSession(t *testing.T) {
	id := "1234"
	resp, err := client.CreateCustomer(CreateCustomerReq{
		ClientCustomerId: "1234",
	})
	require.NoError(t, err)
	require.True(t, resp.Success)

	metadata := TestMetadata{
		Hello: "world",
		Name:  "test",
	}

	req := CreateSessionReq{
		ClientReferenceId: id,
		CustomerId:        resp.Data.Id,
		Mode:              mode.Payment,
		Products: []Product{{
			Name:       "test",
			Quantity:   10,
			UnitAmount: 10,
		}},
		SuccessUrl: "https://test.com",
		CancelUrl:  "https://test.com",
		Metadata:   metadata,
	}
	sessionResp, _, err := client.CreateSession(req)
	require.NoError(t, err)
	assert.Equal(t, 2004, sessionResp.Code)
	assert.Equal(t, metadata.Hello, sessionResp.Data.Metadata["hello"])
	assert.Equal(t, metadata.Name, sessionResp.Data.Metadata["name"])
}
