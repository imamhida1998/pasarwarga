package data

import (
	"log"
	"pasarwarga/database/memory"
	"pasarwarga/database/postgresql"
	"pasarwarga/database/postgresql/config"
	"pasarwarga/helpers"
	"testing"
)

func TestArtikel(t *testing.T) {
	db := postgresql.OpenConnection(config.IMAM)
	Client := memory.RedisCnn()
	defer db.Close()

	branchModel := Createartikel(db)
	err := branchModel.Migrate()
	if err != nil {
		t.Fatal(err.Error())
	}

	// err = branchModel.Truncate()
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }
	id := "01"
	Title := "BRNCH-001"
	Kategori := "Alfa Poris Indah"
	Content := "123.456"

	err = branchModel.CreateArtikelPost(id, Title, Kategori, Content)

	if err != nil {
		t.Fatal(err.Error())
	}

	//update
	Kategori = "Alfa Poris Indah-2"
	Content = "123.456.789"

	err = branchModel.UpdateArtikel(id, Title, Kategori, Content)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = branchModel.DeletedAtArtikel(Title)
	if err != nil {
		t.Fatal(err.Error())
	}

	foundBranch, err := branchModel.FindArtikel(id)
	if err != nil {
		t.Fatal(err)
	}
	if foundBranch == nil {
		t.Fatal("FindOne=Branch=", foundBranch)
	}

	dataset, err := branchModel.ReadArtikel()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("OnUpdated", helpers.ToJson(dataset))
	Client.Set(id,helpers.ToJson(dataset),0)

	// err = branch.Remove(Title)
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }

	// branches, err = branch.GetBranchs()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if len(branches) != 0 {
	// 	t.Fatal("Expected=0, actual=", len(branches))
	// }

}
