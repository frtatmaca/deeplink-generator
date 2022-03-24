package constant

const CREATE_DATABASE = "CREATE DATABASE IF NOT EXISTS "
const CREATE_LINK_TABLE = "CREATE TABLE IF NOT EXISTS links(id bigint(20) primary key auto_increment, link longtext, deep_link longtext, type int, created_at datetime, updated_at datetime, deleted_at datetime)"
const INSERT_LINK = "INSERT INTO `links` (`created_at`,`updated_at`,`deleted_at`,`link`,`deep_link`,`type`) VALUES (?,?,?,?,?,?)"
