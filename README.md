# how the tool works:

this tool can detect technologies of web applications and can even detect cms without forgetting their versions 
this tool is built with golang reputer for it's speed
we can analyze a single url or a file that contains several urls
the result in the terminal is made in the form of tables for easy reading; in parallel a file.xlsx is created and saves the same results as the terminal

to understand the relationships between the functions, you will find a photo of tree diagram : 

<img width="794" alt="Capture d’écran 2022-06-27 à 11 17 01" src="https://user-images.githubusercontent.com/107410271/175919219-e629ecbb-42fd-4b39-9625-458dcb0370fc.png">

## function correlation:

To begin with, this tool can take as input a url or a .txt file that contains several urls.
In the case where we put a file in input:
The function <img width="38" alt="image" src="https://user-images.githubusercontent.com/107410271/175920866-d1f325d3-9b5e-427c-88de-773b05d1db54.png">
opens the file and reads it line by line, for each line we must have a precise url
For each URL:
the function checks if it is
1) a Domain Name:
The function adds “https” or “http” (it all depends on the url)
2) a url that has a valid format
The function leaves the url as it is

The function gives us the Domain Name of the url
According to the Domain Name, the function gives us the IP address of the url
The function takes as input the url of the function
And make sure that the url is valid and does not give an error (if yes return true otherwise return false)
The function detects the title of the web page
If the function return true:



The function reports the value of some specific header (if they exist) that we will use later:
• The name of the “server”
• The value of "X-POWERED-BY", this value tells us the language used in the web application
• The value of "X-ASPNET-VERSION", this value tells us the version of the asp.net language, we deduce the language
The function returns all the code of the page
The function takes the code of the page returned by and first searches all the urls of the source code, detects that the urls which contain the Domain Name of the url, detects if these selected urls contain "PHP" in the directory of the 'url.
The function takes the code of the page returned by and first searches all the urls of the source code, detects that the urls which contain the Domain Name of the url, detects if these selected urls contain "ASPX" in the directory of the 'url.
The function returns true if the function returns for example “X-POWERED-BY =PHP” or if the r function returns true
The function returns true if the function returns either for example "X-POWERED-BY =ASP" or the value of "X-ASPNET-VERSION" is not null or if the function returns true



The function takes the page code returned by
And first searches all the urls of the source code, detects that the urls which contain the Domain Name of the url, detects if these selected urls contain either “wp-…” or “drupal” or “joomla” or “magento” either “dnn_” for DotNetNuke or “DotNetNuke” or “SharePoint” or “spip.php” for spip or “SharePoint”
If yes it returns a variable of type string containing the name of the string found
Detected
It checks if the cms names that I mentioned exist in the [content] tag (we can at the same time find the cms version)

To check the current version of Magento 2 store is to add /magento_version at the end of the URL. This will display the current version of a site using Magento.
To check the current version of the drupal store is to add /CHANGELOG.txt to the end of the URL. This will display the current version of a site using drupal.

Checks if the SPIP CMS is used by checking if the [class] tag contains "''spip_''
Checks if the Magento CMS is used by checking if the [type] tag contains "''magento-init''
Checks if the DotNetNuke CMS is used by checking if the [id] tag contains "''dnn_''


“wp-…”: wp-content/wp-includes/wp-admin

• If the function returns “DotNetNuke” then the function Returns true
• If the function returns “wp-…” then the function Returns true
• If the function returns “drupal” then the function Returns true
• If the function returns "joomla" then the function Returns true
• If the function returns “magento” then the function Returns true


• If the function returns “SharePoint” then the function Returns true
• If the function returns "SPIP" then the


function returns true

We can also save our result of the url analysis in an Excel file automatically
