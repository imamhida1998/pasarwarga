package data

import (
	"pasarwarga/database/postgresql"
	"pasarwarga/database/postgresql/config"
	"testing"
)

func TestKategori(t *testing.T) {
	db := postgresql.OpenConnection(config.IMAM)
	defer db.Close()

	ModelArtikel := Createkategori(db)
	err := ModelArtikel.Migrate() // createdb
	if err != nil {
		t.Fatal(err.Error())
	}
	Id := "Artikel-01"
	Category_Name := "Hello Guys"
	Category_Slug := "SayHello"
	//Content := "Test Hello"
	err, _ = ModelArtikel.CreateCategoryPost(Category_Name,Category_Slug)
	if err != nil {
		t.Fatal(err)
	}
	err = ModelArtikel.UpdateCategory(Id,Category_Name,Category_Slug)
	if err != nil {
		t.Fatal(err)
	}

	err = ModelArtikel.DeleteCategory(Id)
	if err != nil {
		t.Fatal(err)
	}
	//
	//dataset , err := ModelArtikel.ReadCategory()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//log.Println("Data", helpers.ToJson(dataset))

}