# HackSoc Web Application

The official website of [HackSoc](http://hacksoc.com/) from The University of Manchester. We also have a [Facebook](https://www.facebook.com/groups/HackSocManc/) and [LinkedIn](https://www.linkedin.com/company-beta/17980188/) page!

### Getting Started

Follow these instructions to run the website on your local machine.

### Prerequisites

#### Required
* Download and install the Go distribution from [here](https://golang.org/doc/install).

#### Optional
* Enabling reCAPTCHA requires a secret key and a site key. The first two steps of [this tutorial](https://webdesign.tutsplus.com/tutorials/how-to-integrate-no-captcha-recaptcha-in-your-website--cms-23024) explain how to get them.
* To enable SendGrid, create a [free account](https://sendgrid.com/free/) and get an API key.
* A MySQL database is needed by certain features (such as signing up). Run the queries from the ```sql```  folder to generate the required tables in a new database. To connect the website to the database, a data source name (DSN) must be provided. Read [this](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for details on how a DSN string is formatted.

If you choose to ignore these optional prerequisites, certain parts of the website may be deactivated or not work as expected.

### Installing

#### Download the Package
Install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") using the [go tool](http://golang.org/cmd/go/ "go command") from the command line:
```bash
$ go get github.com/hacksoc-manchester/www
```
Make sure [Git is installed](http://git-scm.com/downloads) on your machine and referenced by your system's ```PATH```.

#### Environment Variables
Linux/MacOS - Add the following to the `.env` file in your project root (create it if it does not exist, but do *not* commit it)
```bash
HTTP_PLATFORM_PORT="default port is 8080"
NOREPLY_EMAIL="noreply email used by the SendGrid API"
CONTACT_EMAIL="contact email used by the SendGrid API"
SYMMETRIC_KEY="personal symmetric key (can be of length 16, 24 or 32). keys are more secure"
RECAPTCHA_SITE_KEY="your reCAPTCHA site key"
RECAPTCHA_SECRET_KEY="your reCAPTCHA secret key"
MYSQL_CONNECTION_STRING="your MySQL connection string"
SENDGRID_API_KEY="your SendGrid API key"
FB_APP_ID="your Facebook app ID"
FB_SECRET="your Facebook secret key"
```
Windows - Add the same environment variables as above. You may have to create a file called `.env.` in Windows Explorer for it to proper register the `.env` filename.

#### Starting the Website
Go to the ```hacksoc-manchester/www``` folder and run the following go command in your shell:
```bash
$ go run app.go
```

### Running Tests

Go to the ```hacksoc-manchester/www``` folder and run the following go command in your shell to run the predefined tests:
```bash
$ go test ./...
```

### Authors

See the [CONTRIBUTORS](CONTRIBUTORS) file.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
