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
        - Assigning/adding/removing to a map, array, 
        - Child properties of a model
    - Basic control flow
        - If statements
        - If/else statements
        - For each loops
        
## Documentation
### Getting Started
TODO: Add installation guide when this is actually released
#### Creating an Implementation
To get an instance of an implementation supply the language name along with arguments for that language. Find a list of
supported arguments in the language specific Readme.
```go
impl, err := targets.CreateImplementation(<language name>, <args>)
```
#### Example
This example creates a model with a single string field and a method that will set the value of that field.
```go
package main

import "github.com/JosephNaberhaus/go-delta-sync/agnostic"

func main() {
    impl, err := targets.CreateImplementation("go", map[string]string {"package": "test"})
    if err != nil {
        return err
    }
    
    impl.Model("TestModel", agnostic.Field{Name: "value", Type: types.BaseString})
    	
    body := implementation.Method(
    	"TestModel", 
    	"SetValue",
    	agnostic.Field{Name: "newValue", Type: types.BaseString},
    )
    body.Assign(value.NewOwnField(value.NewId("value")), value.NewId("newValue"))
}
```
The last line shows the nestable nature of the value primitive. The `NewOwnField` wrapper specifies that the sub-value is a field of the model the method belongs to. The `NewId` child then gives the name of that field. Many other value primitives, like `ArrayElement`, will wrap child-values in a similar manner. 