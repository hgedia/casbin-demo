package main

import (
	"fmt"

	"github.com/casbin/casbin"
)

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")

	e.AddPermissionForUser(":id", "Origin", "read")

	//Group 1 Permission
	e.AddPermissionForUser("group_1", "InvoiceNo", "read")
	e.AddPermissionForUser("group_1", "InvoiceDate", "read")
	e.AddPermissionForUser("group_1", "ManufacturerId", "read")
	e.AddPermissionForUser("group_1", "Origin", "read")
	e.AddPermissionForUser("group_1", "MineName", "read")

	e.AddPermissionForUser("group_2", "Origin", "read")
	e.AddPermissionForUser("group_2", "KPNumber", "read")
	e.AddPermissionForUser("group_2", "Weight", "read")

	//Admin gets all permisssions
	e.AddRoleForUser("admin", "group_1")
	e.AddRoleForUser("admin", "group_2")

	//Haresh gets only group 1
	e.AddRoleForUser("haresh", "group_1")
	e.AddRoleForUser("priyav", "group_2")
	e.AddRoleForUser("david", "admin")
	e.SavePolicy()

	//print all permissions
	result, _ := e.GetImplicitPermissionsForUser("haresh")
	fmt.Printf("%#v\n", result)

	//Check if user has permission

	fmt.Println(e.HasPermissionForUser("haresh", "group_1"))
}
