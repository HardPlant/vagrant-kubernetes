# Ansible

## Inventory file

* Inventory

`/etc/ansible/hosts`

```
192.0.2.50
aserver.example.org
bserver.example.org
```

이 시스템들의 `authorized_keys`에 SSH 공개키가 있어야 됨

```
ssh-agent bash
ssh-add ~/.ssh/id_rsa
```

* ping

```bash
ansible all -m ping
ansible all -m ping -u bruce # as bruce
ansible all -m ping -u bruce --become # sudo to root
ansible all -m ping -u bruce --become --become-user batman # sudoing to batman
```

* command

```
ansible all -a "/bin/echo hello"
```

