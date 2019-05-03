# K8S 볼륨

Pod 단위이다 (컨테이너 단위가 아님)
해당 Pod 내의 컨테이너가 볼륨을 공유할 수 있다

### 종류

[](https://kubernetes.io/docs/concepts/storage/volumes/#types-of-volumes)
* 로컬 디스크
* 외장 디스크 인터페이스
 NFS, iSCSI, Fiber Channel

* 오픈 소스 파일 시스템
GlusterFS, Ceph

* 클라우드 디스크
AWS EBS, GCP Persistent, VsphereVolume

##### 쿠버네티스 내에서

* Temp
emptyDir

* Local
hostPath

* Network
GlusterFS
gitRepo
NFS
iSCSI
gcePersistentDisk
AWS EBS
azureDisk
Fiber Channel
Secret
VshereVolume

### emptyDir

Pod가 생성될 때 생성, Pod가 삭제될 떄 삭제
물리 디스크나 메모리 지정 가능

```yml
apiVersion: v1
kind: Pod
metadata:
  name: shared-volumes 
spec:
  containers:
  - name: redis
    image: redis
    volumeMounts:
    - name: shared-storage
      mountPath: /data/shared #(2)
  - name: nginx
    image: nginx
    volumeMounts:
    - name: shared-storage
      mountPath: /data/shared #(2)
  volumes: #(1)
  - name : shared-storage
    emptyDir: {}
```

Pod에 nginx, redis 컨테이너를 생성하고, emptyDir를 공유하는 설정
`volumes` 필드로 공유된 스토리지를 생성했고, 각 컨테이너의 `volumeMounts` 필드를 이용해 `/data/shared` 경로에 마운트시킨다.

```yml
    emptyDir:
        medium: Memory
```

이렇게 지정하면 메모리에 저장이 된다.

```
kubectl apply -f .\shared-volumes_pod.yaml
```

으로 적용한다.

* pod 접근

```
kubectl exec -it shared-volumes --container redis -- /bin/bash
```

### hostPath

노드의 로컬 디스크의 경로를 Pod에서 마운트에 사용한다
같은 hostPath에 있는 볼륨을 여러 Pod 사이에서 공유해 사용한다
Pod가 삭제되도 hostPath에 있는 파일을 삭제하지 않는다
Pod가 **다른** 노드에서 재시작되면 **이전** 노드의 hostPath가 아닌 **현재** 노드의 hostPath에 액세스한다

로그 솔루션 등에 필요하다.

```yml
apiVersion: v1
kind: Pod
metadata:
  name: hostpath
spec:
  containers:
  - name: redis
    image: redis
    volumeMounts:
    - name: terrypath
      mountPath: /data/shared
  volumes:
  - name : terrypath
    hostPath:
      path: /tmp
      type: Directory
```

이 Pod를 배포한 뒤 ID를 얻고, pod 정보에서

```
kubectl describe pod hostpath
```

Node 프로퍼티를 보면

```
Name:               shared-volumes
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               minikube/10.0.2.15
# ...
```

해당 pod가 어느 Node에서 실행되고 있는지 확인 가능하다.
해당 Node를 확인하여 공유 파일을 검증한다.

### gitRepo

git 저장소의 특정 리비전의 내용을 clone해서 emptyDir을 생성한다.

```
apiVersion: v1
kind: Pod
metadata:
 name: gitrepo-volume-pod
spec:
 containers:
 - image: nginx:alpine
   name: web-server
   volumeMounts:
   - name: html
     mountPath: /usr/share/nginx/html
     readOnly: true
   ports:
   - containerPort: 80
     protocol: TCP
 volumes:
 - name: html
   gitRepo:
        repository: https://github.com/luksa/kubia-website-example.git
        revision: master
        directory: .
```

### PersistentVolume, PersistentVolumeClaim

시스템 관리자와 개발자의 역할을 분리할 수 있다.

시스템 관리자가 물리 디스크를 이 기본 오브젝트로 등록하고,
개발자는 Pod를 생성할 떄 볼륨 정의 부분에 PVC를 지정해 연결한다.

(참조: [](https://bcho.tistory.com/1259?category=731548))