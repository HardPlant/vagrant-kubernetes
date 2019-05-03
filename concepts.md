# K8S 컨셉
[](https://bcho.tistory.com/1256?category=731548)

### 마스터, 노드

* Master .. 클러스터 전체를 관리함

* Node .. 컨테이너가 배포되서 실행됨


### 오브젝트

* 기본 오브젝트 .. 가장 기본적인 구성단위

* 컨트롤러 .. 기본 오브젝트 생성, 관리하는 추가 기능

##### 스펙

yaml, json 파일로 정의되는 스펙
오브젝트의 설정정보를 저장함

##### 기본 오브젝트

애플리케이션의 워크로드

* Pod : 컨테이너화된 애플리케이션 모음

* Service : 로드밸런서

* Volume : 디스크
* Namespace : 패키지명

###### Pod

```yml
apiVersion: v1          # 쿠버네티스 API 버전
kind: Pod               # 리소스 종류
metadata:               # 라벨, 리소스 이름 등
  name: nginx
spec:                   # 리소스의 스펙
  containers:               # 컨테이너를 정의함
  - name: nginx             # 컨테이너 이름
    image: nginx:1.7.9      # 컨테이너 이미지
    ports:
    - containerPort: 8090   # 오픈할 포트
```

* 컨테이너들이 IP와 Port를 공유함

컨테이너들이 localhost로 통신 가능함
컨테이너간 디스크 볼륨을 공유 가능함

* MSA의 사이드카 패턴을 구현함 (주변 프로그램을 같이 배포)

이외에도 [여러 패턴이 있음](https://kubernetes.io/blog/2015/06/the-distributed-system-toolkit-patterns/)

###### Volume

여러 디스크를 추상화함

###### Service

Pods에 접근가능한 네트워크 자원 제공
로드밸런서 (하나의 IP, Port에 Pod을 묶음)
Pod가 삭제 후 재생성되면 IP를 재할당받기 떄문에, Pod를 선택하는 방법이 필요함 (label + label selector)

* 라벨 셀렉터 : 서비스에 묶을 Pod를 정의

```yml
kind: Service       # 리소스 종류
apiVersion: v1
metadata:
  name: my-service  # 서비스 이름
spec:               # 스펙
  selector:
    app: myapp
  ports:
  - protocol: TCP
    port: 80        # 노드의 80 포트로 서비스
    targetPort: 9376    # 컨테이너의 9376 포트로 연결
```

###### Namespace

클러스터 내의 논리적인 분리단위

예를 들어.. 개발/운영/테스트 3개 네임스페이스로 만들 수 있음

* 접근 권한 관리
* 리소스 Quota 할당
* 리소스 나눠서 관리

논리적인 분리 단위이므로, 서로 다른 네임스페이스 내부의 Pod끼리 통신 가능

### 라벨

**리소스**를 선택할 수 있도록 함
키/값 쌍으로도 정의 가능함

```
"metadata": {
  "labels": {
    "key1" : "value1",
    "key2" : "value2"
  }
}
```

이후 서비스 등에서 Pod를 연결하려면

```yml
kind: Service
apiVersion: v1
metadata:
  name: my-service
spec:
  selector:
    app: myapp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376`
```

으로 설정을 하면 pod 중에서 라벨 키 `app` 값이 `myapp`인 pod를 골라 서비스에 바인드한다.