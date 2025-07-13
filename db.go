package main

type DB struct {
	DB_URL      string
	DB_USERNAME string
	DB_PASSWORD string
}

func (d *DB) Conn() error {
	return nil
}