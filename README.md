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
the function <img width="34" alt="image" src="https://user-images.githubusercontent.com/107410271/175921702-1b453a29-e2fe-4ac3-8fdd-8c3d18ffd4e1.png">
checks if it is
### 1) a Domain Name:
The function adds “https” 
### 2) a url that has a valid format
The function leaves the url as it is

The function <img width="39" alt="image" src="https://user-images.githubusercontent.com/107410271/175921895-8f9d05dc-b1bc-4b60-bde1-e18b1038ef89.png">
gives us the Domain Name of the url
According to the Domain Name, the function <img width="34" alt="image" src="https://user-images.githubusercontent.com/107410271/175921940-17825a79-6ffc-4470-9f5b-6c178e367775.png">
gives us the IP address of the url

The function <img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175922268-3ecadb76-d87d-4849-8796-c5798e3bc3b6.png">
reports the value of some specific header (if they exist) that we will use later:
##### • The name of the “server”
##### • The value of "X-POWERED-BY", this value tells us the language used in the web application
##### • The value of “X-ASPNET-VERSION”, this value specifies the version of the asp.net language, we deduce the language

The function <img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175922780-906bedbe-491f-4ae0-a5f2-c0f18ffa0558.png">
first searches all the urls of the source code, detects that the urls which contain the Domain Name of the url, detects if these selected urls contain "PHP" in the directory of the 'url.

The function <img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175923202-a55c628e-d72d-4044-94b4-2a3e75d204c1.png">
 first searches all the urls of the source code, detects that the urls which contain the Domain Name of the url, detects if these selected urls contain "ASPX" in the directory of the 'url.

The function <img width="34" alt="image" src="https://user-images.githubusercontent.com/107410271/175923444-2a594888-37c0-4f1c-81dd-18e14a620804.png">
returns true if the function <img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175923464-1e0d83b1-8728-497f-936b-0034bcb18019.png">
returns for example "X-POWERED-BY =PHP" or if the  function  returns true
returns true if the function returns for example "X-POWERED-BY =PHP" or if the  function <img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175923569-096093fc-6fb2-47cf-a94d-78b82746478d.png">
returns true

The function <img width="34" alt="image" src="https://user-images.githubusercontent.com/107410271/175923712-728cc7de-e47d-4281-a5e3-442358a62d9c.png">
returns true if the function<img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175923735-05872cdd-a9cd-49fd-9afd-917b6b314720.png">
 returns either for example "X-POWERED-BY =ASP" or the value of "X-ASPNET-VERSION" is not null or if the function returns true
returns true if the function returns for example "X-POWERED-BY =PHP" or if the  function <img width="30" alt="image" src="https://user-images.githubusercontent.com/107410271/175923569-096093fc-6fb2-47cf-a94d-78b82746478d.png">
