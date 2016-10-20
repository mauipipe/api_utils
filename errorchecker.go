package api_utils

func Errorchecker(e error) {
	if e != nil {
		panic(e)
	}
}