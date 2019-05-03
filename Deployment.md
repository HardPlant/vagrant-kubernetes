# minikube 디플로이먼트 만들기

[](https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/)

### minikube 시작

* minikube 시작

```
minikube start
```

* dashboard 확인

```
minikube dashboard
```

### 디플로이먼트 만들기

* 파드

하나 이상의 컨테이너 그룹이다. 관리와 네트워킹 목적으로 함꼐 묶여 있다.

* 디플로이먼트

파드의 헬스를 검사해 파드의 컨테이너가 종료되었다면 재시작한다.
파드 생성, 스케일링을 관리할 수 있다.

```
kubectl create deployment hello-node --image=gcr.io/hello-minikube-zero-install/hello-node
```

#### 명령어들

```ps1
kubectl get deployments # 디플로이먼트 보기
kubectl get pods #파드 보기
kubectl get events #클러스터 이벤트 보기
kubectl config view #`kubectl` 환경설정 보기
```

#### 서비스 만들기

가상 네트워크 외부에서 접근하려면 파드를 쿠버네티스 `서비스`로 노출해야 한다.

```
kubectl expose deployment hello-node --type=LoadBalancer --port=8080
```

* 서비스 살펴보기

```
kubectl get services
```

* 서비스 접근하기

```
minikube service hello-node
```

#### 서비스 제거하기

```
kubectl delete service hello-node
kubectl delete deployment hello-node
```

#### minikube 정지/제거하기

```
minikube stop
```

명령어로 정지한다.

```
minikube delete
```

명령어로 VM을 제거한다.

#### 디플로이먼트 객체 만들기

디플로이먼트로 `Pods`, `ReplicaSets`를 선언적으로 생성할 수 있다.
`Deployment Object`는 선언적으로 `유지되어야 할 상태`를 지정할 수 있다.

[](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
)
```
```

###### 예제

다음 예제는 ReplicaSet으로 `nginx` 파드 3개를 띄운다.

`controllers/nginx-deployment.yaml`

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment # Deployment 이름
  labels:
    app: nginx
spec:
  replicas: 3 # 복제할 파드의 갯수
  selector: # Deployment가 관리할 Pods를 찾는 방법
    matchLabels:
      app: nginx
  template: # pod 템플릿
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

* `metadata.name` : 디플로이먼트의 이름

* `replicas`

* `selector`

여기서는 `Pods` 템플릿에 설정되어 있는 라벨값(`app:nginx`)으로 간단하게 찾고 있다.

* `template`

`template.labels`: pod를 라벨링한다. (`app: ngnx`)

`template.spec` : 실행할 컨테이너를 지정한다. (여러 개 가능)

`image` 필드를 이용해 `nginx:1.7.9` 이미지를 Docker Hub에서 찾는다.
`name` 필드로 컨테이너 이름을 `nginx`로 지정하고
`ports` 필드로 `80` 포트를 연다.

* 실행

```
kubectl apply -f https://k8s.io/examples/controllers/nginx-deployment.yaml
```

* 확인

```
kubectl get deployments
```

#### Deployment 업데이트하기

`ngix:1.9.1` 이미지를 `nginx:1.7.9`로 업데이트하려면

```
kubectl --record deployment.apps/nginx-deployment set image deployment.v1.apps/nginx-deployment nginx=nginx:1.9.1
```

으로 수정을 하거나, yaml 파일을 수정하고 (`.spec.template.spec.containers[0].image : nginx:1.9.1`)

```
kubectl edit deployment.v1.apps/nginx-deployment
```

을 실행한다.

롤아웃 상태를 확인하려면

```
kubectl rollout status deployment.v1.apps/nginx-deployment
```

* 레플리카 세트 확인하기

```
kubectl get rs
```

으로 롤아웃으로 대체되고 있는 레플리카 세트를 확인할 수 있다.

