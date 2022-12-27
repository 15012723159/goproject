package model

import "fmt"

func SelectMember(id int) Member {
	member := new(Member)
	error := DbCon.Debug().Where("id = ?", id).Find(member).Error

	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println("id = ", id)
	fmt.Println(member)
	return *member

}
