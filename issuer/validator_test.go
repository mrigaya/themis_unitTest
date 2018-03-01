package issuer

import (
	"testing"
)

func TestValidateMACAddress(t *testing.T) {
	
	// Valid MacAddresses 
	assert.Equal(t, true, ValidateMACAddress("14cfe3213211"), "Mac validation SUCCESS is expected")
	assert.Equal(t, true, ValidateMACAddress("14:32:AB:13:c2:11"), "Mac validation SUCCESS is expected")
	assert.Equal(t, true, ValidateMACAddress("14.cf.e3.21.32.11"), "Mac validation SUCCESS is expected")
	assert.Equal(t, true, ValidateMACAddress("14-32-AB-13-C2-11"), "Mac validation SUCCESS is expected")
	assert.Equal(t, true, ValidateMACAddress("14,cf,e3,21,32,11"), "Mac validation SUCCESS is expected")
	assert.Equal(t, true, ValidateMACAddress("1432AB13C211"), "Mac validation SUCCESS is expected")
	assert.Equal(t, true, ValidateMACAddress("14-32,AB-13:C2.11"), "Mac validation SUCCESS is expected")
	
	// Invalid MacAddresses
		
	// Contains special character "{" which is not valid, expected delimiters are ":-.,"
	assert.Equal(t, false, ValidateMACAddress("14:cf{e3:21:32:11"), "Mac validation failed is expected")
	assert.Equal(t, false, ValidateMACAddress("14cfe32*3211"), "Mac validation failed is expected")
	assert.Equal(t, false, ValidateMACAddress("12345678911&"), "Mac validation failed is expected")
	
	// Invalid hexadecimal value
	assert.Equal(t, false, ValidateMACAddress("14xyz3213211"), "Mac validation failed is expected")
	// MacAddress length should be equal to 12
	assert.Equal(t, false, ValidateMACAddress("14cfe123456789"), "Mac validation failed is expected")
	assert.Equal(t, false, ValidateMACAddress("14cfe323211"), "Mac validation failed is expected")
}
