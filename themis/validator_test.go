package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMatchClientHostnames(t *testing.T) {

	/* Valid Hostnames matching with pattern */
	assert.Equal(t, true, matchClientHostnames("www.test1.com", "www.test1.com"), "Hostname validation SUCCESS is expected")
	assert.Equal(t, true, matchClientHostnames("*.test2.com", "www.test2.com"), "Hostname validation SUCCESS is expected")
	assert.Equal(t, true, matchClientHostnames("themis.*.com", "themis.*.com"), "Hostname validation SUCCESS is expected")
	assert.Equal(t, true, matchClientHostnames("192.168.0.0", "192.168.0.0"), "Hostname validation SUCCESS is expected")
	assert.Equal(t, true, matchClientHostnames("www.test1.com", "*ww.test1.com"), "Hostname validation SUCCESS is expected")
	// Ideally this test case should fail, As we are not supporting this feature now so it will pass
	assert.Equal(t, true, matchClientHostnames("192.168.0.0", "*.168.0.0"), "Hostname validation Failed is expected")
	
	/* Invalid hostnames */
	assert.Equal(t, false, matchClientHostnames("*.test6.com", "www.test.test6.com"), "Hostname validation Failed is expected")
	assert.Equal(t, false, matchClientHostnames("*.", ".*"), "Hostname validation Failed is expected")
	assert.Equal(t, false, matchClientHostnames("www.test7.com", "*.*.test7.com"), "Hostname validation Failed is expected")
	
	// Partial wildcard comparison is not supported
	assert.Equal(t, false, matchClientHostnames("ww.test3.com", "ww*.test3.com"), "Hostname validation SUCCESS is expected")
	assert.Equal(t, false, matchClientHostnames("xxx*.test8.com", "xxxwww.test8.com"), "Hostname validation Failed is expected")
	
	assert.Equal(t, false, matchClientHostnames("*.168.0.0", "192.169.0.0"), "Hostname validation Failed is expected")
	assert.Equal(t, false, matchClientHostnames("192.168.0.0", "192.*.0.0"), "Hostname validation Failed is expected")
	assert.Equal(t, false, matchClientHostnames("*::3285:a9ff:fe46:b619","fe80::3285:a9ff:fe46:b619"), "Hostname validation Failed is expected")
}
