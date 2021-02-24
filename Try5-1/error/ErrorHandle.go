package error

type Stringer interface {
	String() string
}

type I interface {}

type ErrorImpl error {
	func Error() string {
		return ""
	}
}

func ToStringer(v I) (Stringer, error) {

}
