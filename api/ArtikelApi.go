package api

import (
	"encoding/json"
	"net/http"
	"pasarwarga/data"
	"pasarwarga/database/memory"
	"pasarwarga/database/postgresql"
	"pasarwarga/database/postgresql/config"
	"pasarwarga/helpers"
)

type ArtikelAPI struct {
	BaseApi

	data.Artikel
}

func (artikelApi ArtikelAPI) CreateArtikel(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&artikelApi)
	if err != nil {
		artikelApi.Error(w, err)
		return
	}

	db := postgresql.OpenConnection(config.IMAM)
	client := memory.RedisCnn()
	defer db.Close()
	ModelArtikel := data.Createartikel(db)
	err = ModelArtikel.CreateArtikelPost(artikelApi.Id, artikelApi.Title, artikelApi.Category_id, artikelApi.Content)
	if err != nil {
		artikelApi.Error(w, err)
		return
	} else {
		branch, err := ModelArtikel.FindArtikel(artikelApi.Id)
		if err != nil {
			artikelApi.Error(w, err)
			return
		} else {
			artikelApi.Json(w, branch, http.StatusOK)
			client.Set(artikelApi.Title,helpers.ToJson(branch),0 )
			return
		}
	}
}

func (artikelApi ArtikelAPI) UpdateArtikel(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&artikelApi)
	if err != nil {
		artikelApi.Error(w, err)
		return
	}

	db := postgresql.OpenConnection(config.IMAM)
	//client := memory.RedisCnn()
	defer db.Close()

	ModelArtikel := data.Createartikel(db)
	err = ModelArtikel.UpdateArtikel(artikelApi.Id, artikelApi.Title, artikelApi.Category_id, artikelApi.Content)
	if err != nil {
		artikelApi.Error(w, err)
		return
	} else {
		branch, err := ModelArtikel.FindArtikel(artikelApi.Id)
		if err != nil {
			artikelApi.Error(w, err)
			return
		} else {
			artikelApi.Json(w, branch, http.StatusOK)
			return
		}
	}

}

func (artikelApi ArtikelAPI) ListArtikel(w http.ResponseWriter, r *http.Request) {
	db := postgresql.OpenConnection(config.IMAM_POSTSQL)
	//client := memory.RedisCnn()
	defer db.Close()
	ModelArtikel := data.Createartikel(db)
	Artikels, err := ModelArtikel.ReadArtikel()
	if err != nil {
		artikelApi.Error(w, err)
		return
	} else {
		artikelApi.Json(w, Artikels, http.StatusOK)
	}

}

// func (artikelApi ArtikelAPI) DetailArtikelByTitle(w http.ResponseWriter, r *http.Request) {
// 	db := postgresql.OpenConnection(config.IMAM_POSTSQL)
// 	defer db.Close()
// 	FindTitle := artikelApi.QueryParam(r, "")
// 	ModelArtikel := data.Createartikel(db)
// 	Artikel, err := ModelArtikel.FindArtikel(FindTitle)
// 	if err != nil {
// 		artikelApi.Error(w, err)
// 		return
// 	} else {
// 		artikelApi.Json(w, Artikel, http.StatusOK)
// 	}

// }

// func (artikelApi ArtikelAPI) DeleteArtikelByArtikel(w http.ResponseWriter, r *http.Request) {
// 	db := postgresql.OpenConnection(config.IMAM_POSTSQL)
// 	defer db.Close()
// 	Title := artikelApi.QueryParam(r, "title")
// 	branchModel := data.Createartikel(db)
// 	branch, err := branchModel.de(Title)
// 	if err != nil {
// 		artikelApi.Error(w, err)
// 		return
// 	} else {
// 		artikelApi.Json(w, branch, http.StatusOK)
// 		return
// 	}
// }
