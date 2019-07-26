# rest-example-go
Dump database

    DROP TABLE IF EXISTS "users";
    DROP SEQUENCE IF EXISTS users_seq;
    CREATE SEQUENCE users_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;
    
    CREATE TABLE "public"."users" (
        "id" integer DEFAULT nextval('users_seq') NOT NULL,
        "first_name" character varying(255) NOT NULL,
        "last_name" character varying(255) NOT NULL,
        "email" character varying(255) NOT NULL,
        "password" character varying(255) NOT NULL,
        "sex" smallint DEFAULT '1' NOT NULL,
        CONSTRAINT "users_pkey" PRIMARY KEY ("id")
    ) WITH (oids = false);
    
    INSERT INTO "users" ("id", "first_name", "last_name", "email", "password", "sex") VALUES
    (2,	'My Name',	'My Name',	'My@Name.ru',	'My Name',	1),
    (3,	'Duncan',	'Stewart',	'mogopemis@mailinator.net',	'1',	1),
    (4,	'Teagan',	'Harrison',	'adminweb@adminweb.com',	'2',	1),
    (5,	'Duncan1',	'Stewart',	'mogopemis@mailinator.net',	'',	1),
    (10,	'Jon',	'Doe',	'qwe@qwe.ru',	'a',	0),
    (13,	'Duncan2',	'Stewart2',	'mog@mailinator.net',	'w',	1),
    (15,	'Josh',	'Green',	'qwe@qwe.ru',	'a',	0),
    (1,	'Qwer',	'Qwertov',	'qwert@mail.ru',	'',	1);

Example config.json

    {
    
    "database": {
    
    "connectionType": "postgres",
    
    "connectionString": "user=postgres dbname=apirepository sslmode=disable"
    
    },
      
    "server": {
    
    "port": 3000
    
    }
    
    }
