# how the tool works:

this tool can detect technologies of web applications and can even detect cms without forgetting their versions 
this tool is built with golang reputer for it's speed
we can analyze a single url or a file that contains several urls
the result in the terminal is made in the form of tables for easy reading; in parallel a file.xlsx is created and saves the same results as the terminal

to understand the relationships between the functions, you will find a photo of tree diagram : 

<img width="794" alt="Capture d’écran 2022-06-27 à 11 17 01" src="https://user-images.githubusercontent.com/107410271/175919219-e629ecbb-42fd-4b39-9625-458dcb0370fc.png">

## function correlation:

To begin with, this tool can take as input a url or a file.txt  that contains several urls.
In the case where we put a file in input:
The function <img width="38" alt="image" src="https://user-images.githubusercontent.com/107410271/175920866-d1f325d3-9b5e-427c-88de-773b05d1db54.png">
opens the file and reads it line by line, for each line we must have a precise url
For each URL:
the function checks if it is
### 1) a Domain Name:
The function adds “https” 
### 2) a url that has a valid format
The function leaves the url as it is
