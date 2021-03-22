#/bin/sh

propath=$(cd `dirname $0`;pwd)
pluginPath=${propath}/plugins

setArgsKey=("registry_address")
setArgsDefault=("127.0.0.1:8500")
getArgs=$*
getArgValue=()

for ((i=0; i<${#setArgsKey[*]}; i++)); do
	setOption="--${setArgsKey[i]}"
	argVal=${setArgsDefault[i]}

	for getArg in $getArgs; do 
		argKey=${getArg%=*}

		if [[ "$setOption" == "$argKey" ]]; then 
			argVal=${getArg#*=}
		fi

		echo $argVal
	done

	getArgValue[i]=$argVal
done

registryAddress=getArgValue[0]

curl "http://${registryAddress}/"
