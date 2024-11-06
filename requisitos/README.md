# Requisitos necessários para subir o ambiente K8S no WSL


Possuir o wsl2 instalado e após esse processo seguir com a instalação do docker, nos seguintes passos abaixo:

### Install docker
É possível rodar o Docker dentro do WSL do Windows, sem precisar instalar o Docker Desktop. Abaixo os detalhes da instalação.
```
Documentação oficial: https://docs.docker.com/engine/install/ubuntu/ https://docs.docker.com/engine/install/linux-postinstall/

E um artigo (em inglês) bastante detalhado do processo: https://dev.to/bowmanjd/install-docker-on-windows-wsl-without-docker-desktop-34m9

```
### Instalação do WSL2
Containers normalmente rodam no Linux e para rodar esses containers Linux dentro do Windows uma ótima opção é instalar o WSL (Windows Subsystem for Linux). O WSL Permite instalar uma distribuição do Linux (por padrão a distribuição Ubuntu) dentro do Windows

Para verificar se você já tem o WSL2 instalado, abra um Prompt do PowerShell e execute:
```
wsl -l -v
```

### Etapa 1 – Habilitar o Subsistema do Windows para Linux
```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```

### Etapa 2 – Habilitar o recurso de Máquina Virtual
```
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

### Etapa 3 – Baixar o pacote de atualização do kernel do Linux e em seguida efetuar a instalação
```
https://wslstorestorage.blob.core.windows.net/wslblob/wsl_update_x64.msi
```

### Etapa 4 – Definir o WSL 2 como a sua versão padrão
```
wsl --set-default-version 2
```

### Etapa 5 – Instalar a distribuição do Linux de sua escolha
```
Ubuntu 20.04 LTS - https://aka.ms/wslubuntu2004
Ubuntu 18.04 LTS - https://aka.ms/wsl-ubuntu-1804
Debian GNU/Linux - https://aka.ms/wsl-debian-gnulinux
Fedora Remix para WSL - https://github.com/WhitewaterFoundry/WSLFedoraRemix/releases/
```

### Etapa 6 - Após escolher a distribuição a ser utilizada, execute o comando abaixo no terminal powershell modificando para o seu caminho raiz e o nome de sua distribuição, ex:
```
Add-AppxPackage C:\Users\ramps\Downloads\Ubuntu_2004.2020.424.0_x64.appx
```

# Instalação do Docker no WSL
Ubuntu 20.10+
Se você estiver rodando o Ubuntu em uma versão anterior ao Ubuntu 20.10, é necessário atualizar:
Verificar a versão instalada:
```
cat /etc/os-release
```

Se for inferior a 20.10, atualizar o Ubuntu:
```
sudo apt-get update
sudo apt dist-upgrade
sudo apt install update-manager-core
sudo apt autoremove
sudo do-release-upgrade -d
```
Ao final do processo, vai ocorrer um shutdown do linux. Feche o terminal e então abra outro prompt do Ubuntu (Iniciar / pesquisar “Ubuntu”).

### Utilizar iptables legadas
Executar:
```
sudo update-alternatives --config iptables
```
(selecionar a opção “iptables-legacy”)

### Instalar o Docker
Remover versões antigas:
```
sudo apt-get remove docker docker-engine docker.io containerd runc
```

Se não tiver docker instalado, vai exibir: “Unable to locate package docker-engine”

Instalar dependências:
```
sudo apt-get update && sudo apt-get upgrade
sudo apt-get install ca-certificates curl gnupg lsb-release
sudo apt-get autoremove
```

Adicionar chave do repositório oficial do Docker:
```
sudo mkdir -p /etc/apt/keyrings

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

Instalar o Docker Engine
```
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

Teste:
```
docker --version
docker compose version
```

### Adicionar permissões
Executar:
```
sudo groupadd docker
sudo usermod -aG docker $USER
```

### Executando & Iniciar o Docker
Executar:
```
sudo dockerd
```
Se for iniciado com sucesso, será exibido: API listen on /mnt/wsl/shared-docker/docker.sock

### Após efetuar a instalação execute o seguinte comando abaixo para habilitar o service do docker 
```
sudo service docker start  
```

### Verificar se o status do docker ficou  * Docker is running
```
sudo service docker status
```

# E por fim, possuir o k3d instalado para gerenciar o kubernets, Install current latest release

```
wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

Site, procurar a seção Overview e descer até Install Script onde obterá o comando para realizar a instalação conforme acima.

```
https://k3d.io/v5.5.1/
```

# Instalação kubectl, kubectx e kubens
Documentação oficial K8S kubectl
```
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/
```

kubectx e kubens
Como utilizar - https://gasparbarancelli.com/post/kubectx-e-kubens-ferramentas-para-auxiliar-na-administracao-do-Kubernetes?lang=pt
```
sudo git clone https://github.com/ahmetb/kubectx /opt/kubectx
sudo ln -s /opt/kubectx/kubectx /usr/local/bin/kubectx
sudo ln -s /opt/kubectx/kubens /usr/local/bin/kubens
```

# Criar alias dentro do seu SO ubuntu para agilizar a digitação
Editar o seu arquivo "vim /root/.bashrc" e procurar pelo bloco alias e inserir os comandos abaixo e ao final aplicar com o comando "source /root/.bashrc"

```
alias k='kubectl'
alias kgh='kubectl get hpa'
alias kx='kubectx'
alias kb='kubens'
alias kgp='kubectl get pod'
alias kdp='kubectl delete pod'
alias kgn='kubectl get namespace'
alias kgs='kubectl get svc'
```

# WSL - MACETES AND COMANDOS TROUBLESHOOT
```
wsl --list --running   -> Listar imagens operando
 
wsl --list   -> Listar imagens instaladas
 
wsl --shutdown   -> Desligar tudo: Build 18917+
 
wsl -t <DistroName>   -> Encerrar distribuição específica: Windows 1903+
 
netsh winsock reset   -> problema com logon, executar comando e reiniciar a máquina para resolver o problema
````
