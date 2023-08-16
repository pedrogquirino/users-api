CREATE SCHEMA IF NOT EXISTS usersapi;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS usersapi.address (
	district varchar NULL,
	street varchar NULL,
	zipcode varchar NULL,
	"number" varchar NULL,
	id varchar NULL DEFAULT uuid_generate_v1(),
	user_id varchar NULL
);

ALTER TABLE "usersapi"."address" ADD CONSTRAINT address_fk FOREIGN KEY (user_id) REFERENCES "usersapi"."users"(id) ON DELETE CASCADE ON UPDATE CASCADE;