// Package lightsail adapts the lego AWS Lightsail DNS
// provider for Caddy. Importing this package plugs it in.
package lightsail

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/lightsail"
)

func init() {
	caddytls.RegisterDNSProvider("lightsail", NewDNSProvider)
}

// NewDNSProvider returns a new AWS Lightsail DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return lightsail.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
