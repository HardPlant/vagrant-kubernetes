# Docker 이미지 빌드

### docker 툴박스 설치

[](https://medium.com/@maumribeiro/running-your-own-docker-images-in-minikube-for-windows-ea7383d931f6)
Hyper-V를 설치하지 않고 Docker Client를 사용하기 위해서, [Docker Toolbox](https://github.com/docker/toolbox/releases)를 사용한다.

### docker 클라이언트를 minikube 안에 설정

```
minikube docker-env --shell powershell
```

Docker ToolBox VM 대신 minikube 내의 docker 대몬을 이용하게 환경설정을 해준다.

### 더미 프로그램 만들기

#### Go 설치

[](https://golang.org/doc/install?download=go1.12.4.windows-amd64.msi)

#### 샘플 코드

```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func runProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := []byte("Hello Universe! (because World is too small...)")
	w.Write(response)
	return
}

//NewRouter api router
func newRouter() *mux.Router {
	apiRouter := mux.NewRouter().StrictSlash(true) //mainRouter.PathPrefix("/api").Subrouter().StrictSlash(true)
	//Create Routes
	apiRouter.HandleFunc("/", runProgram).Methods("GET")
	return apiRouter
}
func main() {
	// launch server in port 8500
	log.Fatal(http.ListenAndServe(":8500", newRouter()))
}
```

gorilla/mux가 없다면

```
go get github.com/gorilla/mux
```

### 도커파일 만들기

```dockerfile
#Note: golang:onbuild is a very straightforward way for you to build your own GO app image
FROM golang:onbuild
EXPOSE 8500
```

### 이미지 빌드, 컨테이너 생성

(위에서 했던 Windows 환경설정)

```ps1
minikube start
minikube docker-env
minikube docker-env | Invoke-Expression
```

```
minikube docker-env | Invoke-Expression
```
해당 설정은 쉘이 열릴 때마다 수행해야 한다.

#### 도커 이미지 빌드

```
docker build -t dummy:v0 .
```

### 컨테이너 실행

* 실행

```ps1
kubectl run hello-universe --image=dummy:v0 --port=8500
```

* 네트워크 포트 열기

컨테이너에 접근하기 위해 포트를 열어야 한다.

```ps1
kubectl expose deployment hello-universe --type="LoadBalancer"
```

서비스가 잘 생성됬는지 확인하려면

```
kubectl get services
```

서비스의 url을 얻으려면

```
minikube service hello-universe --url
```

### 스케일링

```
kubectl scale deployment hello-universe --replicas=3
```

### 참고

image 없이 run을 한 뒤 image를 올려도 k8s가 자동으로 pods를 띄워준다.

### 참고자료

[](https://medium.com/@maumribeiro/running-your-own-docker-images-in-minikube-for-windows-ea7383d931f6)

[](https://medium.com/humanscape-tech/kubernetes-%EB%8F%84%EC%9E%85-%EC%A0%84-minikube-%EC%82%AC%EC%9A%A9%EA%B8%B0-2eb2b6d8e444)

