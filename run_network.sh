#!bin/sh

rm -rf pki/database.db

#NUMMIXES=$1
#echo $NUMMIXES

#for (( j=0; j<$NUMMIXES; j++ ));
#do
#    go run main.go -typ=mix -id="Mix$j" -host=localhost -port=$((9990+$j)) > logs/"Mix$j".log &
#done

go run main.go -typ=mix -id=Mix1 -host=localhost -port=9998 > logs/Mix1.log &
go run main.go -typ=mix -id=Mix2 -host=localhost -port=9999 > logs/Mix2.log &
go run main.go -typ=provider -id=Provider -host=localhost -port=9997 > logs/Provider.log ;

# trap call ctrl_c()
trap ctrl_c SIGINT SIGTERM SIGTSTP
function ctrl_c() {
        echo "** Trapped SIGINT, SIGTERM and SIGTSTP"
        kill_port 9998
        kill_port 9999
        kill_port 9997
}

function kill_port() {
    PID=$(lsof -t -i:$1)
    echo "$PID"
    kill -TERM $PID || kill -KILL $PID
}




