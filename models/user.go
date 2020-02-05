/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:25:34
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-03 00:49:02
 */
package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

const PROFILE_TABNAME = "tb_profile"

//用户
type User struct {
	Id        int64    //beego默认Id为主键,且自增长
	Nickname  string   `orm:"unique"` //用户名唯一
	Password  string   //密码
	Status    int64    //认证状态
	AvatarUrl string   //头像
	Profile   *Profile `orm:"rel(one)"` //设置一对一的反向关系
	//Manager   *Manager `orm:"rel(one)"`
	Posts []*Post `json:"post" orm:"reverse(many)"`
}

//判断一个用户是否存在
func OneUserExist(nickName string) bool {
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable("tb_user").Filter("nickname", nickName).One(&user)
	if err != orm.ErrNoRows {
		return true
	}
	return false
}

//返回一个用户的id
func GetOneuserID(nickName string) int64 {
	o := orm.NewOrm()
	user := User{}
	_ = o.QueryTable("tb_user").Filter("nickname", nickName).One(&user)
	return user.Id
}

func UpdateProfile(userID interface{}, userInfo map[string]interface{}) error {
	name := userInfo["name"].(string)
	stuId := userInfo["stu_id"].(string)
	school := userInfo["school"].(string)
	profession := userInfo["profession"].(string)
	sex, _ := strconv.Atoi(userInfo["sex"].(string))
	qqNumber := userInfo["qq_number"].(string)
	email := userInfo["email"].(string)
	telNum := userInfo["telNum"].(string)
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("tb_user").Filter("id", userID.(int64)).One(&user)
	proId := user.Profile.Id
	ret, err := o.Raw("UPDATE tb_profile SET name = ?,stu_id = ?,school = ?,profession=?, sex = ?, q_q_number = ?,email = ?,tel_number = ? WHERE id = ?", name, stuId, school, profession, sex, qqNumber, email, telNum, proId).Exec()
	fmt.Println(ret)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(userId int64) (*User, error) {
	user := User{Id: userId}
	o := orm.NewOrm()
	err := o.Read(&user)
	return &user, err
}