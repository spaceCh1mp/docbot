#   DOCBOT

Update status for the #100daysofcode challenge via twitter API endpoint. 

>       go run main.go -name=fileName

                        Or

>       go build

>       ./docbot -name=fileName

Flags to be set are

*    dir
        - Program defaults to ./log/ when no file path is specified
    
*    name
        - No value is set for the name flag and must be specified
    
*    rel
        - file path is relative by default, set this flag to false to set a path from your home directory

Example 

For file in $HOME/Documents/log/today.txt

    > ./docbot -dir=Documents/log -name=today.txt -relative=false

* flags are case sensitive i.e (Documents != documents)

Twitter access token is stored in a json file in the main directory, edit token.template.json and rename as
        
        token.json