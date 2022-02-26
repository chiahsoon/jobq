package jobq

type Task interface {
	Desc() string
	Run()
}
