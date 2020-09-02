package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type clickSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x clickSort) Len() int {
	return len(x.t)
}
func (x clickSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}
func (x clickSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

var clicks []string

func less(x, y *Track) bool {
	for _, click := range clicks {
		switch click {
		case "Title":
			if x.Title == y.Title {
				continue
			}
			return x.Title < y.Title
		case "Year":
			if x.Year == y.Year {
				continue
			}
			return x.Year < y.Year
		case "Artist":
			if x.Artist == y.Artist {
				continue
			}
			return x.Artist < y.Artist
		case "Album":
			if x.Album == y.Album {
				continue
			}
			return x.Album < y.Album
		case "Length":
			if x.Length == y.Length {
				continue
			}
			return x.Length < y.Length
		}
	}
	return false
}

func main() {
	//clicks = []string{"Title", "Year", "Length", "Artist", "Album"}
	//sort.Sort(clickSort{tracks, less})
	//printTracks(tracks)
	//
	//clicks = []string{"Year", "Title", "Length", "Artist", "Album"}
	//sort.Sort(clickSort{tracks, less})
	//printTracks(tracks)
	http.HandleFunc("/", home)
	http.HandleFunc("post", post)
	log.Println(http.ListenAndServe("localhost:1234", nil))
}

var users []string

func home(w http.ResponseWriter, _ *http.Request) {
	if err := homePage.Execute(w, users); err != nil {
		log.Printf("%v", err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v", err)
		return
	}
	fmt.Println(r.PostFormValue("buf"))
	fmt.Println(r.PostFormValue("Title"))
	fmt.Println(r.PostFormValue("name"))
}

var homePage = template.Must(template.New("home").Parse(
	`<html><body>
<form action="/post"method="post"><br/>
<input type='button'value='Title'id='but1'/>
<input type='button'value='Artist'id='but2'/>
<input type='button'value='Album'id='but3'/>
<input type='button'value='Year'id='but4'/>
<input type='button'value='Length'id='but5'/>
</form></body></html>
`))
