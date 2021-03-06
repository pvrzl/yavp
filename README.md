## YAVP (Yet Another Validation Package)
[![GoDoc](https://godoc.org/github.com/pvrzl/yavp?status.png)](https://godoc.org/github.com/pvrzl/yavp)
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
		yavp.GenericValidator(
			"name", 
			yavp.RequiredString.Validate(f.Name),
			yavp.When(f.Name != "foo", errors.New("error occured")),
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
yavp.RequiredString.WithErrorMessage("is required").Validate("hello")
```

Custom Validation 
```go
func IsFoo() yavp.StringValidator {
	return yavp.NewStringValidator(func(s string, err error) error {
		if s == "foo" {
			return nil
		}
		return err
	}, errors.New("is error"))
}
```

you can also create your own validator, please see the source code for reference.


