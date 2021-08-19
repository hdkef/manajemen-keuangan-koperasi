package driver

import (
	"encoding/json"
	"manajemen-keuangan-koperasi/models"

	"github.com/gomodule/redigo/redis"
)

func (C *RedisDriver) SetCacheMember(keyname string, member models.Member) error {
	//first marshall member struct to json
	data, err := json.Marshal(member)
	if err != nil {
		return err
	}
	//set member
	_, err = C.C.Do("SET", keyname, data)
	if err != nil {
		return err
	}
	return nil
}

func (C *RedisDriver) GetCacheMember(keyname string) (models.Member, error) {
	//get slice of byte of json
	data, err := redis.Bytes(C.C.Do("GET", keyname))
	if err != nil {
		return models.Member{}, err
	}
	//unmarshal json to struct
	var member models.Member
	err = json.Unmarshal(data, &member)
	if err != nil {
		return models.Member{}, err
	}
	return member, nil
}

func (C *RedisDriver) DelCacheMember(keyname string) error {
	_, err := C.C.Do("Del", keyname)
	return err
}
