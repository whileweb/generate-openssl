## Generate self-signed OpenSSL certificates.

1. Build application
```
go build -o generate_ssl main.go
```
2. Edit alt_names.ext file: add necessary local names as DNS records. Also add these names to /etc/hosts:
```
127.0.0.1 app.local admin.app.local
```
3. Generate certificates (will be created in the current directory)
```
./generate_ssl -h

Usage of ./generate_ssl:
  -C string
        a string (default "UA")
  -CN string
        a string (default "localhost.local")
  -L string
        a string (default "Kyiv")
  -O string
        a string (default "Localhost-Certificates")
  -ST string
        a string (default "Kyiv")
  -crtpath string
        a string (default "certs")
  -days uint
        a uint16 (default 1024)
  -extpath string
        a string (default "alt_names.ext")
```
4. Then import <strong>RootCA.crt</strong> to <strong>chrome://settings/certificates?search=https</strong> in tab <strong>Authorities</strong>
![alt text](https://raw.githubusercontent.com/oleksiivelychko/generate-openssl/master/screen.png)
