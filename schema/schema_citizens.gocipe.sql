-- generated by gocipe 2cdecd62f8cc1e3e1b51a96e61f1519e5d31bc3eb806037239eeb7707acf58b5; DO NOT EDIT

DROP TABLE IF EXISTS citizens;

CREATE TABLE citizens (
	"id" CHAR(36),
	"surname" VARCHAR(255) NOT NULL,
	"othernames" VARCHAR(255) NOT NULL,
	"gender" CHAR(1) NOT NULL,
	"dob" DATE NOT NULL,
	
	PRIMARY KEY ("id")
);
