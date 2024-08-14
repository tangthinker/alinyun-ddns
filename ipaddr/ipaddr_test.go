package ipaddr

import "testing"

func TestIpaddr(t *testing.T) {

	ipv4, err := GetIPv4Addr()

	if err != nil {
		t.Fatal(err)
	}

	t.Log("IPv4:", ipv4)

	ipv6, err := GetIPv6Addr()

	if err != nil {
		t.Fatal(err)
	}

	t.Log("IPv6", ipv6)

}
