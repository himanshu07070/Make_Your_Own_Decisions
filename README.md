# Make_Your_Own_Decisions

[Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) is a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

- Turn to page 30 to go up  ladder.
- Turn to page 95 to venture down the dark passage.

The goal of this project is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc). Stories will be provided via a JSON file.

Major things done :
1. I Used the `html/template` package to create  HTML pages. 
2. Created an `http.Handler` to handle the web requests instead of a handler function.
3. Used the `encoding/json` package to decode the JSON file. 
