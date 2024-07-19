# goSmuggling
Series of HTTP Request Smuggling scenario collected in this repository. 

## Apache ReverseProxy Confs Steps/Checklist

1. Create configuration file on ```/etc/apache2/sites-available```.
2. Make sure that port you are using has listening status on  ```/etc/apache2/ports.conf``` file ``` "Listen 1920"```.
3. UFW disabled or allowed reverse proxy port
```ufw allow 1920/tcp comment "smuggling reverseProxy"```.
4. Apache proxy modules has been enabled successfully ```sudo a2enmod proxy proxy_http proxy_balancer lbmethod_byrequests```.
5. Site/Configuration has been enabled by ```a2ensite```.
6. DNS records has been defined on ```/etc/hosts``` both your client and reverse proxy server.
7. Do not forget this commands for T-Shoot "Always restart apache"
```sh
sudo systemctl restart apache2
sudo systemctl status apache2
apachectl -M
apachectl -S
tail -f /var/log/apache2/{LOGFILES}
```