package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20190318_232922 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20190318_232922{}
	m.Created = "20190318_232922"

	migration.Register("CreateUsersTable_20190318_232922", m)
}

// Run the migrations
func (m *CreateUsersTable_20190318_232922) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS `users` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,`user_name` varchar(255) NOT NULL DEFAULT ''  UNIQUE,`password` varchar(255) NOT NULL DEFAULT '' ,`created` datetime NOT NULL,`updated` datetime NOT NULL) ENGINE=InnoDB;")

}

// Reverse the migrations
func (m *CreateUsersTable_20190318_232922) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
