mkdir -p $GOPATH/src/github.com/afnarqui
cd $GOPATH/src/github.com/afnarqui
dep init
docker run --name godep -P --volumes-from volumen -v c:/Users/anaranjo/proyectos:/local -it afnarqui/godep bash
docker inspect godep | grep -i tcp
dep ensure --add github.com/joho/godotenv
curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash - && sudo apt-get install --yes nodejs
sudo npm install db-migrate db-migrate-mysql -g

https://www.youtube.com/watch?v=Bv60lebPu24&list=PLTxFJWe_410yjXmAE90-eUWs9xxIyWNU8&index=29

https://github.com/go-chi/chi
go get -u github.com/go-chi/chi

docker run --name gochi -p 8082:8082 -p 8083:8083 --volumes-from volumen -v c:/Users/afnarqui/proyectos;/local -it afnarqui/gochi bash
cd $GOPATH/src/github.com/afnarqui

chi go run main.go


https://www.youtube.com/watch?v=VAGodyl84OY

docker network create afn

docker exec -it 6aab ./cockroach sql --insecure

CREATE DATABASE account;
CREATE DATABASE image;

CREATE USER account_user WITH PASSWORD 'account_user';
CREATE USER image_user WITH PASSWORD 'quintero1.';

GRANT ALL ON DATABASE account TO account_user;
GRANT ALL ON DATABASE image TO image_user;


> git clone https://github.com/callistaenterprise/goblog.git
> git checkout P13


docker run --name postgressafn -e POSTGRES_PASSWORD=test -d postgres

docker exec -it postgressafn psql -U postgres

CREATE USER golang PASSWORD 'golang';

CREATE DATABASE gocrud OWNER golang

CREATE TABLE estudiantes (
     id SERIAL NOT NULL,
     NAME VARCHAR(50) NOT NULL,
     age SMALLINT NOT NULL,
     active BOOLEAN NOT NULL,
     create_at TIMESTAMP NOT NULL DEFAULT NOW(),
     update_at TIMESTAMP
);

dt
d estudiantes

go get github.com/lib/pq


docker network create -d bridge roachnet

docker run -d --name=roach1 --hostname=roach1 --net=roachnet -p 26257:26257 -p 8080:8080  -v c:/Users/afnarqui/proyectos:/local  coc
kroachdb/cockroach:v19.1.0 start --insecure

docker run -d \
--name=roach1 \
--hostname=roach1 \
--net=roachnet \
-p 26257:26257 -p 8080:8080  \
-v "c:/Users/afnarqui/proyectos:/cockroach/cockroach-data"  \
cockroachdb/cockroach:v19.1.0 start --insecure

docker run -d --name=roach1 --hostname=roach1 --net=roachnet -p 26257:26257 -p 8080:8080  -v c:/Users/afnarqui/proyectos:/local  cockroachdb/cockroach:v19.1.0 start --insecure

docker exec -it roach1 ./cockroach sql --insecure

CREATE DATABASE test;

CREATE TABLE test.cantidad (id INT PRIMARY KEY, balance DECIMAL);

INSERT INTO test.cantidad VALUES (1, 1000.50);

SELECT * FROM test.cantidad;

\q


docker run --name gochiv1 --link=roach1 --net=roachnet --volumes-from volumen -v c:/Users/afnarqui/proyectos:/local -it afnarqui/godep:v5 bash


https://github.com/afnarqui/godep-mysql-vue.git

go get github.com/lib/pq


wget -qO- https://binaries.cockroachdb.com/cockroach-v19.1.0.linux-amd64.tgz | tar  xvz

cp -i cockroach-v19.1.0.linux-amd64/cockroach /usr/local/bin

cockroach help
cockroach start --insecure

docker run --name godepv1 -p 8080:8080 --volumes-from volumen -v c:/Users/afnarqui/proyectos:/local -it afnarqui/godep:v5 bash

cd $GOPATH/src/github.com/afnarqui

git clone https://github.com/afnarqui/godep-mysql-vue.git

cd godep-mysql-vue
go get github.com/lib/pq

go run crud.go

wget -qO- https://binaries.cockroachdb.com/cockroach-v19.1.0.linux-amd64.tgz | tar  xvz

cp -i cockroach-v19.1.0.linux-amd64/cockroach /usr/local/bin

cockroach help
cockroach start --insecure

docker run --name gocockroa -p 8080:8080 -p 8081:8081 --volumes-from volumen -v c:/Users/afnarqui/proyectos:/local -it afnarqui/godep:v6 bash

cd $GOPATH/src/github.com/afnarqui

cockroach start --insecure
docker exec -it gocockroa bash
cd $GOPATH/src/github.com/afnarqui
cd godep-mysql-vue
go run crud.go

go get github.com/googollee/go-socket.io


<!-- Load required Bootstrap and BootstrapVue CSS -->
<link type="text/css" rel="stylesheet" href="//unpkg.com/bootstrap/dist/css/bootstrap.min.css" />
<link type="text/css" rel="stylesheet" href="//unpkg.com/bootstrap-vue@latest/dist/bootstrap-vue.min.css" />

<!-- Load polyfills to support older browsers -->
<script src="//polyfill.io/v3/polyfill.min.js?features=es2015%2CMutationObserver" crossorigin="anonymous"></script>

<!-- Load Vue followed by BootstrapVue -->
<script src="//unpkg.com/vue@latest/dist/vue.min.js"></script>
<script src="//unpkg.com/bootstrap-vue@latest/dist/bootstrap-vue.min.js"></script>


github.com/go-chi/chi/middleware

go get -u -v github.com/mattn/go-sqlite3

go build
./main -migrate

curl -X POST http://localhost:8081/notes -H 'Content-Type: application/json' -d '{"title": "Primera nota", "description": "Esta es una nota de prueba"}'   


curl -X GET http://localhost:8081/notes

curl -X PUT http://localhost:8080/notes -H 'Content-Type: application/json' -d '{"id": 1, "title": "Primera nota editada", "description": "Esta es una nota de prueba editada"}'


curl -X DELETE http://localhost:8080/notes?id=1

curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -

sudo apt-get install -y nodejs

npm install -g @vue/cli

vue --version

npm install --global vue-cli

vue init webpack
front

npm run dev

npm run build

npm install axios --save

npm i bootstrap-vue

npm i vue bootstrap-vue bootstrap