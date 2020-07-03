package values

// A value within a name/id/variable
type IdStruct struct {
	name string
}

func (v IdStruct) Name() string {
	return v.name
}

func Id(name string) IdStruct {
	return IdStruct{name: name}
}
