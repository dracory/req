package req

// Package req provides functions for working with HTTP requests
//
// The package is imported like this:
//
//	import "github.com/gouniverse/req"
//
// # Example
//
// To get the value of a GET parameter:
//
//	value := req.Value(r, "key")
//
// To get the value of a POST parameter:
//
//	value := req.Value(r, "key")
//
// To get the value of a GET or POST parameter:
//
//	value := req.Value(r, "key")
//
// To check if a GET parameter exists:
//
//	if req.HasGet(r, "key") {
//	    // key exists
//	}
//
// To check if a POST parameter exists:
//
//	if req.HasPost(r, "key") {
//	    // key exists
//	}
//
// To check if a GET or POST parameter exists:
//
//	if req.Has(r, "key") {
//	    // key exists
//	}
//
// To get all GET parameters:
//
//	all := req.AllGet(r)
//
// To get all POST parameters:
//
//	all := req.AllPost(r)
//
// To get all GET and POST parameters:
//
//	all := req.All(r)
//
// To get a map of GET parameters:
//
//	map := req.Map(r, "key")
//
// To get a map of POST parameters:
//
//	map := req.Map(r, "key")
//
// To get a map of GET or POST parameters:
//
//	map := req.Map(r, "key")
