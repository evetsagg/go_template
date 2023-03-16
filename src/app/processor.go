package main

import (
	"encoding/json"
	"go_template/src/database"
	"go_template/src/model"
	"io"
	"net/http"

	business "go_template/src/business"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
)

type Processor struct {
	logger     LoggingI
	prop       *properties.Properties
	productDao *database.ProductDao
}

// func NewProcessor(logger LoggingI, prop *properties.Properties, productDao *database.ProductDao) *Processor {
// 	return &Processor{logger: logger, prop: prop, productDao: productDao}
// }

func NewProcessor(logger LoggingI, prop *properties.Properties) *Processor {
	return &Processor{logger: logger, prop: prop}
}

func (p *Processor) returnAllArticles(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("RETURN ALL ARTICLES")
	//If needed, delegate to some business logic , using dao
	businessImpl := business.New(p.logger)
	businessImpl.ProcessSomething()
	json.NewEncoder(w).Encode(Articles)
}

func (p *Processor) returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func (p *Processor) createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	//
	reqBody, _ := io.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func (p *Processor) deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func (p *Processor) createProduct(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	//
	reqBody, _ := io.ReadAll(r.Body)

	var product model.Product
	json.Unmarshal(reqBody, &product)
	// update our global Articles array to include
	// our new Article
	p.productDao.Create(&product)
	json.NewEncoder(w).Encode(product)

}
func (p *Processor) returnSingleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	//intVar, err := strconv.Atoi(key)
	// if err != nil{
	//     p.logger.Error(err)
	//     w.WriteHeader(http.StatusBadRequest)
	//     w.Header().Set("Content-Type", "application/json")
	//     resp := make(map[string]string)
	//     resp["message"] = "id must be integer value"
	//     jsonResp, err := json.Marshal(resp)
	//     if err != nil {
	//         p.logger.Fatal(err)
	//     }
	//     w.Write(jsonResp)
	// } else {
	product := &model.Product{Id: key}
	result := p.productDao.FindById(product)
	json.NewEncoder(w).Encode(result)

	//}

}
func (p *Processor) returnAllProducts(w http.ResponseWriter, r *http.Request) {

	//If needed, delegate to some business logic , using dao
	results := p.productDao.FindAll()
	json.NewEncoder(w).Encode(results)
}

// Sample object
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func init() {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}
