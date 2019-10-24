#!/bin/bash
set -ex

# Turning Swap off
swapoff -a

# Installing Prerequisites
apt install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Installing Docker
apt update
apt install -y docker-ce

# Adding user to docker group
usermod -aG docker $USER

# Setting Docker daemon to same cgroup as kubelet
bash -c 'cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF'

# Creating systemd directory for docker
mkdir -p /etc/systemd/system/docker.service.d

# Restarting docker
systemctl daemon-reload
systemctl restart docker

# Installing dependencies for kubernetes tools
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo bash -c "cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF"

# Installing kubernetes packages
apt-get update
apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl
