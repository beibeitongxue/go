package pay_package

type Package struct {
	Name  string
	Money string
	Data  string
	Time  string
}
type ExpensivePackage struct {
	Name  string
	Money string
	Data  string
	Time  string
}
type CustomPackage struct {
	Name  string
	Money string
	Data  string
	Time  string
	Nodes string
}
type ExtraPackage struct {
	Name string
	Data string
}

func NewPackage(name, date, time, money string) *Package {
	return &Package{
		Name:  name,
		Data:  date,
		Time:  time,
		Money: money,
	}
}
func NewExpensivePackage(name, date, time, money string) *ExpensivePackage {
	return &ExpensivePackage{
		Name:  name,
		Data:  date,
		Time:  time,
		Money: money,
	}
}
func NewCustomPackage(name, date, time, money, nodes string) *CustomPackage {
	return &CustomPackage{
		Name:  name,
		Data:  date,
		Time:  time,
		Money: money,
		Nodes: nodes,
	}
}
func NewExtraPackage(name, date string) *ExtraPackage {
	return &ExtraPackage{
		Name: name,
		Data: date,
	}
}

func NormalA() *Package {
	return &Package{
		Name:  "Normal A",
		Data:  "100G",
		Time:  "31",
		Money: "30",
	}
}
func NormalB() *Package {
	return &Package{
		Name:  "Normal B",
		Data:  "200G",
		Time:  "31",
		Money: "45",
	}
}
func NormalC() *Package {
	return &Package{
		Name:  "Normal C",
		Data:  "100G",
		Time:  "90",
		Money: "88",
	}
}
func NormalD() *Package {
	return &Package{
		Name:  "Normal D",
		Data:  "200G",
		Time:  "90",
		Money: "128",
	}
}

func ExpensiveA() *ExpensivePackage {
	return &ExpensivePackage{
		Name:  "Expensive A",
		Data:  "200G",
		Time:  "365",
		Money: "238",
	}
}
func ExpensiveB() *ExpensivePackage {
	return &ExpensivePackage{
		Name:  "Expensive B",
		Data:  "300G",
		Time:  "365",
		Money: "298",
	}
}
func ExpensiveC() *ExpensivePackage {
	return &ExpensivePackage{
		Name:  "Expensive C",
		Data:  "500",
		Time:  "365",
		Money: "488",
	}
}

func CustomA() *CustomPackage {
	return &CustomPackage{
		Name:  "CustomPackage A",
		Data:  "1TB",
		Time:  "90",
		Money: "480",
		Nodes: "50",
	}
}
func CustomB() *CustomPackage {
	return &CustomPackage{
		Name:  "CustomPackage B",
		Data:  "2TB",
		Time:  "90",
		Money: "680",
		Nodes: "100",
	}
}
func CustomC() *CustomPackage {
	return &CustomPackage{
		Name:  "CustomPackage C",
		Data:  "1TB",
		Time:  "364",
		Money: "1680",
		Nodes: "50",
	}
}
func CustomD() *CustomPackage {
	return &CustomPackage{
		Name:  "CustomPackage D",
		Data:  "2TB",
		Time:  "365",
		Money: "2560",
		Nodes: "100",
	}
}

func ExtraA() *ExtraPackage {
	return &ExtraPackage{
		Name: "Extra A",
		Data: "100G",
	}
}
func ExtraB() *ExtraPackage {
	return &ExtraPackage{
		Name: "Extra B",
		Data: "200G",
	}
}
func ExtraC() *ExtraPackage {
	return &ExtraPackage{
		Name: "Extra C",
		Data: "500G",
	}
}
