package service

import (
	"context"
	"fmt"
	proxy "go-grpc-server/service/api"
)

type StudentService struct {
}

func (this *StudentService) GetStudentsByClass(ctx context.Context, request *proxy.StudentRequest) (*proxy.StudentResponse, error) {
	fmt.Println("客户请求: GetStudentsByClass 方法")
	var students int32 = 0
	if request.Class == proxy.Class_C1901 {
		students = 1901
	} else if request.Class == proxy.Class_C1902 {
		students = 1902
	} else if request.Class == proxy.Class_C1903 {
		students = 1903
	} else {
		students = 1910
	}
	return &proxy.StudentResponse{Students: students}, nil
}
