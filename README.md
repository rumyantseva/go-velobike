# go-velobike

A Go client library for the API of the velobike.ru project.

API style based on the best practices of [google/go-github](https://github.com/google/go-github).

## Versioning

Current major version is 0.x.x.

There are [few warnings from go linter](https://github.com/rumyantseva/go-velobike/issues/2) here. I'd like to fix these warnings, but it'll break backward compatibility. So, these changes are planned for the next major version of the library.

If you use this library, please vendor it using a package management tool to not to have backward compatibility problems in the future. 

If you have any questions, feel free to [create an issue](https://github.com/rumyantseva/go-velobike/issues/new), I'll try to answer as soon as possible.

## Supported methods

* GET /profile
* POST /profile/authorize
* GET /ride/parkings
* GET /ride/history

## Roadmap

* Add electric bikes
* Add GET /content/news method
