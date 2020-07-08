package test

// A test file in an arbitrary programming language
type Implementation interface {
	Write(fileName string)
	Test(testCase Case)
}
