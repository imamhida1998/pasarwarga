package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pasarwarga/data"
	"pasarwarga/database/memory"
	"pasarwarga/database/postgresql"
	"pasarwarga/database/postgresql/config"
	"pasarwarga/helpers"
)

type KategoriApi struct {
	BaseApi

	data.Kategori

}
func(kategoriApi KategoriApi)CreateKategori(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&kategoriApi)
	if err != nil {
		kategoriApi.Error(w,err)
	}
	client := memory.RedisCnn()
	db := postgresql.OpenConnection(config.IMAM)
	defer db.Close()
	defer client.Close()
	ModelArtikel := data.Createkategori(db)
	dataset, err := ModelArtikel.CreateCategoryPost(kategoriApi.Category_name,kategoriApi.Category_slug)
	if err != nil {
		kategoriApi.Error(w, err)
		return
	} else {

		kategoriApi.Json(w,dataset,http.StatusOK)
		err = client.Set(kategoriApi.Category_name, helpers.ToJson(dataset), 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (kategoriApi KategoriApi)UpdateArtikel(w http.ResponseWriter,r *http.Request){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&kategoriApi)
	if err != nil {
		kategoriApi.Error(w , err)
	}
	db := postgresql.OpenConnection(config.IMAM_POSTSQL)
	defer db.Close()

	ModelArtikel := data.Createkategori(db)
	err = ModelArtikel.UpdateCategory(kategoriApi.Id,kategoriApi.Category_name,kategoriApi.Category_slug)
	if err != nil {
		kategoriApi.Error(w , err)
		return
	} else {
		kategoriApi.Json(w , ModelArtikel , http.StatusOK)
		//err = client.Set(kategoriApi.Category_name, helpers.ToJson(dataset), 0).Err()
	}

}

