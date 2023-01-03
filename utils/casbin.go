package utils

// import (
// 	"chat/constant"
// 	. "chat/global"
// 	"errors"
// 	"sync"

// 	"github.com/casbin/casbin/v2"
// 	gormadapter "github.com/casbin/gorm-adapter/v3"
// )

// var (
// 	enforcer *casbin.SyncedEnforcer
// 	once     sync.Once
// )

// func init() {
// 	Logger(constant.Info, "casbin", map[string]string{"mes": "init casbin"})
// 	once.Do(func() {
// 		adapter, _ := gormadapter.NewAdapterByDB(Db)
// 		enforcer, _ = casbin.NewSyncedEnforcer("conf/model.conf", adapter)
// 	})
// 	err := enforcer.LoadPolicy()
// 	Error(err, "casbin", "load policy failed")
// }

// func Authorization(role string, source string, operation constant.OperationEnums) bool {

// 	ok, _ := enforcer.Enforce(role, source, operation)
// 	return ok
// }

// func Police(role string, source string, operation constant.OperationEnums) (bool, error) {

// 	switch operation {

// 	case constant.ADD:
// 		not_exist, err := enforcer.AddPolicy(role, source, operation)
// 		Error(err, "InitConfig", "read config failed")
// 		return !not_exist, nil
// 	case constant.DELETE:
// 		not_exist, err := enforcer.RemovePolicy(role, source, operation)
// 		Error(err, "InitConfig", "read config failed")
// 		return !not_exist, nil
// 	case constant.READ:
// 		// enforcer.UpdatePolicy([]string{})
// 	case constant.UPDATE:

// 	default:
// 		return false, errors.New("got an invalid operation value")
// 	}
// 	return true, nil
// }
