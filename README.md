# HackSoc Web Application

The official website of the [Hack Society](http://hacksoc.com/) from The University of Manchester. You can also check us on [Facebook](https://www.facebook.com/groups/HackSocManc/)

### Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

#### Primary
* The Go distribution [here](https://golang.org/doc/install).

#### Optional
* ReCaptcha Google: get your Secret and Site key [here](https://webdesign.tutsplus.com/tutorials/how-to-integrate-no-captcha-recaptcha-in-your-website--cms-23024), follow all steps until step 3.
* SendGrid: create Free account [here](https://sendgrid.com/free/) and get your API key [here](https://sendgrid.com/docs/Classroom/Basics/API/what_is_my_api_key.html).
* MySql Database: for MySql connection string, check [here](https://github.com/go-sql-driver/mysql#dsn-data-source-name).

The optional prerequisites let you have a mailing system, a database and a Google ReCaptcha, if you decide you only want the primary ones, the website will run without problems but with some warnings, related to these prerequisites.

### Installing

#### Download the package
Simple install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") with the [go tool](http://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/hacksoc-manchester/www
```
Make sure [Git is installed](http://git-scm.com/downloads) on your machine and in your system's `PATH`.

#### Environment Variables
Linux - Add the following to your relevant shell profile (a relog is required):
```bash
export HTTP_PLATFORM_PORT="Default Port is 8080"
export NOREPLY_EMAIL="noreply email used by sendgrip api"
export CONTACT_EMAIL="contact email used by sendgrip api"
export SYMMETRIC_KEY="personal symm key, size 16, 24, 32. Longer = better security"
export RECAPTCHA_SITE_KEY="your ReCaptcha Site Key"
export RECAPTCHA_SECRET_KEY="your ReCaptcha Secret Key"
export MYSQL_CONNECTION_STRING="your mysql connection string"
export SENDGRID_API_KEY="your sendgrid api key"
```
Windows - Add the same Environment Variables as above, check [here](http://www.computerhope.com/issues/ch000549.htm) for help.

#### Running the app
Go to hacksoc-manchester/www folder and run the following go command in your shell:
```bash
$ go run app.go
```

### Running the Tests

Go to hacksoc-manchester/www folder and run the following go command in your shell to run the predefined tests:
```bash
$ go test ./...
```

### Authors

* **Andrei Muntean** - *Initial work* - [andreimuntean](https://github.com/andreimuntean)

See also the list of [contributors](https://github.com/hacksoc-manchester/www/graphs/contributors) who participated in this project.

### License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
