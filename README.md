## Generate local OpenSSL certificates.

1. Set script executable
```
chmod +x generate.sh
```
2. Edit alt_names.ext file: add necessary local names as DNS records. Also add these names to /etc/hosts like:
```
127.0.0.0 app.local admin.app.local
```
3. Edit path to store generated certificates. Change <strong>CERTSPATH</strong> environment variable value.
```
CERTSPATH=$HOME/Documents/test
```
4. Generate certificates. Also, you can change others values like CommonName or CityName etc.
```
./generate.sh
```
![alt text](https://raw.githubusercontent.com/oleksiivelychko/generate-openssl/master/screens/screen_1.png)
5. Then import <strong>RootCA.crt</strong> to <strong>chrome://settings/certificates?search=https</strong> in tab <strong>Authorities</strong>
![alt text](https://raw.githubusercontent.com/oleksiivelychko/generate-openssl/master/screens/screen_2.png)