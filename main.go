package main

import (
	"fmt"

	"github.com/casbin/casbin"
)

func printUserPermissions(e *casbin.Enforcer, usr string) {
	fmt.Println("Persmissions for user : ", usr)
	result, _ := e.GetImplicitPermissionsForUser(usr)
	for _, v := range result {
		fmt.Printf("%s:%s\n", v[0], v[1])
	}
	fmt.Println("=============================")
}

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")

	//UID 100 - OBI 100
	//UID 200 -!OBI 100

	//Group 1 Permission Manufacturers
	e.AddPermissionForUser("group_1", "InvoiceNo", "read")
	e.AddPermissionForUser("group_1", "InvoiceDate", "read")
	e.AddPermissionForUser("group_1", "ManufacturerId", "read")
	e.AddPermissionForUser("group_1", "Origin", "read")
	e.AddPermissionForUser("group_1", "MineName", "read")

	//Group 2 Permissions Retailers
	e.AddPermissionForUser("group_2", "Origin", "read")
	e.AddPermissionForUser("group_2", "KPNumber", "read")
	e.AddPermissionForUser("group_2", "Weight", "read")

	//Admin gets all permisssions
	e.AddRoleForUser("admin", "group_1")
	e.AddRoleForUser("admin", "group_2")

	e.AddRoleForUser("uuid-100-204-100", "group_1")

	//Haresh gets only group 1
	e.AddRoleForUser("haresh", "group_1")
	e.AddRoleForUser("priyav", "group_2")
	e.AddRoleForUser("david", "admin")
	e.SavePolicy()

	//print all permissions
	printUserPermissions(e, "haresh")
	printUserPermissions(e, "priyav")
	printUserPermissions(e, "david")

	//Check if user has permission
	fmt.Println(e.HasPermissionForUser("haresh", "InvoiceNo", "read"))
	fmt.Println(e.HasPermissionForUser("group_1", "InvoiceNo", "read"))
}
