package migrations

import (
	"database/sql"
	"fmt"
	"log"
)

func RunMigrations(db *sql.DB) {
	array := []string{
		`CREATE TABLE IF NOT EXISTS public.credentials (
			ID SERIAL PRIMARY KEY,
			login varchar(450) UNIQUE,
			password varchar(450)
		);`,
		`CREATE TABLE IF NOT EXISTS public.service (
			ID SERIAL PRIMARY KEY,
			fk_login integer,
			CONSTRAINT fk_login FOREIGN KEY (fk_login) REFERENCES credentials (ID),
			name text,
			image text,
			description text,
			value double precision,
			positive_evaluations integer,
			negative_evaluations integer
		);`,
		`CREATE TABLE IF NOT EXISTS public.commit (
			ID SERIAL PRIMARY KEY,
			fk_login integer,
			CONSTRAINT fk_login FOREIGN KEY (fk_login) REFERENCES credentials (ID),
			fk_service_page integer,
			CONSTRAINT fk_service_page FOREIGN KEY (fk_service_page) REFERENCES service (ID),
			commit text
		)`,
		`CREATE TABLE IF NOT EXISTS public.likes (
			fk_login integer NOT NULL,
			fk_service_page integer NOT NULL,
			liker integer,
			PRIMARY KEY(fk_login, fk_service_page)
		)`,
	}

	for _, migrate := range array {
		_, err := db.Exec(migrate)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(migrate)
	}

}
