package e2e

import (
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kot13/vertigo/client"
	"github.com/kot13/vertigo/client/operations"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	transport := httptransport.New("localhost:"+os.Getenv("PORT"), "/api", nil)

	c := client.New(transport, strfmt.Default)

	response, err := c.Operations.GetHealthCheck(operations.NewGetHealthCheckParams())

	assert.Nil(t, err)
	if t.Failed() {
		return
	}

	assert.NotEqual(t, "", response.Payload)
	if t.Failed() {
		return
	}
}
