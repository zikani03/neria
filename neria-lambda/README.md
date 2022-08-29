neria-lambda
============

Simple AWS Lambda function that extracts named entities from articles. It uses the [prose](https://github.com/jdkato/prose) library for NER and allows you to specify the URL and a JQuery-like element selector to specify the elements to extract text from which uses [goquery](https://github.com/PuerkitoBio/goquery)

## Example Event / Request

```json
{
    "url": "https://www.nyasatimes.com/chilima-says-malawi-is-a-best-investment-place-in-sadc-region-and-beyond/",
    "selector": "#content div.nyasa-content"
}
```

## Building 

```sh
$ git clone https://github.com/zikani03/neria-lambda

$ cd neria-lambda

# Remember to build your handler executable for Linux!
$ GOOS=linux GOARCH=amd64 go build -o main main.go

$ zip main.zip main
```

**On Windws with Powershell**: Run `build.ps1`


---

Copyright (c) Zikani Nyirenda Mwase 