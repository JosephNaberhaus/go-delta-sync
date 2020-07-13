# Agnostic
 - **What?** A code generation library that can target different languages.
 - **Why?** For higher-level code generation tools.
 - **How?** Code generation is abstracted behind a set of Go interfaces that support a limited subset of language features.
 ## Support
 #### Languages
  - Go
  - TypeScript
 #### Language Features
  - Models
    - Supports  a subset of Go types that translate well to other languages
  - Methods
    - Variable assignment
        - Temporary variables
        - Creating maps/arrays
        - Assigning/adding/removing to maps and arrays
        - Child properties of a model
    - Basic control flow
        - If statements
        - If/else statements
        - For each loops
        
## Documentation
### Getting Started
TODO: Add installation guide when this is actually released
#### Creating an Implementation
To get an instance of an implementation supply the language name along with language specific arguments.
```go
impl, err := targets.CreateImplementation(<language name>, <args>)
```
#### Example
This example creates a model with a single string field, and a method that will set the value of that field.
```go
package main

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets"
)

func main() {
    impl, err := targets.CreateImplementation("go", map[string]string {"package": "test"})
    if err != nil {
        return err
    }
    
    impl.Model("TestModel", agnostic.Field{Name: "value", Type: types.BaseString})
    	
    body := impl.Method(
    	"TestModel", 
    	"SetValue",
    	agnostic.Field{Name: "newValue", Type: types.BaseString},
    )
    body.Assign(value.NewOwnField(value.NewId("value")), value.NewId("newValue"))

    impl.Write("test")
}
```
Running this example will produce file called "test.go" which will contain the following code.
```go
package test

type TestModel struct {
	value string
}

func (t *TestModel) SetValue(newValue string) {
	t.value = newValue
}
```
To export to another language, you can just change the values passed into the `CreateImplementation` function. For instance, changing its parameters to `"typescript", map[string]string {}` will produce the following "test.ts" file.
```typescript
export class TestModel{
	value: string;
	public SetValue(newValue: string) {
		this.value = newValue;
	}
}
```