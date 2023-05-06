package main

import (
	"fmt"
	"goflatbuffer/AdminService"

	flatbuffers "github.com/google/flatbuffers/go"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	builder := flatbuffers.NewBuilder(1024)
	userName := builder.CreateString("IsaacDSC")
	email := builder.CreateString("isaac8.silva@hotmail.com")
	AdminService.UserStart(builder)
	AdminService.UserAddId(builder, 1)
	AdminService.UserAddName(builder, userName)
	AdminService.UserAddEmail(builder, email)
	user := AdminService.UserEnd(builder)
	builder.Finish(user)

	buf := builder.FinishedBytes()
	userDeserialized := AdminService.GetRootAsUser(buf, 0)

	fmt.Printf("%+v\n",
		User{
			Id:    int(userDeserialized.Id()),
			Name:  string(userDeserialized.Name()),
			Email: string(userDeserialized.Email()),
		})
}
