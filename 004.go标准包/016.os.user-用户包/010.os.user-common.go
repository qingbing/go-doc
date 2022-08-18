package main

import (
	"fmt"
	"log"
	"os/user"
)

func printUser(u *user.User, msg string) {
	fmt.Printf("====== %s =======\n", msg)
	fmt.Printf("Name: %+v\n", u.Name)
	fmt.Printf("Uid: %+v\n", u.Uid)
	fmt.Printf("Gid: %+v\n", u.Gid)
	fmt.Printf("HomeDir: %+v\n", u.HomeDir)
	fmt.Printf("Username: %+v\n", u.Username)
	gs, err := u.GroupIds()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("用户名: %+v\n", gs)
	}
}

func printGroup(g *user.Group, msg string) {
	fmt.Printf("====== %s =======\n", msg)
	fmt.Printf("Name: %+v\n", g.Name)
	fmt.Printf("Gid: %+v\n", g.Gid)
}

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	} else {
		printUser(u, "user.Current")
	}

	u, err = user.Lookup("root")
	if err != nil {
		log.Fatal(err)
	} else {
		printUser(u, "user.Lookup")
	}

	u, err = user.LookupId("501")
	if err != nil {
		log.Fatal(err)
	} else {
		printUser(u, "user.LookupId")
	}

	g, err := user.LookupGroup("admin")
	if err != nil {
		log.Fatal(err)
	} else {
		printGroup(g, "user.LookupGroup")
	}

	g, err = user.LookupGroupId("12")
	if err != nil {
		log.Fatal(err)
	} else {
		printGroup(g, "user.LookupGroup")
	}
}
