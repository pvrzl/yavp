## YAVP (Yet Another Validation Package)

the name says it all, inspired by https://www.npmjs.com/package/validator

#### Features
- validate without struct tag
- able to merge error into one json
- customizable
- extremely easy to create and use custom validation rules

#### Example
Simple example
```go
yavp.RequiredString.Validate("hello")
```

Struct Validation(with json output)
```go
import "github.com/pvrzl/yavp"

type Form struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (f *Form) Validate() error {
	return yavp.MergeErrors(
		yavp.ValidateString(
			"name",
			f.Name,
			yavp.RequiredString,
		),
		yavp.ValidateString(
			"address",
			f.Address,
			yavp.RequiredString,
		),
	)
}
```

the output of above code will be
```JSON
{
    "name": "is required",
    "address": "is required"
}

```

Change error message
```go
yavp.RequiredString.WithError(errors.New("is required"))
yavp.RequiredString.WithErrorMessage("is required")
```

Custom Validation 
```go
func IsFoo() StringValidator {
	return NewStringValidator(func(s string, err error) error {
		if s == "foo" {
			return nil
		}
		return err
	}, errors.New("is error"))
}
```

you can also create your own validator, please see the source code for reference.


