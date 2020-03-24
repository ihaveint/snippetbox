This a Golang web application which I'm developing while reading Alex Edwards [Let's go](https://lets-go.alexedwards.net/). There might be some differences from the book

## project URL
I try to maintain a working version of this project in the following link : [snippets](http://snippets.ihaveint.com)

### Prerequisites

- you need to install [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- also you need to install mysql
- create a user in mysql - you can follow the below lines in terminal( 'newuser' and 'user_password' can be replaced by anything you want)
```
mysql
mysql-> CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'user_password';
mysql-> CREATE DATABASE snippetbox;
mysql-> GRANT ALL PRIVILEGES ON snippetbox.* TO 'newuser'@'localhost';
mysql-> exit
```

### Installation

you can install the projcet by running the below command :
```
go get -u github.com/ihaveint/snippetbox/...
```

you need to add a self-signed TLS certificate :
```
mkdir project_path/tls
cd project_path/tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=lslocalhost
cd ..
```

#### for the rest of this guide, you shuold be in the 'snippetbox' folder

then you have to migrate the database : (enter the username and password that you created above)
```
migrate -database mysql://newuser:user_password@/snippetbox -path ./migrations/ up
```

### Running the project

you have to set envrionment variables for mysql usernamme and password :
```
export MySQLUser="username"
export MySQLPass="password"
```

before running the project, you have to activate extended globbing by running the below command :
```
shopt -s extglob
```

finally, you can run the project by :
```
go run cmd/web/!(*_test).go
```

#### NOTE :
because the TLS is self-signed, the first time you run the app in a browser, you might get a warning about security problems, based on how your browser works, there should be an option to proceed anyway, if you do this one time, later on you shouldn't get any warnings. for production, it's a good practice if your TLS is not self-signed. you can use [Let's Encrypt](https://letsencrypt.org) for example !

