# Go Graphql Starter Kit

A boilerplate for golang Graphql starter project. Aim for this project is to help get started with new projects.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

You need to have a postgres instance running and dep should be installed for installing dependencies

### Installing

A step by step series of examples that tell you have to get a development env running

clone the repo

```
git clone https://github.com/akmittal/go-graphql-starter.git
```

Go to cloned directory and Run dep ensure

```
dep ensure
```

Change the database connection for postgres in config/db.go


Run project

```
go run main.go
```

## Running the tests

Need to add tests


## Built With

* [graphql-go](https://github.com/neelance/graphql-go) - Graphql Library
* [ozzo-validation](https:/github.com/go-ozzo/ozzo-validation) - Validation Library
* [pq](https://rgithub.com/lib/pq) - Postgres Driver
* [bcrypt](https://golang.org/x/crypto/bcrypt) - Bcrypt for JWT tokens
* [sqlx](https:/github.com/jmoiron/sqlx) - sql helper library

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/akmittal/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/akmittal/go-graphql-starter/tags). 

## Authors

* **Amit Mittal** - *Initial work* - [Akmittal](https://github.com/Akmittal)

See also the list of [contributors](https://github.com/akmittal/go-graphql-starter/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

