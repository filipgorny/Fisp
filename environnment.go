package main

type wxlExecution struct {
	parameters List
}

// type wxlMethodResult = LToken

// type wxlMethod interface {
// 	Execute(wxlExecution) LToken
// }

// type wxlDefinitions map[LToken]wxlMethod

// type wxlEnvironment struct {
// 	version     string
// 	definitions wxlDefinitions
// }

// func createEnvironment() wxlEnvironment {
// 	var definitions wxlDefinitions

// 	definitions[LToken{LT_SYMBOL, "+", 0}] = wxlAdd{}

// 	return wxlEnvironment{"0.0.1", definitions}
// }
