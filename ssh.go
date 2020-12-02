package library

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	Defaultkeypath = "~/.ssh/id_rsa"
	Defaulttimeout = time.Second
)

type SSH struct {
	Host     string
	Port     string
	User     string
	Password string
	Type     string
	Keypath  string
	Timeout  time.Duration
}

// new ssh
func Newssh(host, port, user, password, typ, keypath string, timeout time.Duration) *SSH {
	return &SSH{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Type:     typ,
		Keypath:  keypath,
		Timeout:  timeout,
	}

}

// 返回session
func (s *SSH) Conn() *ssh.Client {
	conf := ssh.ClientConfig{
		Timeout:         s.Timeout,
		User:            s.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //不够安全
	}
	if s.Type == "password" {
		conf.Auth = []ssh.AuthMethod{ssh.Password(s.Password)}
	} else {
		conf.Auth = []ssh.AuthMethod{privateKeyPath(s.Keypath)}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	sshClient, _ := ssh.Dial("tcp", addr, &conf)
	/*if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	//执行远程命令
	combo,err := session.CombinedOutput("whoami; cd /; ls -al;echo https://github.com/dejavuzhou/felix")
	if err != nil {
		log.Fatal("远程执行cmd 失败",err)
	}
	log.Println("命令输出:",string(combo))*/
	return sshClient
}

// 从ssh key文件中读取认证信息
func privateKeyPath(sshKeyPath string) ssh.AuthMethod {
	if sshKeyPath[:2] == "~/" {
		sshKeyPath = filepath.Join(userHome(), sshKeyPath[2:])
	}
	buff, err := ioutil.ReadFile(sshKeyPath)
	if err != nil {
		fmt.Errorf("Error while reading SSH key file: %v", err)
		return nil
	}
	signer, err := ssh.ParsePrivateKey(buff)
	if err != nil {
		fmt.Errorf("Error while Parse Private SSH key file: %v", err)
		return nil
	}
	return ssh.PublicKeys(signer)
}

// 获取~目录
func userHome() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	homeDrive := os.Getenv("HOMEDRIVE")
	homePath := os.Getenv("HOMEPATH")
	if homeDrive != "" && homePath != "" {
		return homeDrive + homePath
	}
	return os.Getenv("USERPROFILE")
}
