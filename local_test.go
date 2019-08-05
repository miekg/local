package local

import (
	"context"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
)

var testcases = []struct {
	question string
	qtype    uint16
	rcode    int
}{
	{"localhost.", dns.TypeA, dns.RcodeSuccess},
	{"localhost.", dns.TypeMX, dns.RcodeSuccess},
	{"a.localhost.", dns.TypeA, dns.RcodeNameError},
	{"1.0.0.127.in-addr.arpa.", dns.TypePTR, dns.RcodeSuccess},
}

func TestAny(t *testing.T) {
	req := new(dns.Msg)
	l := &Local{}

	for _, tc := range testcases {
		req.SetQuestion(tc.question, tc.qtype)
		rec := dnstest.NewRecorder(&test.ResponseWriter{})
		_, err := l.ServeDNS(context.TODO(), rec, req)

		if err != nil {
			t.Errorf("Expected no error, but got %q", err)
			continue
		}
		if rec.Msg.Rcode != tc.rcode {
			t.Errorf("Expected rcode %d, got %d", tc.rcode, rec.Msg.Rcode)
		}
	}
}
