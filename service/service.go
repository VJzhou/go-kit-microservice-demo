package service

// 定义接口
type Service interface {
	Add (num1 ,num2 int) int
}

type addService struct {}

// 工厂函数
func NewService() Service {
	return &addService{}
}

// 实现接口
func (a addService) Add(num1, num2 int) int {
	return num1 + num2
}
