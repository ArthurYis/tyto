### 
# @Description: 
 # @Author: Arthur
 # @Date: 2019-08-16 16:09:28
 # @LastEditTime: 2019-08-29 13:55:28
 # @LastEditors: Arthur
 ###

#  linux 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tyto main.go
#  mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
#  windows
# CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go