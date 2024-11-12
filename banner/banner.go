package banner

import (
	"fmt"
)

// prints the version message
const version = "v0.0.1"

func PrintVersion() {
	fmt.Printf("Current subdomainfuzz version %s\n", version)
}

// Prints the Colorful banner
func PrintBanner() {
	banner := `
                 __         __                          _         ____                
   _____ __  __ / /_   ____/ /____   ____ ___   ____ _ (_)____   / __/__  __ ____ ____
  / ___// / / // __ \ / __  // __ \ / __  __ \ / __  // // __ \ / /_ / / / //_  //_  /
 (__  )/ /_/ // /_/ // /_/ // /_/ // / / / / // /_/ // // / / // __// /_/ /  / /_ / /_
/____/ \__,_//_.___/ \__,_/ \____//_/ /_/ /_/ \__,_//_//_/ /_//_/   \__,_/  /___//___/
`
	fmt.Printf("%s\n%85s\n\n", banner, "Current subdomainfuzz version "+version)
}
