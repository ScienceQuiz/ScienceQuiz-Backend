package _interface

import "github.com/ScienceQuiz-Backend/model"

type User interface {
	// 새 사용자 생성 메서드
	Create(*model.User)
	// (userId 기준) 사용자 존재 여부 확인 메서드
	CheckUserExist(string) bool
	// 매개변수로 받은 문자열의 아이디를 가지고 있는 사용자 구조체 반환
	GetUserById(string) *model.User
}
