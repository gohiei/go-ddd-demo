package dddcore

// Input represents the input parameter for a use case or operation.
type Input interface{}

// Output represents the output result of a use case or operation.
type Output interface{}

// UseCase represents a use case or operation in the application.
// It is parameterized with types I for the input and O for the output.
type UseCase[I Input, O Output] interface {
	// Execute executes the use case with the given input and returns the output and an error if applicable.
	Execute(*I) (*O, error)
}
