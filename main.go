package main

import (
	"fmt"

	"github.com/casbin/casbin"
)

/*
User story
 - Users in system User01, User02
 - Group 1 : InvoiceNo, InvoiceDate,ManufactuererID, Origin,MineName
 - Group 2 : Origin, KPNumber, Weight

1. As manufacturer : manu-id-01 , I want to give access group_1 perm to User01
2. As manufacturer: manu-did-02, I want to give access group_2 perm to User02

3. As manufacturer: manu-id-01 , I want to give access to group_2 only for oid "object-id-01" to "User01"
4. As manufacturer: manu-id-01 , I want to give access to group 1
*/

func printUserPermissions(e *casbin.Enforcer, usr string) {
	fmt.Println("Permissions for user : ", usr)
	result, _ := e.GetImplicitPermissionsForUser(usr)
	for _, v := range result {
		fmt.Printf("%s\n", v)
	}
	fmt.Println("=============================")
}

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")

	//UID 100 - OBI 100
	//UID 200 -!OBI 100

	//Group 1 Permission Manufacturers
	e.AddPermissionForUser("role:manu-id-01_group", "InvoiceNo")
	e.AddPermissionForUser("role:manu-id-01_group", "InvoiceDate")
	e.AddPermissionForUser("role:manu-id-01_group", "Origin")

	//Group 2 Permission Manufacturers
	e.AddPermissionForUser("role:manu-id-02_group_1", "InvoiceNo")
	e.AddPermissionForUser("role:manu-id-02_group_1", "InvoiceDate")
	e.AddPermissionForUser("role:manu-id-02_group_1", "Origin")

	//Group 3 Permission Manufacturer no 2
	e.AddPermissionForUser("role:manu-id-02_group_2", "Origin")
	e.AddPermissionForUser("role:manu-id-02_group_2", "KPNumber")
	e.AddPermissionForUser("role:manu-id-02_group_2", "Weight")

	//Haresh gets only group 1
	e.AddRoleForUser("haresh", "role:manu-id-02_group_2")
	e.AddRoleForUser("haresh", "role:manu-id-01_group")

	e.AddPermissionForUser("haresh", "dia", "oid", "muid", "Origin", "deny")

	e.SavePolicy()

	//print all permissions
	printUserPermissions(e, "haresh")
}
