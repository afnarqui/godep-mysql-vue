mkdir -p $GOPATH/src/github.com/afnarqui
cd $GOPATH/src/github.com/afnarqui
dep init
docker run --name godep -P --volumes-from volumen -v c:/Users/anaranjo/proyectos:/local -it afnarqui/godep bash
docker inspect godep | grep -i tcp
dep ensure --add github.com/joho/godotenv
