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
      mountPath: /data/shared
  - name: nginx
    image: nginx
    volumeMounts:
    - name: shared-storage
      mountPath: /data/shared
  volumes:
  - name : shared-storage
    emptyDir: {}
```

Pod에 nginx, redis 컨테이너를 생성하고, emptyDir를 공유하는 설정
