package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"net/http/httputil"
	"sort"
	"sync"
)

func main() {
	fmt.Println("Here we go!")
	db := NewDB()

	http.Handle("/", http.FileServer(http.Dir("www")))

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, true)
		fmt.Printf("%s\n", b)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "{}")
	})

	//API Stubs
	http.HandleFunc("/api/quotes/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var quote Quote
			d := json.NewDecoder(r.Body)
			err := d.Decode(&quote)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			quote.Email = html.EscapeString(quote.Email)
			quote.Quote = html.EscapeString(quote.Quote)
			db.SetQuote(quote)

			w.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(w)
			err = e.Encode(&quote)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "GET":
			quotes := db.GetSortedQuotes()
			w.Header().Set("Content-Type", "application/json")
			e := json.NewEncoder(w)
			err := e.Encode(&quotes)
			if err != nil {
				fmt.Println(err)
				return
			}
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	http.HandleFunc("/api/vote", func(w http.ResponseWriter, r *http.Request) {
		db.VoteForQuote(r.URL.Query().Get("email"))
		quote := db.GetQuote(r.URL.Query().Get("email"))
		w.Header().Set("Content-Type", "application/json")
		e := json.NewEncoder(w)
		err := e.Encode(&quote)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}

// DB CODE
type DB struct {
	sync.RWMutex
	quotes map[string]Quote
}

func NewDB() *DB {
	var db = new(DB)
	db.quotes = make(map[string]Quote)
	return db
}

func (db *DB) SetQuote(quote Quote) {
	//Write Lock
	db.Lock()
	defer db.Unlock()
	quote.Vote = 0
	db.quotes[quote.Email] = quote
}

func (db *DB) VoteForQuote(email string) {
	db.Lock()
	defer db.Unlock()
	if v, ok := db.quotes[email]; ok {
		v.Vote = v.Vote + 1
		db.quotes[email] = v
	}
}

func (db *DB) GetQuote(email string) Quote {
	//Read Lock
	db.RLock()
	defer db.RUnlock()
	return db.quotes[email]
}

func (db *DB) GetSortedQuotes() []Quote {
	//Read Lock
	db.RLock()
	defer db.RUnlock()
	ret := make([]Quote, 0)
	for _, v := range db.quotes {
		ret = append(ret, v)
	}
	sort.Sort(ByVote(ret))

	return ret
}

type Quote struct {
	Email string `json:"email"`
	Quote string `json:"quote"`
	Vote  int64  `json:"vote"`
}

type ByVote []Quote

func (a ByVote) Len() int           { return len(a) }
func (a ByVote) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVote) Less(i, j int) bool { return a[i].Vote < a[j].Vote }
