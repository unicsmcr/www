# HackSoc Web Application

The official website of [HackSoc](http://hacksoc.com/) from The University of Manchester. We also have a [Facebook](https://www.facebook.com/groups/HackSocManc/) and [LinkedIn](https://www.linkedin.com/company-beta/17980188/) page!

### Getting Started

Follow these instructions to run the website on your local machine.

### Prerequisites

#### Required
* The Go distribution [here](https://golang.org/doc/install).

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
Linux - Add the following to your relevant shell profile (a relog is required):
```bash
export HTTP_PLATFORM_PORT="default port is 8080"
export NOREPLY_EMAIL="noreply email used by the SendGrid API"
export CONTACT_EMAIL="contact email used by the SendGrid API"
export SYMMETRIC_KEY="personal symmetric key (can be of length 16, 24 or 32). Longer keys are more secure"
export RECAPTCHA_SITE_KEY="your reCAPTCHA site key"
export RECAPTCHA_SECRET_KEY="your reCAPTCHA secret key"
export MYSQL_CONNECTION_STRING="your MySQL connection string"
export SENDGRID_API_KEY="your SendGrid API key"
```
Windows - Add the same environment variables as above. Check [here](http://www.computerhope.com/issues/ch000549.htm) for help.

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

See the list of contributors [here](https://github.com/hacksoc-manchester/www/graphs/contributors).

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
