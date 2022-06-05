package service

import (
	"github.com/google/uuid"
	"github.com/vitorvidaldev/msn/src/domain/entity"
	"github.com/vitorvidaldev/msn/src/domain/repository"
	"github.com/vitorvidaldev/msn/src/domain/vo"
)

func CreateUser(createVO vo.CreateUserVO) vo.UserVO {
	userEntity := entity.User{
		User_id:  uuid.New(),
		Email:    createVO.Email,
		Username: createVO.Username,
		Password: createVO.Password}

	userEntity = repository.InsertUser(userEntity)

	return vo.UserVO{
		User_id:  userEntity.User_id,
		Email:    userEntity.Email,
		Username: userEntity.Username}
}

func LoginUser(loginVO vo.LoginUserVO) vo.UserVO {

}

func GetUser(userId uuid.UUID) vo.UserVO {

}

func DeleteUser(userId uuid.UUID) {

}
