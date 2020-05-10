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
```
![alt text](https://raw.githubusercontent.com/oleksiivelychko/generate-openssl/master/screens/screen_1.png)

4. Then import <strong>RootCA.crt</strong> to <strong>chrome://settings/certificates?search=https</strong> in tab <strong>Authorities</strong>
![alt text](https://raw.githubusercontent.com/oleksiivelychko/generate-openssl/master/screens/screen_2.png)
