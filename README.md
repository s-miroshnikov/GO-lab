## GO-lab

This is testing project for learning GO language. <br>
Most of them are just copied tutorials from the Tour of Go. 

### web-gin
Simple web-server based on gin-gonic/gin, pgx v4 driver and alexeyco/pig. <br>

<p>
In this example I added PostgreSQL database 'library' with such tables: <br>
Magazines, Articles, article_types, author.

This script creates connection to PostgreSQL database with DSN parameters <br>
<code> postgres://reader1:Password1@db/library </code> <br>
receives data from tables with SELECT query and populates structures. <br>
After that we can see our data in the generated HTML page (template is /templates/library.tmpl) on the web server (port: 8080).
</p>