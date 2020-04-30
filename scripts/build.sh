#!/bin/bash

set -e

BUILD_HOME=deployments/bin

build() {
    if [[ ! -d ${BUILD_HOME} ]];then
        mkdir -p ${BUILD_HOME}
    fi

    buildDir=./app/interface/$2/cmd
    if [[ "$1" == "srv" ]]; then
        buildDir=./app/service/$2/cmd
    fi

    if [[ -d ${buildDir} ]];then
        for f in ${buildDir}/main.go; do
            if [[ -f ${f} ]];then
                dir=${BUILD_HOME}/$1/$1_$2
                if [[ ! -d ${dir} ]];then
                    mkdir -p ${BUILD_HOME}/$1/$1_$2
                fi

                cp ${buildDir}/config.yaml ${BUILD_HOME}/$1/$1_$2
                CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o ${BUILD_HOME}/$1/$1_$2/$1_$2 ${buildDir}
                echo build over: $1_$2;
            fi \
        done
    fi
}

buildAll() {
    # web
    build web web

    # srv
    build srv web
}

case $1 in
    all) echo "build all"
    buildAll
    echo "make all ok"
    ;;
    one) echo "build:"$2_$3
    if [[ -z $2 || -z $3 ]];then
        echo "param error"
        exit 2
    fi
    build $2 $3
    ;;
    *)
    echo -e "\n\tusage: \n\n\
\tbuild one: build.sh one srv web\n\n\
\tbuild all: build.sh all\n"
    exit 2
    ;;
esac
