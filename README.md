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

#### ssh config 받기

```
vagrant ssh-config --host demo >> ~/.ssh/config
```

ssh [hostname]으로 접근 가능하게 해줌

#### 상태 확인

```
vagrant status
```

#### 명령 요약

```bash
vagrant -v          #: 버전 확인
vagrant status      #: 현재 프로젝트의 가상 이미지 상태 요약
vagrant global-status #: 호스트 머신 전체의 Vagrant 가상 이미지들의 상태 확인
vagrant up          #: Vagrant 가상 이미지 시작
vagrant halt            #: 가상 인스턴스 강제 종료
vagrant destroy             #: 가상 이미지 종료 및 기존 이미지 삭제
vagrant suspend             #: 가상 인스턴스 하이버네이트, 상태 보존
vagrant resume          #: 중지된 인스턴스 시작
vagrant reload          #: 변경된 VagrantFile 적용
vagrant ssh             #: 현재 프로젝트의 가상 이미지에 ssh 접근
```

#### Multi-VM

* ip, host 변경 필요

```
vm.hostname
vm.network "public_network", ip: "10.0.2.21"
```

전체 파일은...

```Vagrantfile
Vagrant.configure("2") do |config|
  config.vm.provision "shell", inline: "echo Hello"

  config.vm.define "master" do |master|
    master.vm.box = "centos66"
    master.vm.hostname = "TestVM1"
    master.vm.network "public_network", ip: "10.0.2.21"
    # modify mem, cpu
    config.vm.provider "virtualbox" do |v|
      v.memory = 1024
      v.cpus = 2
    end
  end

  config.vm.define "minion" do |minion|
    minion.vm.box = "centos66"
    minion.vm.hostname = "TestVM2"
    minion.vm.network "public_network", ip: "10.0.2.22"
  end
end
```

#### 자원 할당

```vagrantfile
# vm 이름설정
config.vm.provider "virtualbox" do |v|
  v.name = "my_vm"
end

# CPU, MEM 조정
config.vm.provider "virtualbox" do |v|
  v.memory = 1024
  v.cpus = 2
end

# 또는 아래와 같이 할 수 있다.
config.vm.provider "virtualbox" do |v|
  v.customize ["modifyvm", :id, "--memory", "2048"]
  v.customize ["modifyvm", :id, "--cpus", "2"]   
end
[출처] [vagrant] Multi Machine (vm 여러개 띄우기) & CPU개수 조정|작성자 asdf
```

#### 참고자료

[](https://rorlab.org/rblogs/232)
[](http://taewan.kim/post/vagrant_intro/)
[](http://blog.naver.com/PostView.nhn?blogId=sory1008&logNo=220759894761&categoryNo=0&parentCategoryNo=0&viewDate=&currentPage=1&postListTopCurrentPage=1&from=postView)

### Kubernetes 설치

#### Master Node

###### kubectl 설치

* kubectl 다운로드

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
```

* $PATH 경로에 넣기

```
sudo mv ./kubectl /usr/local/bin/kubectl
```

* 확인

```
kubectl version
```

###### minikube 설치

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/ 
```

#### 참고자료

[](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

