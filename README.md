# Ship vm in container
So you can launch a vm by :
```
docker run -v /fanux/centos7.qcow2:/img/centosy.qcow2 -v /dev:/dev --privileged sealvm init
```

And ssh in it :
```
ssh root@172.17.0.2
```

# What is the difference with kata
Sealvm is not a micro or light vm, it will start systemd and sshd in wm. It also using runc runtime, just run qemu in docker.
