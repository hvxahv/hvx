package services

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"hvxahv/pkg/database"
	"hvxahv/pkg/structs"
	"log"
)

// GetAccountData 获取账户数据，将数据返回给调用者, 在用户登录之后调用的方法
func GetAccountData(u string) *structs.Accounts {
	a := new(structs.Accounts)

	db := database.GetMaria()

	db.Debug().Table("accounts").Where("username = ?", u).First(&a)
	return a
}

// GetActorData 获取 Actor ,
func GetActorData(u string) *structs.Accounts {
	a := new(structs.Accounts)

	rdb := database.GetRDB()
	db := database.GetMaria()
	// 判断查询的 key 是否存在,如果不存在, 将在数据库中查询并将数据持久化到 redis
	isKey, err := redis.Bool(rdb.Do("EXISTS", u))
	if err != nil {
		log.Println("检查 keys 时出错:", err)
	}
	if isKey != true {
		db.Debug().Table("accounts").Where("username = ?", u).First(&a)
		go accountCache("SET", a, u)
		return a
	} else {
		a := accountCache("GET", a, u)
		return a
	}
}

// 用于验证用户登录的方法, 增加了判断用户是否在数据库中存在的查询
// VerifyAccount
func VerifyAccounts(u string) *structs.Accounts {
	a := new(structs.Accounts)
	db := database.GetMaria()

	if db.Debug().Table("accounts").Where("username = ?", u).First(&a).RecordNotFound() {
		db.Debug().Table("accounts").Where("username = ?", u).First(&a)
		return a
	} else  {
		return a
	}
}

// accountCache 使用缓存数据库避免数据库的重复查询
func accountCache(method string, a *structs.Accounts, k string) *structs.Accounts {
	rdb := database.GetRDB()
	switch method {
	case "SET":
		v, _ := json.Marshal(a)
		if _, err := rdb.Do("SETNX", k, v); err != nil {
			log.Printf("Actor 持久化到数据库失败: %s", err)
		}
	case "GET":
		res, err := redis.Bytes(rdb.Do("GET", k))
		if err != nil {
			log.Println("Redis 获取 Actor 数据失败:", err)
		}
		if err := json.Unmarshal(res, &a); err != nil {
			log.Println("将 redis 取到的 Actor 数据 unmarshal 失败:", err)
		}
		return a
	default:

	}
	return nil
}