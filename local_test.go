package local

import (
	"context"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
)

func TestAny(t *testing.T) {
	req := new(dns.Msg)
	req.SetQuestion("localhost.", dns.TypeSOA)
	l := &Local{}

	rec := dnstest.NewRecorder(&test.ResponseWriter{})
	_, err := a.ServeDNS(context.TODO(), rec, req)

	if err != nil {
		t.Errorf("Expected no error, but got %q", err)
	}

	println(rec.String())
}
