SP Worker
==========

This worker provides all the API for SP

for more info, read the docs below
https://github.com/ramvasanth/sp/blob/master/docs.MD

## Development
install Go language compiler and mysql server using brew or linux installer such as apt-get or yum
#### for windows
```
download golang from https://golang.org/dl/
```
#### for linux flavours
```sh
$ sudo apt-get install golang-go
$ sudo apt-get install mysqld

```
You need to set up your [GOPATH](http://golang.org/doc/code.html#GOPATH).
For example, you can put the following lines in ~/.bash_profile or ~/.zshrc depends on the shell type

```
export GOPATH="$HOME/gocode"
export PATH=$PATH:$GOPATH/bin
```
create the following folder structure
```sh
cd ~
mkdir -p ~/gocode/src/github.com/ramvasanth/sp
```
so the folder structure would look like ~/gocode/src/github.com/ramvasanth/sp

Now cd into github.com/ramvasanth/sp folder and clone the repo
```sh
cd ~/gocode/src/github.com/ramvasanth/sp
git clone https://github.com/ramvasanth/sp
```
### Build and Start

```sh
$ ./script/build # to build the app
$ ./script/start # to start the app
```
###  Build and Start for windows
```sh
$ cp .env.example .env # change it match your local setting
$ ./script/build.bat # to build the app
$ ./script/start.bat # to start the app
```

### Test
```sh
$ ./script/test # to test the app
```