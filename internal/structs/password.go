package structs

// PasswordParamethers holds password parameters
type PasswordParamethers struct {
	RequireUpper  bool
	RequireLower  bool
	RequireNumber bool
	RequireSymbol bool
	RequireLength int
}

var PasswordConfig PasswordParamethers
