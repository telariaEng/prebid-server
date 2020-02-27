package telaria

import (
	"github.com/prebid/prebid-server/adapters/adapterstest"
	"testing"
)

/**
 * Verify adapter names are setup correctly.
 */
func TestTelariaAdapterNames(t *testing.T) {
	adapter := NewTelariaBidder("https://{{.Host}}")
	adapterstest.VerifyStringValue(adapter.Name(), "telaria", t)
}

/**
 * Verify adapter SkipNoCookie is correct.
 */
func TestTelariaAdapterSkipNoCookiesFlag(t *testing.T) {
	adapter := NewTelariaBidder("https://{{.Host}}")
	adapterstest.VerifyBoolValue(adapter.SkipNoCookies(), false, t)
}

/**
 * Verify bidder has the proper URL
 */
func TestTelariaAdapterEndpoint(t *testing.T) {
	adapter := NewTelariaBidder("https://{{.Host}}")
	adapterstest.VerifyStringValue(adapter.URI, "https://ads.tremorhub.com/ad/rtb/prebid", t)
}

func testJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "telariatest", new(TelariaAdapter))
}