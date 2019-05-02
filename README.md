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

##### 초기화

초기화 작업으로 `Vagrantfile` 설정 파일을 생성한다.

`vagrant init`




### 참고자료

[](https://rorlab.org/rblogs/232)

[](http://taewan.kim/post/vagrant_intro/)