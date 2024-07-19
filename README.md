# Git Hub API
This project is a Go application that extracts data from the GitHub API and saves it locally in JSON and CSV formats. It retrieves information about GitHub users, organizations, and repositories and provides a simple way to manage this data.

## Features
* Extract data from GitHub users, organizations, and repositories.
* Save extracted data in JSON format.
* Convert JSON data to CSV format for easier analysis.

### How to test it

Before running the main.go file, you need to set a environment variable with your own GITHUB token

* For Linux/MacOS: export GITHUB_TOKEN=your_token_here
* For Windows: set GITHUB_TOKEN=your_token_here

### TO DO
* Append new data to existing CSV files without overwriting.

## Author
Tito Muller (follow me on [LinkedIn](https://www.linkedin.com/in/joao-vittor-muller-99142b13a/))
