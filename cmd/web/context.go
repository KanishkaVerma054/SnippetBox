package main

/*
	// 12.2 Request context for authentication/authorization

	// define a custom contextKey type and an isAuthenticatedContextKey variable, so that we have a unique key
	// we can use to store and retrieve the authentication status from a request context (without the risk of naming collisions).
*/
type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")