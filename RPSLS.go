/*

Rules:
Scissors cut Paper
Paper covers Rock
Rock crushes Lizard
Lizard poisons Spock
Spock smashes Scissors
Scissors decapitate Lizard
Lizard eats Paper
Paper disproves Spock
Spock vaporizes Rock
Rock crushes Scissors

 _____       _ _          ___      _ _            ______                         _         
/  __ \     | (_)        / _ \    | | |           | ___ \                       | |      _ 
| /  \/ ___ | |_ _ __   / /_\ \ __| | | ___ _ __  | |_/ / __ ___  ___  ___ _ __ | |_ ___(_)
| |    / _ \| | | '_ \  |  _  |/ _` | |/ _ \ '__| |  __/ '__/ _ \/ __|/ _ \ '_ \| __/ __|  
| \__/\ (_) | | | | | | | | | | (_| | |  __/ |    | |  | | |  __/\__ \  __/ | | | |_\__ \_ 
 \____/\___/|_|_|_| |_| \_| |_/\__,_|_|\___|_|    \_|  |_|  \___||___/\___|_| |_|\__|___(_)
______           _     ______                       _____      _                          _     _                  _   _____                  _    
| ___ \         | |    | ___ \                     /  ___|    (_)                        | |   (_)                | | /  ___|                | |   
| |_/ /___   ___| | __ | |_/ /_ _ _ __   ___ _ __  \ `--.  ___ _ ___ ___  ___  _ __ ___  | |    _ ______ _ _ __ __| | \ `--. _ __   ___   ___| | __
|    // _ \ / __| |/ / |  __/ _` | '_ \ / _ \ '__|  `--. \/ __| / __/ __|/ _ \| '__/ __| | |   | |_  / _` | '__/ _` |  `--. \ '_ \ / _ \ / __| |/ /
| |\ \ (_) | (__|   <  | | | (_| | |_) |  __/ |    /\__/ / (__| \__ \__ \ (_) | |  \__ \ | |___| |/ / (_| | | | (_| | /\__/ / |_) | (_) | (__|   < 
\_| \_\___/ \___|_|\_\ \_|  \__,_| .__/ \___|_|    \____/ \___|_|___/___/\___/|_|  |___/ \_____/_/___\__,_|_|  \__,_| \____/| .__/ \___/ \___|_|\_\
                                 | |                                                                                        | |                    
                                 |_|                                                                                        |_|                    
*/

package main

import ("fmt"
        "math/rand"
        "time"
        //"net/http"
)
var patterns = []string{"NULL", "Scissors","Paper", "Rock", "Lizard", "Spock"}

func main(){
        done := false
        for done == false{
        fmt.Println("\n\tWelcome to RPSLS")
        fmt.Println("Please choose one of the following:\n")
        fmt.Println("1.\tPlay")
        fmt.Println("0.\tExit Game\n")
        fmt.Print("Please select a game mode: ")
        var mode int;
        fmt.Scanf("%d", &mode);
        switch mode {
            default: fmt.Println("\nError: Please choose 0 or 1\n\n")
            case 0: done = true			// bie bie
            case 1: ai()
        }
    }
}
func printSelections(){
    fmt.Println("\n===========================")
    fmt.Println("\tSelections")
    fmt.Println("===========================\n")
    fmt.Println("1. Scissors")
    fmt.Println("2. Paper")
    fmt.Println("3. Rock")
    fmt.Println("4. Lizard")
    fmt.Println("5. Spock\n")
}
func ai(){
    var p1, p2, result int
    r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))	// makes a rand instance
    // ask for input
    printSelections()
    fmt.Print("Please enter your choice: ")
    fmt.Scanf("%d", &p1)

    n, err := fmt.Scanf("%f\n", &inputSquare)
    if err != nil || n != 1 {
        // handle invalid input
        fmt.Println(n, err)
        return
    }
	
    // make ai's choice
    p2 = r.Intn(5) + 1;
    fmt.Println("\nThe computer's choice is", patterns[p2], "\n")
    // Judge
    result = judge(p1, p2, true)
    // Print
    switch result {
    case -1: fmt.Println("Computer wins\n")
    case 0 : fmt.Println("Tie\n")
    case 1, 2, 3, 4 : fmt.Println("You win!\n")
    }
}
func judge(p1 int, p2 int, displayR bool) int {
	
    var ret	int;
    death := "breaks the code of"
    // determine result
    switch (p1 - p2){
    default:        ret = -     // p1 lose
    case 0:         ret = 0     // tie
    case -1:        ret = 1     // p1 win | scissors -> paper | paper -> rock | rock -> lizard | lizard -> spock
    case -3:        ret = 2     // p1 win | scissors -> lizard | paper -> rock
    case 2:         ret = 3     // p1 win | spock -> rock | rock -> scissors | lizard -> paper
    case 4:         ret = 4     // p1 win | spock -> scissors
    }
    // determine how the computer dies
    if displayR == true {
        switch (ret) {
        default:
            fmt.Println("rip code")         // congrats if u get this
        case -1:
            death = "is beaten by"          // lose
        case 0:
            death = "ties with"				// tie
        case 1: 		
            if p1 == 1 {
                death = "cuts"              // scissors -> paper
            }
            if p1 == 2 {
                death = "covers"            // paper -> rock
            }
            if p1 == 3 {
                death = "crushes"           // rock -> lizard
            }
            if p1 == 5 {
                death = "disproves"         // lizard -> spock
            }
        case 2: 		
            if p1 == 1 {
                death = "decapitate"        // scissors -> lizard
            }
            if p1 == 2 {
                death = "covers"            // paper -> rock
            }
        case 3: 		
            if p1 == 5 {
                death = "vaporizes"         // spock -> rock
            }
            if p1 == 3 {
                death = "crushes"           // rock -> scissors
            }
            if p1 == 4 {
                death = "eats"              // lizard -> paper
            }
        case 4:
            death = "smashes"               // spock -> scissors
        }
    }
    fmt.Println(patterns[p1], death, patterns[p2], "\n")
    return ret
}