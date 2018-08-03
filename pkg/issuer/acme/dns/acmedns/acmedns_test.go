package acmedns

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	acmednsLiveTest     bool
	acmednsHost         string
	acmednsAccountsJson []byte
	acmednsDomain       string
)

func init() {
	acmednsHost = os.Getenv("ACME_DNS_HOST")
	acmednsAccountsJson = []byte(os.Getenv("ACME_DNS_ACCOUNTS_JSON"))
	acmednsDomain = os.Getenv("ACME_DNS_DOMAIN")
	if len(acmednsHost) > 0 && len(acmednsAccountsJson) > 0 {
		acmednsLiveTest = true
	}
}

func TestLiveAcmeDnsPresent(t *testing.T) {
	if !acmednsLiveTest {
		t.Skip("skipping live test")
	}
	provider, err := NewDNSProviderHostBytes(acmednsHost, acmednsAccountsJson)
	assert.NoError(t, err)

	err = provider.Present(acmednsDomain, "", "123d==")
	assert.NoError(t, err)
}
