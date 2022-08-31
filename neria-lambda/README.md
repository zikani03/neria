neria-lambda
============

Simple AWS Lambda function that extracts named entities from articles. It uses the [prose](https://github.com/jdkato/prose) library for NER and allows you to specify the URL and a JQuery-like element selector to specify the elements to extract text from which uses [goquery](https://github.com/PuerkitoBio/goquery)

## Example Event / Request

### With Text from a URL 

```json
{
    "Url": "https://www.nyasatimes.com/chilima-says-malawi-is-a-best-investment-place-in-sadc-region-and-beyond/",
    "Selector": "#content div.nyasa-content",
	"Text": ""
}
```

### With Text you already have

You can also send the actual text to avoid the scraping process.

```json
{
    "Url": "",
    "Selector": "",
	"Text": "Some text ...."
}
```



## Example Response

The above request with text from a URL should give the following response:

```json
{
	"Entities": [
		{
			"EntityType": "GPE",
			"Name": "Malawi Vice President Salous Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "LOCATION",
			"Name": "Southern Africa Development Community SADC"
		},
		{
			"EntityType": "GPE",
			"Name": "Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Dubai"
		},
		{
			"EntityType": "GPE",
			"Name": "United Arab Emirates"
		},
		{
			"EntityType": "PERSON",
			"Name": "Lazarus Chakwera"
		},
		{
			"EntityType": "GPE",
			"Name": "Dubai Expo"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Dubai"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi Veep"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi National Day Saturday"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Centre"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi Investment"
		},
		{
			"EntityType": "PERSON",
			"Name": "Trade Centre"
		},
		{
			"EntityType": "GPE",
			"Name": "MITC"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "PERSON",
			"Name": "MITC"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Lands Act"
		},
		{
			"EntityType": "ORGANIZATION",
			"Name": "Ministry"
		},
		{
			"EntityType": "GPE",
			"Name": "Lands"
		},
		{
			"EntityType": "PERSON",
			"Name": "MITC"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "PERSON",
			"Name": "MITC"
		},
		{
			"EntityType": "PERSON",
			"Name": "/>Malawi Veep"
		},
		{
			"EntityType": "PERSON",
			"Name": "Mary Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Dubai"
		},
		{
			"EntityType": "GPE",
			"Name": "United Arab Emirates"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "PERSON",
			"Name": "Online Visa"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Africa"
		},
		{
			"EntityType": "GPE",
			"Name": "Veep"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "UAE"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Emirates Airlines"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Africa"
		},
		{
			"EntityType": "GPE",
			"Name": "MITC"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "GPE",
			"Name": "Dubai Expo"
		},
		{
			"EntityType": "PERSON",
			"Name": "MITC"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "PERSON",
			"Name": "Assistant Public"
		},
		{
			"EntityType": "PERSON",
			"Name": "Nellie Mononga"
		},
		{
			"EntityType": "GPE",
			"Name": "Malawi"
		},
		{
			"EntityType": "PERSON",
			"Name": "Chilima"
		},
		{
			"EntityType": "GPE",
			"Name": "Dubai"
		},
		{
			"EntityType": "PERSON",
			"Name": "Lazarus Chakwera"
		},
		{
			"EntityType": "ORGANIZATION",
			"Name": "COMESA Heads"
		},
		{
			"EntityType": "GPE",
			"Name": "State"
		},
		{
			"EntityType": "GPE",
			"Name": "Egypt"
		},
		{
			"EntityType": "PERSON",
			"Name": "Follow"
		},
		{
			"EntityType": "PERSON",
			"Name": "Subscribe Nyasa"
		}
	]
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