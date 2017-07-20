package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		x, err := strconv.Atoi(strings.Replace(r.URL.Path, "/", "", -1))
		if err != nil {
			fmt.Fprintln(w, `
			<!doctype html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport"
					  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Square Generator</title>
			</head>
			<body>
			<h2>Usage:</h2>
			<h4>Add a number to url path.</h4>
			<h4>Example:</h4><a href="/10">`+r.Host+`/10</a>
			</body>
			</html>`)
		} else {
			fmt.Fprintln(w, `
			<head>
				<meta charset="UTF-8">
				<meta name="viewport"
					  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Square Creator</title>
				<style>pre { line-height:0;}</style>
			</head>`)
			for i := 1; i <= x; i++ {
				io.WriteString(w, "<pre>")
				for j := 1; j <= x; j++ {
					if j == x {
						io.WriteString(w, "#")
					} else {
						io.WriteString(w, "# ")
					}
				}
				io.WriteString(w, "</pre>")
			}
		}
	})

	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		x, err := strconv.Atoi(strings.Replace(r.URL.Path, "/data/", "", -1))
		if err != nil {
			fmt.Fprintln(w, `
			<!doctype html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport"
					  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Data Generator</title>
			</head>
			<body>
			<h2>Usage:</h2>
			<h4>Add a number to url path.</h4>
			<h4>Example:</h4><a href="/data/10">`+r.Host+`/data/10</a>
			<h4>This will generate 10KB data</h4>
			</body>
			</html>`)
		} else {
			if x < 1 {
				io.WriteString(w, "Minimum 1KB")
				io.WriteString(w, `<h4>Example:</h4><a href="/data/10">`+r.Host+`/data/10</a>`)
			} else {
				fmt.Fprint(w, "<!doctype html><head><style>div{display: none;}</style></head><body><div>")

				for i := 1; i <= (x*1024)-86; i++ {
					io.WriteString(w, "#")
				}
			}
			io.WriteString(w, "</div></body>")
		}
	})

	http.ListenAndServe(":8080", nil)
}
