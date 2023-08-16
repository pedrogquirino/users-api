CREATE DATABASE usersapi;

CREATE SCHEMA IF NOT EXISTS usersapi;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS usersapi.users (
	"name" varchar NULL,
	age varchar NULL,
	id varchar NOT NULL DEFAULT uuid_generate_v1(),
	email varchar NULL,
	"password" varchar NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_un UNIQUE (id)
);