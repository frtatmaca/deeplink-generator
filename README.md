# DeepLink
Deep Link

## About

This project is a challenge for Trendyol.

Simply, i designed and coded deeplink creator in the project.

Besides, our database(MySql) is also stored to links request in order not to lose the data.

This project was developed by Go Language.


## Usage

You can directly run server.go file to use program.

Four endpoints are defined to use, these are;

```Shell
To Generate a new deeplink

POST http://localhost:8000/urltodeeplink
{
	"link":"https://www.trendyol.com/casio/saat-p-1925865?boutiqueId=439892&merchantId=105064"
}
```

```Shell
To Generate a new link

POST http://localhost:8000/deeplinktourl
{
	"link":"ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064"
}

```

## Run

You can run the project via command:

```Shell
go run server.go
```

## Docker

You can dockerize the project via command:

```Shell
docker build -t deeplink .
```

You can run the project via command:
```Shell
docker run -v /logs:/logs -p 8000:8000 -it deeplink
```

## Testing

This project has integration tests in specs/linkendpoints_test.go file.

You can test the code via this file.

## Swagger

http://localhost:8000/swagger/index.html

## About me

I am Fırat Atmaca.

I have been working on software projects since 2013.

You can contact with me via [Linkedin](https://www.linkedin.com/in/firat-atmaca-469b2769/)

