# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "perk/ubuntu-2204-arm64"
  config.vm.synced_folder "./", "/bootcamp"
  config.vm.provision "shell", path: "install.sh"
  config.vm.network "forwarded_port", guest: 4000, host: 8080
end
