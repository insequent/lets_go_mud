package main

type Alias struct {
	Output string
}

type Trigger struct {
	Match  string
	Output string
}

type Class struct {
	Aliases  map[string]*Alias
	Triggers map[string]*Trigger
}

type Configuration struct {
	// General Settings
	HostPort      string
	RepeatOnEnter bool

	// Classes
	Active   map[string]struct{}
	Profiles map[string]*Class
}

func ParseConfig(data, file string) error {
	return nil
}
