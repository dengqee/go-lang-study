package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	priceStr := req.URL.Query().Get("price")
	newPrice, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "price value err: %q\n", priceStr)
		return
	}

	db[item] = dollars(newPrice)
	fmt.Fprintf(w, "update success %s: %s\n", item, priceStr)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) list(w http.ResponseWriter, req *http.Request) {
	var shopList = template.Must(template.New("shopList").Parse(`
<h1>shopList</h1>
<table>
<tr style='text-align:left'>
<th>item</th>
<th>	</th>
<th>price</th>
</tr>
</table>
`))
	shopList.Execute(w, nil)
	const templ = `<p>{{.A}} {{.B}}</p>`
	type data struct {
		A string
		B dollars
	}
	t := template.Must(template.New("escape").Parse(templ))
	for item, price := range db {
		var dat = data{item, price}
		err := t.Execute(w, dat)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
