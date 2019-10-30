# batteriesnotincluded
Building a webstack in Go

# install rbenv
	brew install rbenv
	rbenv init
	curl -fsSL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-doctor | bash
	rbenv install 2.6.1
	gem install thin rails

# install go
	https://golang.org/dl/
	download go 1.13.1 and run the package to install
	go version

# run benchmarks

Benchmarks originally run on 2.9ghz 8gb late 2015, Macbook Pro 13


## Ruby 2.6.1
		ruby thin.rb   		- 3309.26 [#/sec] (mean)	- ab -k -n 10000 -r http://127.0.0.1:8080/
		ruby http.rb 			- 4795.89 [#/sec] (mean)  - ab -k -n 10000 -r http://127.0.0.1:8081/
	 	puma -t 1 puma.ru	- 11490.19 [#/sec] (mean) - ab -k -n 10000 -r http://127.0.0.1:9292/

## Go 1.13.1
		Go run main.go  	- 11976.32 [#/sec] (mean) - ab -k -n 10000 -r http://127.0.0.1:8082/

**BE CAREFUL RUNNING YOUR OWN TESTS
	Notice how I run these on different ports
	Requests per second:    704.66 [#/sec] (mean) (GO)
	https://www.ncftp.com/ncftpd/doc/misc/ephemeral_ports.html
