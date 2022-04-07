package test_helpers

type HttpTestDescription struct {
	Description   string // description of the test case
	Route         string // route path to test
	Method        string // http method
	ExpectedCode  int    // expected HTTP status code
	Authorization string
	Body          []byte
	ExpectedBody  interface{}
}
