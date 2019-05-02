# Vagrant + K8s

## Vagrant

코드로 가상머신 제어

* VirtualBox, VMWare 지원

* chef로 추가 설정 가능

* 유용한 플러그인

```bash
vagrant plugin install vagrant-vbguest # vagrant guest 버전 업데이트
vagrant plugin install vagrant-librarian-chef-nochef # vagrant 시작 시 chef 실행
```

### 시작하기

`Vagrant` 작업을 할 폴더 생성

`mkdir vm_develop && cd vm_develop`

#### 작업 단계

* Box 이미지 다운로드

`vagrant box add`

* Box 이미지를 이용해 프로젝트 생성

* 프로젝트의 `Vagrantfile` 수정

* 프로젝트의 가상 인스턴스 시작

* 가상 이미지 접속, 작업

* 가상 이미지 종료

Box는 Vagrant 가상 이미지로 사용할 수 있는 바이너리임
[](https://atlas.hashicorp.com/boxes/search)에서 hashicorp사가 제공하는 box 이미지 확인 가능

박스는 `~/vagrant.d/boxes` 디렉터리에 저장됨

##### box 설치

```bash
vagrant box add  ubuntu/trusty64
vagrant box list
```


##### 초기화

초기화 작업으로 `Vagrantfile` 설정 파일을 생성한다.

```
vagrant init ubuntu/trusty64
```

#### 구동 및 접근

```
vagrant up
vagrant ssh
```

### 참고자료

[](https://rorlab.org/rblogs/232)
