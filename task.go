package jobq

type Task interface {
	Run() error
	Label() string
}
