package main

import (
	"fmt"

	"github.com/casbin/casbin"
)

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	//admin
	e.AddPermissionForUser("admin", "ad_campaign", "GET")
	e.AddPermissionForUser("admin", "ad_campaign", "LIST")
	//area-admin
	e.AddPermissionForUser("area_ad_admin", "campaign", "WRITE")
	e.AddPermissionForUser("area_ad_admin", "adgroup", "WRITE")
	e.AddPermissionForUser("area_ad_admin", "adcreative", "WRITE")
	//assigns role of area_ad_admin to admin.
	e.AddRoleForUser("admin", "area_ad_admin")

	//end users.
	e.AddRoleForUser("kevin", "admin")
	e.SavePolicy()

	//print.
	result, _ := e.GetImplicitPermissionsForUser("kevin")
	fmt.Printf("%#v", result)
}
