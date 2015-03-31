# go-spa-kit [![wercker status](https://app.wercker.com/status/1b40701f67b072d03c5a3d174c46f8ab/m "wercker status")](https://app.wercker.com/project/bykey/1b40701f67b072d03c5a3d174c46f8ab)
###The ultimate component backend for single page applications written in go

####General

The goal of go-spa-kit is providing a construction kit for single page applications that runs inside a docker container.
#####Features:
* Webserver
* Adminbackend
* Plugin Manager

####Setup
####Usage
####Development
####Changelog


Adminbackend
Route: /admin


#Thoughts

Plugins

* Route /plugins/<plugin-name>
* Route Scoping <%= PLUGIN_URL%>

Need:
* getResponse(): string, bool
* getAdminResponse(): string, bool
* getName(): string
* setCofig(string): err
* init(): err
* getMetaData(string): err
