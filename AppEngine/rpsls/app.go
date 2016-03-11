package rpsls

import (
    "fmt"
    //"github.com/julienschmidt/httprouter"
    "net/http"
    //"log"
    "text/template"
    "math/rand"
    "time"
)

// goapp deploy -application proud-outrider-106502 app.yaml

type Choice struct {
	HumanChoice, AiChoice, Death, Result string
}

var patterns = []string{"NULL", "Scissors","Paper", "Rock", "Lizard", "Spock"}

const doc = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>
	RPSLS
	</title>
	</head>
	<body>
	<h1>
	You chose {{.HumanChoice}}
	</h1>
	<h1>
	Computer chose {{.AiChoice}}
	</h1>
	<h1>
	{{.HumanChoice}} {{.Death}} {{.AiChoice}}
	</h1>
	<h1>
	{{.Result}}
	</h1>
	<h2>
	<FORM><INPUT Type="button" VALUE="Back" onClick="history.go(-1);return true;"></FORM>
	</body>
	</html>
`


func init() {
    //router := httprouter.New()
    //fs := http.FileServer(http.Dir("static"))

    http.HandleFunc("/", Index)
    http.HandleFunc("/rock", Rock)
    http.HandleFunc("/paper", Paper)
    http.HandleFunc("/scissors", Scissors)
    http.HandleFunc("/lizard", Lizard)
    http.HandleFunc("/spock", Spock)
    //http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

    //log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>
	RPSLS - Colin Adler
	</title>
	</head>
	<body>
	<h1>
	Welcome to Rock Paper Scissors Lizard Spock
	</h1>
	<h5>
	Created by Colin Adler
	</h5>
	<h2>
	Please click on a link to make your choice
	</h2>
	<h2>
	<a href="rock" class="btn btn-primary btn-lg" role="button">Rock</a>
	<a href="paper" class="btn btn-primary btn-lg" role="button">Paper</a>
	<a href="scissors" class="btn btn-primary btn-lg" role="button">Scissors</a>
	<a href="lizard" class="btn btn-primary btn-lg" role="button">Lizard</a>
	<a href="spock" class="btn btn-primary btn-lg" role="button">Spock</a>
	</h2>
	</body>
	</html>
	`)
}

func Rock(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("RPSLS").Parse(doc)
    if err == nil {
    	var p1, p2 int
    	var winner string
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// makes a rand instance
		p1 = 3
		p2 = r.Intn(5) + 1;
		// Judge
		result, death := judge(p1, p2, true)
		// Print
		switch result {
		case -1: winner = "Computer wins"
		case 0 : winner = "Tie"
		case 1, 2, 3, 4 : winner = "You win!"
	}
    	choice := Choice{patterns[p1], patterns[p2], death, winner}
    	tmpl.Execute(w, choice)
    	//time.Sleep(3000 * time.Millisecond)
    	//http.Redirect(w, r, "https://github.com", 301)
    }
}

func Paper(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("RPSLS").Parse(doc)
    if err == nil {
    	var p1, p2 int
    	var winner string
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// makes a rand instance
		p1 = 2
		p2 = r.Intn(5) + 1;
		// Judge
		result, death := judge(p1, p2, true)
		// Print
		switch result {
		case -1: winner = "Computer wins"
		case 0 : winner = "Tie"
		case 1, 2, 3, 4 : winner = "You win!"
	}
    	choice := Choice{patterns[p1], patterns[p2], death, winner}
    	tmpl.Execute(w, choice)
    }
}

func Scissors(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("RPSLS").Parse(doc)
    if err == nil {
    	var p1, p2 int
    	var winner string
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// makes a rand instance
		p1 = 1
		p2 = r.Intn(5) + 1;
		// Judge
		result, death := judge(p1, p2, true)
//		_ , death := judge(p1, p2, true)
		// Print
		switch result {
		case -1: winner = "Computer wins"
		case 0 : winner = "Tie"
		case 1, 2, 3, 4 : winner = "You win!"
	}
    	choice := Choice{patterns[p1], patterns[p2], death, winner}
    	tmpl.Execute(w, choice)
    }
}

func Lizard(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("RPSLS").Parse(doc)
    if err == nil {
    	var p1, p2 int
    	var winner string
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// makes a rand instance
		p1 = 4
		p2 = r.Intn(5) + 1;
		// Judge
		result, death := judge(p1, p2, true)
		// Print
		switch result {
		case -1: winner = "Computer wins"
		case 0 : winner = "Tie"
		case 1, 2, 3, 4 : winner = "You win!"
	}
    	choice := Choice{patterns[p1], patterns[p2], death, winner}
    	tmpl.Execute(w, choice)
    }
}

func Spock(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("RPSLS").Parse(doc)
    if err == nil {
    	var p1, p2 int
    	var winner string
		r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// makes a rand instance
		p1 = 5
		p2 = r.Intn(5) + 1;
		// Judge
		result, death := judge(p1, p2, true)
		// Print
		switch result {
		case -1: winner = "Computer wins"
		case 0 : winner = "Tie"
		case 1, 2, 3, 4 : winner = "You win!"
	}
    	choice := Choice{patterns[p1], patterns[p2], death, winner}
    	tmpl.Execute(w, choice)
    }
}

func judge(p1 int, p2 int, displayR bool) (int, string){
	
	var ret	int
//	var fullDeath string
	death := "breaks the code of"
	// determine result
	switch (p1 - p2){
	default: 		ret = -1	// p1 lose
	case 0:			ret = 0		// tie
	case -1:		ret = 1		// p1 win
		//fmt.Println("this is case -1")	//scissors -> paper | paper -> rock | rock -> lizard | lizard -> spock
	case -3:		ret = 2		// p1 win
		//fmt.Println("this is case -3") 	//scissors -> lizard | paper -> rock
	case 2:			ret = 3		// p1 win
		//fmt.Println("this is case -2")	//spock -> rock | rock -> scissors | lizard -> paper
	case 4:			ret = 4		// p1 win
		//fmt.Println("this is case -4") 	// spock -> scissors
	}
	// display result
	if displayR == true {
		switch (ret) {
		default:
			fmt.Println("rip code")			// congrats if u get this

		case -1:
			death = "is beaten by"			// lose

		case 0:
			death = "ties with"				// tie

		case 1: 		
			if p1 == 1 {
				death = "cuts"				// scissors -> paper
			}
			if p1 == 2 {
				death = "covers"			// paper -> rock
			}
			if p1 == 3 {
				death = "crushes"			// rock -> lizard
			}
			if p1 == 4 {
				death = "disproves"			// lizard -> spock
			}

		case 2: 		
			if p1 == 1 {
				death = "decapitate"		// scissors -> lizard
			}
			if p1 == 2 {
				death = "covers"			// paper -> rock
			}

		case 3: 		
			if p1 == 5 {
				death = "vaporizes"			// spock -> rock
			}
			if p1 == 3 {
				death = "crushes"			// rock -> scissors
			}
			if p1 == 4 {
				death = "eats"				// lizard -> paper
			}

		case 4:
			death = "smashes"				// spock -> scissors
		}
	}
	return ret, death
}

// cd /Users/colinadler/GitHub/GoApp