Example text based game using Golang objects.

## To Do list

- Build menu screen and display
    - This should include
        - Available actions
        - Character Health
        - Opponent Health (If relevant)

- Build display methods
    - Configuration of displays for ASCII/Art above menu
    - Scaling for multiple window sizes
    - Stored ASCII art(?)

- Story Development
    - What locations is the character a part of
    - Where does the character explore? 
    - Why?

- Action Development
    - Location specific actions
    - Inventory management
    - 

- Deployment tooling
    - API integration for web services (`api.go`)
    - CICD structure out of Github
    - Containerizing deployable resources


## Development

Main structure is within the `worldMain.go` file. 

To initiate a run, use `go build .` and `go run .`

