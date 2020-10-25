package math

var name = "唯一"

func Add(a int64, b int64) int64 {
	return a * b
}

func GetName() (int64, int64) {
	return cc(2, 3)
}

func cc(a int64, b int64) (int64, int64) {
	return a, b
}

func GetN() string {
	user := User{}
	return user.Name
}
