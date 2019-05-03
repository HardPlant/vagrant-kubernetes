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

### 컨트롤러

애플리케이션 설정/배포는 앞의 4개 기본 오브젝트로 가능함
기본 오브젝트를 편하게 관리하기 위해 컨트롤러를 사용함

###### 종류
Replication Controller (aka RC), Replication Set, DaemonSet, Job, StatefulSet, Deployment

##### Replication Controller

Equality 기반 Selector를 사용함
(ReplicaSet의 이전 버전, ReplicaSet은 Set 기반 Selector 사용)

Pod를 관리함
지정된 숫자로 기동, 관리함

```yml
spec:
    replicas: 3
    selector: # 해당 셀렉터를 이용해..
        app: nginx
    template:
        metadata:               # Pod 정의에서 그대로 붙여넣음
            name: nginx
            labels:
                app: nginx # 이 키/값 쌍이 설정된 Pods
        spec:                 
            containers:       
            - name: nginx     
                image: nginx:1.7.9
                ports:
                - containerPort: 8090
```

##### Deployment

Replica Set의 상위 추상화 개념임
실제로 사용함

###### 배포 시나리오

* 블루/그린

새 버전 Pod을 3개 띄운 후 모든 트래픽을 즉시 리다이렉트

* 롤링 업그레이드

새 버전 Pod를 1개 띄우고 구버전 Pod 1개를 죽여가며 만듬

이 과정을 추상화함

##### DaemonSet

각 노드에서 Pod을 하나만 돌게 하는 컨트롤러
모니터링, 로그 수집, GPU/Nvme SSD를 사용하는 경우

##### Job

배치 작업
(one-time 파일 변환, 주기적 ETL 배치 작업)
Job 종료 시 Pod를 같이 종료한다

```yml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl # 커맨드를 같이 입력한다
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      restartPolicy: Never
  backoffLimit: 4
```

커맨드 결과를 모니터링해 결과가 실패/성공하는지 확인하고,
실패했다면 설정에 따라 재시도시킬 수 있다.
같은 작업을 여러 번 수행하게 하는 기능이 있고 (completion), 이를 병렬적으로 할 것인지를 지정할 수 있다. (parallelism)

##### Cron Job

스케줄로 Job 컨트롤러에 있는 Pod를 실행한다.

```yml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox
            args:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure
```

##### StatefulSet

DB 등 상태가 있는 Pod을 관리하는 컨트롤러