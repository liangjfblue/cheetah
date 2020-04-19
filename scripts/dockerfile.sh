#!/bin/bash

set -e

gen(){
    pname=$1_$2

    filepath=./deployments/bin/$1/${pname}
    if [[ ! -d ${filepath} ]];then
        mkdir -p ${filepath}
    fi

    if [[ ! -f ${filepath}/Dockerfile ]];then
        touch ${filepath}/Dockerfile
    fi

cat>${filepath}/Dockerfile<<EOF
FROM alpine
COPY . .
ENTRYPOINT ["./${pname}"]
EOF
    echo "create dockerfile over: $pname"
}

genAll() {
    gen srv user
}

case $1 in
    all) echo "build all dockerfile"
    genAll
    ;;
    one) echo "create dockerfile start:" $2_$3
    if [[ -z $2 || -z $3 ]];then
        echo "param error"
        exit 2
    fi
    gen $2 $3
    ;;
    *)
    echo -e "\n\tusage: \n\n\
\tfirst run build.sh \n\n\
\tthen run dockerfile.sh\n\
\t1)dockerfile.sh all\n\
\t2)dockerfile.sh one [web/srv] [name] :make Dockerfile web common\n"
    exit 2
    ;;
esac