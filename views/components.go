package views

type ColorVariant string

const (
	PrimaryColor   ColorVariant = "primary"
	SecondaryColor ColorVariant = "secondary"
	DangerColor    ColorVariant = "danger"
	SuccessColor   ColorVariant = "success"
)

func GetColorClass(color ColorVariant) string {
	switch color {
	case PrimaryColor:
		return "bg-blue-600 hover:bg-blue-700"
	case SecondaryColor:
		return "bg-gray-600 hover:bg-gray-700"
	case DangerColor:
		return "bg-red-600 hover:bg-red-700"
	case SuccessColor:
		return "bg-green-600 hover:bg-green-700"
	default:
		return "bg-gray-600 hover:bg-gray-700"
	}
}
