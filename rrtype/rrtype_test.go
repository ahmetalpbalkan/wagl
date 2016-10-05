package rrtype

import (
	"testing"

	"github.com/miekg/dns"
)

func TestIsSupported(t *testing.T) {
	cases := []struct {
		rrType    uint16
		supported bool
	}{
		// supported
		{dns.TypeA, true},
		{dns.TypeSRV, true},
		{dns.TypeAAAA, true},

		// some others
		{dns.TypeCNAME, false},
		{dns.TypeNS, false},
		{dns.TypeMX, false},
	}
	for _, c := range cases {
		out := IsSupported(c.rrType)
		if out != c.supported {
			t.Fatal("wrong value for", dns.TypeToString[c.rrType])
		}
	}
}

func TestToRR(t *testing.T) {
	ToRR(dns.TypeAAAA, "api.domain", "2001:cdba::3257:9652")
}
