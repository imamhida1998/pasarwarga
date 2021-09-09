package data

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Artikel struct {
	BaseData
	Id          string `json:"artikel_id"`
	Title       string `json:"title"`
	Category_id string `json:"category_id"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

func Createartikel(db *sql.DB) Artikel {
	artikel := Artikel{}
	artikel.DB = db
	return artikel
}

func CreateVirtualAccountWithTransaction(transaction *sql.Tx) Artikel {
	artikel := Artikel{}
	artikel.Transaction = transaction
	artikel.UseTransaction = true
	return artikel
}

func (artikel Artikel) Migrate() error {

	sqlDropTable := "DROP TABLE IF EXISTS public.artikel CASCADE"

	sqlCreateTable := `
CREATE TABLE public.artikel
(
    artikel_id character varying NOT NULL,
    title character varying NOT NULL,
	category_id character varying NOT NULL,
	content text NOT NULL,
	created_at character varying NOT NULL,
	updated_at character varying,
	deleted_at character varying,
    CONSTRAINT artikel_pkey PRIMARY KEY (artikel_id)
)

TABLESPACE pg_default;

ALTER TABLE public.artikel
    OWNER to postgres;
	`

	_, err := artikel.DB.Exec(sqlDropTable)
	if err != nil {
		return err
	}

	_, err = artikel.DB.Exec(sqlCreateTable)
	if err != nil {
		return err
	}

	return nil

}

func (artikel Artikel) selectQuery() string {
	sql := `select artikel_id,title,category_id,content,created_at,updated_at,deleted_at from artikel`
	return sql
}

func (artikel Artikel) fetchRow(cursor *sql.Rows) (Artikel, error) {
	cekdata := Artikel{}
	err := cursor.Scan(
		&cekdata.Id,
		&cekdata.Title,
		&cekdata.Category_id,
		&cekdata.Content,
		&cekdata.CreatedAt,
		&cekdata.UpdatedAt,
		&cekdata.DeletedAt)
	if err != nil {
		return Artikel{}, err
	}
	return cekdata, nil
}

func (artikel Artikel) CreateArtikelPost(id, Title, Category_id, Content string) error {
	sql := `insert into artikel (artikel_id, title, category_id, content, created_at) values 
	($1, $2, $3, $4, $5)`
	result, err := artikel.Exec(sql, id, Title, Category_id, Content, time.Now().Format(time.RFC1123))
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Artikel.Add.AffectedRows=", affectedRows)
	return nil
}

func (artikel Artikel) UpdateArtikel(Id string, Title, Category_id, Content string) error {
	sql := `update artikel 
	set title=$2,
	category_id=$3,
	content=$4,
	updated_at=$5

	where artikel_id=$1
	`
	result, err := artikel.Exec(sql, Id, Title, Category_id, Content, time.Now().Format(time.RFC1123))
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Artikel.Update.AffectedRows=", affectedRows)

	return nil
}

func (artikel Artikel) DeletedAtArtikel(Title string) error {
	sql := `update artikel 
	set deleted_at=$2

	where title=$1
	`
	result, err := artikel.Exec(sql, Title, time.Now().Format(time.RFC1123))
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Artikel.Update.AffectedRows=", affectedRows)

	return nil
}

func (artikel Artikel) ReadArtikel() ([]Artikel, error) {
	sql := artikel.selectQuery()
	cursor, err := artikel.Query(sql)
	if err != nil {
		return nil, err
	}

	artikels := []Artikel{}

	for cursor.Next() { //looping rows
		b, err := artikel.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		artikels = append(artikels, b)
	}

	return artikels, nil
}

func (artikel Artikel) RemoveArtikel(ArtikelName string) error {
	sql := "delete from artikel where artikel_id=$1"
	result, err := artikel.Exec(sql, ArtikelName)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("Artikel.Remove.AffectedRows=", affectedRows)
	return nil
}

func (artikel Artikel) FindArtikel(artikel_id string) (*Artikel, error) {
	sql := `select * from artikel where artikel_id=$1`
	cursor, err := artikel.Query(sql, artikel_id)
	if err != nil {
		return nil, err
	}
	if cursor.Next() {
		b, err := artikel.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		return &b, nil
	}
	return nil, errors.New("Artikel tidak ditemukan")
}
