str=$(cat .env | xargs)
ENV=(${str//=/ })
MODULE_NAME=${ENV[1]}


if [[ "$1" == "make:controller" ]]
then 
    if [[ "$2" != "" ]]
    then
        touch "./controllers/$2.go"

        echo "package controllers"                                          >> ./controllers/$2.go
        echo ""                                                             >> ./controllers/$2.go
        echo "import ("                                                     >> ./controllers/$2.go
        echo "   view \"$MODULE_NAME/utils/view\""                          >> ./controllers/$2.go
        echo \"  "net/http"\"                                               >> ./controllers/$2.go
        echo ")"                                                            >> ./controllers/$2.go
        echo ""                                                             >> ./controllers/$2.go
        echo "func Get$2(w http.ResponseWriter, r *http.Request) {"         >> ./controllers/$2.go
        echo "    var data = map[string]interface{}{"                       >> ./controllers/$2.go
        echo "        \"title\": \"Boilerplate by zuzustack\","             >> ./controllers/$2.go
        echo "        \"name\":  \"Hello World\","                          >> ./controllers/$2.go
        echo "    }"                                                        >> ./controllers/$2.go
        echo ""                                                             >> ./controllers/$2.go
        echo "    view.Render(\"index.html\", data, w)"                     >> ./controllers/$2.go
        echo "}"                                                            >> ./controllers/$2.go
    else
        echo "Please input file name for new controller"
    fi
fi

if [[ "$1" == "start" ]]
then
    sed -i "s/$MODULE_NAME/\"$2\"/g" ./index.go
    sed -i "s/$MODULE_NAME/\"$2\"/g" ./route/web.go 
    sed -i "s/$MODULE_NAME/\"$2\"/g" ./go.mod 
    sed -i "s/$MODULE_NAME/\"$2\"/g" ./controllers/*.go 
fi