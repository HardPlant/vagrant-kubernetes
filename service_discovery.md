# K8S Service Discovery

[](https://bcho.tistory.com/1262)

## 서비스 정의 생성

* 지정한 IP로 생성 가능

* 여러 Pods를 묶어 로드 밸런싱 가능

* 고유한 DNS 이름을 가질 수 있음

```yml
apiVersion: v1
kind: Service
metadata:
  name: dummy-service
spec:
  type: ClusterIP
  selector:
    app: dummy
  ports:
  - port: 3000
    targetPort: 8080
```

## 서비스 타입

### ClusterIP

디폴트 설정이다. 서비스에 내부 IP를 할당한다. 클러스터 내부에서만 이 서비스에 접근 가능하다.

### LoadBalancer

클라우드 벤더에서 제공한다. 외부 IP를 가지고 있는 로드밸런서를 할당한다. 외부 IP를 가지고 있기 떄문에 외부에서 접근 가능하다.

### NodePort

클러스터 IP 및 노트 IP, Port를 모두 공개한다.

### ExternalName

쿠버네티스 내부 -> 외부 서비스를 호출
Pods가 클러스터 IP를 가지고 있기 떄문에, 외부 서비스 응답 같은 것 처리하는 문제 등 때문에 NAT 설정 등 복잡한 설정이 필요하다.
특히 클라우드 환경의 DB 등을 호출할 필요가 있을 때 사용한다.

```yml
kind: Service
apiVersion: v1
metadata:
  name: my-service
  namespace: prod
spec:
  type: ExternalName
  externalName: my.database.example.com
```

서비스 타입이 External이면 들어오는 모든 요청을 포워딩한다. (프록시처럼)

### Headless Service

서비스는 K8S가 제공하는 Cluster IP, External IP를 통해 접근 가능하게 된다.
MSA에서 엔드포인트 IP를 찾는 기능이 필요하다 (Service Discovery)
만약 서비스 디스커버리 솔루션을 사용한다면, 쿠버네티스가 Cluster IP 대신 DNS 이름을 가지게 한다.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello-node-svc-headless
spec:
  clusterIP: None
  selector:
    app: hello-node
  ports:
    - name: http
      port: 80
      protocol: TCP
      targetPort: 8080
```

이 상태에서 `kubectl get svc`로 조회해보면 해당 서비스에 Cluster IP가 할당되지 않는다.
대신 다른 Pod에서 `nslookup`으로 서비스 dns 이름을 조회 가능함을 확인할 수 있다.

### Service Discovery

생성된 서비스의 IP를 알 수 있는 방법

* DNS

쿠버네티스 내부에서만 통신하는 경우
[서비스명].[네임스페이스명].svc.cluster.local 이라는 DNS 명으로 서비스에 접근 가능하다. (내부 DNS에 등록)
DNS에서 리턴해주는 IP는 클러스터 IP이다.


* External IP


```yml
apiVersion: v1
kind: Service
metadata:
  name: hello-node-svc
spec:
  selector:
    app: hello-node
  ports:
    - name: http
      port: 80
      protocol: TCP
      targetPort: 8080
  externalIPs:
  - 80.11.12.11 # 해당 부분을 지정
```

* LoadBalancer

구글 클라우드의 경우..

```
gcloud compute addresses create [IP 리소스명] --region [리전]
```
로 IP를 만든 뒤에

```
apiVersion: v1
kind: Service
metadata:
  name: hello-node-svc
spec:
  selector:
    app: hello-node
  ports:
    - name: http
      port: 80
      protocol: TCP
      targetPort: 8080
  type: LoadBalancer
  loadBalancerIP: 35.200.64.17
```