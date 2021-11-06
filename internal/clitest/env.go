// this package is for cli test
// Do not use this package in production code.
package clitest

import "os"

//
// usage
// ```go
// func TestAPIToken(t *testing.T) {
//     reset = setTestEnv("Token","……")
//     defer reset()
//
//     // some test
// }
// ```
//
// set one time environment variable
func SetTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}
