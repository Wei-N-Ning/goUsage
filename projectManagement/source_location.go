package project_management

import "runtime"


/*
Functionally equivalent to Python function that returns the value of the __file__ variable
 */
func SourceLocation() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}
