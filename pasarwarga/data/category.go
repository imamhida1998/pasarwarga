package data

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
)

type Kategori struct {
	BaseData
	Id            string `json:"category_id"`
	Category_name string `json:"category_name"`
	Category_slug string `json:"category_slug"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

func Createkategori(db *sql.DB) Kategori {
	kategori := Kategori{}
	kategori.DB = db
	return kategori
}

func (kategori Kategori) Migrate() error {

	sqlDropTable := "DROP TABLE IF EXISTS public.kategori CASCADE"

	sqlCreateTable := `
CREATE TABLE public.kategori
(
    category_id character varying NOT NULL,
    category_name character varying NOT NULL,
	category_slug character varying NOT NULL,
	created_at character varying NOT NULL,
	updated_at character varying,
	deleted_at character varying,
    CONSTRAINT kategori_pkey PRIMARY KEY (category_id)
)

TABLESPACE pg_default;

ALTER TABLE public.kategori
    OWNER to postgres;
	`

	_, err := kategori.DB.Exec(sqlDropTable)
	if err != nil {
		return err
	}

	_, err = kategori.DB.Exec(sqlCreateTable)
	if err != nil {
		return err
	}

	return nil

}

func (kategori Kategori) selectQuery() string {
	sql := `select category_id,category_name,category_slug,created_at,updated_at,deleted_at from kategori`
	return sql
}

func (kategori Kategori) fetchRow(cursor *sql.Rows) (Kategori, error) {
	cekdata := Kategori{}
	err := cursor.Scan(
		&cekdata.Id,
		&cekdata.Category_name,
		&cekdata.Category_slug,
		&cekdata.CreatedAt,
		&cekdata.UpdatedAt,
		&cekdata.DeletedAt)
	if err != nil {
		return Kategori{}, err
	}
	return cekdata, nil
}

func (kategori Kategori) CreateCategoryPost(Category_name, Category_slug string) (error, error) {
	sql := `insert into kategori (category_id,category_name,category_slug,created_at) values ($1,$2,$3,$4)`
	Id := uuid.New().ID()
	result, err := kategori.Exec(sql,
		Id,
		Category_name,
		Category_slug,
		time.Now().Format(time.RFC822))
	if err != nil {
		return err, nil
	}
	resultrows, err := result.RowsAffected()
	if err != nil {
		return err, nil
	}
	log.Println("Add Row=", resultrows)
	return nil, nil
}

func (kategori Kategori) UpdateCategory(Id, Category_name, Category_slug string) error {
	sql := `update kategori set category_name =$2,category_slug=$3,updated_at=$4 where category_id=$1`
	result, err := kategori.Exec(sql,
		Id,
		Category_name,
		Category_slug,
		time.Now().Format(time.RFC822))
	if err != nil {
		return err
	}
	resultRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("UpdatedCategory= ", resultRows)
	return nil
}
func (kategori Kategori) DeleteCategory(Id string) error {
	sql := `update kategori set deleted_at=$2 where category_id=$1`
	result, err := kategori.Exec(sql,
		Id,
		time.Now().Format(time.RFC822))
	if err != nil {
		return err
	}
	resultRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("UpdatedCategory= ", resultRows)
	return nil
}


func (kategori Kategori) DetailCategory() ([]Kategori, error) {
	sql := kategori.selectQuery()
	cursor, err := kategori.Query(sql)
	if err != nil {
		return nil, err
	}
	datakategori := []Kategori{}
	for cursor.Next() {
		dataset, err := kategori.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		datakategori = append(datakategori, dataset)
	}
	return datakategori, nil
}


func (kategori Kategori) RemoveArtikel(KategoriName string) (error, error) {
	sql := `delete from kategori where category_name=$1`
	result, err := kategori.Exec(sql, KategoriName)
	if err != nil {
		return err, nil
	}
	AffectedRows, err := result.RowsAffected()
	if err != nil {
		return err, nil
	}
	log.Println("Remove AffectedRows=", AffectedRows)
	return nil, nil
}
func (kategori Kategori) ReadCategory() (*Kategori, error) {
	sql := `select category_id,category_name,category_slug,created_at,updated_at from kategori where deleted_at=$1`
	deletedata := "nil"
	cursor, err := kategori.Query(sql, deletedata)
	if err != nil {
		return nil, err
	}
	if cursor.Next() {
		cek, err := kategori.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		return &cek, nil
	}
	return nil, errors.New("Nama Kategori Tidak Ditemukan")
}

func (kategori Kategori) FindCategory(kategori_name string) (*Kategori, error) {
	sql := `select category_id,category_name,category_slug,created_at,updated_at from kategori where deleted_at=$1`
	cursor, err := kategori.Query(sql, kategori_name)
	if err != nil {
		return nil, err
	}
	if cursor.Next() {
		cek, err := kategori.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		return &cek, nil
	}
	return nil, errors.New("Nama Kategori Tidak Ditemukan")
}
