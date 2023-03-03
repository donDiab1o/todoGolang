# todoGolang
ToDo List, using Golang

Frontend : HTML, JS

Backend: Golang

BD: PostgreSQL
Web-Server: Nginx 


Setup instructions:
Services which we need to use this app: Nginx, PostgreSQL(in my case in docker)

Create and configure ur Nginx with locale folder of app, like this:


server {
	listen "URPORT" default_server;
	listen [::]:"URPORT" default_server;
	
	root /var/www/todo/client; (ur destination addr)

	index index.html;

	server_name _;

	location / {
		try_files $uri $uri/ =404;
		add_header Last-Modified $date_gmt;
        	add_header Cache-Control 'no-store, no-cache';
        	if_modified_since off;
        	expires off;
        	etag off;
	}
	add_header 'Access-Control-Allow-Origin' 'origin-list';
}



So u can start nginx and check in browser localhost on ur port

Also we need PostgreSQL such as Data Base, i use postgresql in docker container

Create data base "todo" with relevant table, in table we would have 2 columns: ID and Tasks
Ur database config u need to set up in config file, template of config file u can see in repository.
Thats all.





