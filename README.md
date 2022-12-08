# Advent Of Code 2022 - Go

## Folder Structure
The folder structure is as follows:

    .
    ├── .devcontainer      # contains the devcontainer.json configuration file
    ├── DayX               # Folder containing the source code for a given X day
    |    ├── Input      # The folder containing the input data
    |    └── main.go    # The GO program solving the given Advent Calendar day's exercise
    ├── utils              # Folder containing the common utility functions
    |    ├── utils.go           # The utils GO package
    └── README.md

## How To Run The Code
In order to run the different Advent Of Code 2022 solutions, follows these steps:
1. From the root directory, run the devcontainer solution, or otherwise install the GO binaries in your current system in order to be able to run GO programs.
2. Change directory into one of the days' folders.
3. Run `go run main.go`.