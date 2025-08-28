package req

// Package req provides functions for working with HTTP requests
//
// The package is imported like this:
//
//	import "github.com/dracory/req"
//
// # Example
//
// To get the value of a GET parameter:
//
//	value := req.GetString(r, "key")
//
// To get the value of a POST parameter:
//
//	value := req.GetString(r, "key")
//
// To get the value of a GET or POST parameter:
//
//	value := req.GetString(r, "key")
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
//	all := req.GetAllGet(r)
//
// To get all POST parameters:
//
//	all := req.GetAllPost(r)
//
// To get all GET and POST parameters:
//
//	all := req.GetAll(r)
//
// To get a map of GET parameters:
//
//	m := req.GetMap(r, "key")
//
// To get a map of POST parameters:
//
//	m := req.GetMap(r, "key")
//
// To get a map of GET or POST parameters:
//
//	m := req.GetMap(r, "key")
