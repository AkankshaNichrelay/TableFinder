# TableFinder
A simulation of OpenTable app to browse and book tables at restaurants. (in the making)

## Running Locally
* Add your environment variables in config.env
* Run unit tests
    * All packages
        * make docker-test
    * A single package
        * make docker-tests PKG=[path to package]
* Running Docker Compose for MySQL and Redis
    * make docker-up
    * make docker-down
* Build
    * make docker-build
* Run 
    * make docker-run
* Postman
    * Set the tablefinder_url to the local url of app. Example: localhost:3000

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/3563c2780f7b7f5593c5?action=collection%2Fimport#?env%5Blocal%5D=W3sia2V5IjoidGFibGVmaW5kZXJfdXJsIiwidmFsdWUiOiJsb2NhbGhvc3Q6MzAwMCIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0Iiwic2Vzc2lvblZhbHVlIjoibG9jYWxob3N0OjMwMDAiLCJzZXNzaW9uSW5kZXgiOjB9XQ==)

