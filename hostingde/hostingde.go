// Package hostingde adapts the lego hosting.de DNS
// provider for Caddy. Importing this package plugs it in.
package cloudflare

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/hostingde"
)

func init() {
	caddytls.RegisterDNSProvider("hostingde", NewDNSProvider)
}

// NewDNSProvider returns a new hosting.de DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Email address
//         credentials[1] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return hostingde.NewDNSProvider()
	case 2:
		config := hostingde.NewDefaultConfig()
		config.APIKey = credentials[0]
		config.ZoneName = credentials[1]
		return hostingde.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
