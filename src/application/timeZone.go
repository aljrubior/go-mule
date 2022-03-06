package application

type TimeZone int

const (
	Belize TimeZone = iota
	Cancun
	Chicago
	Guatemala
	Managua
	Matamoros
	Menominee
	Merida
	Monterrey
	Regina
	Resolute
	Tegucigalpa
	Winnipeg
)

func (timeZone TimeZone) String() string {
	return [...]string{
		"America/Belize",
		"America/Cancun",
		"America/Chicago",
		"America/Guatemala",
		"America/Managua",
		"America/Matamoros",
		"America/Menominee",
		"America/Merida",
		"America/Monterrey",
		"America/Regina",
		"America/Resolute",
		"America/Tegucigalpa",
		"America/Winnipeg"}[timeZone]

}
