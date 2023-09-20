package control

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"cmdb_backend/global"
	"cmdb_backend/logger"
	"cmdb_backend/property"
	"os"

	//"github.com/vertica/vertica-sql-go/logger"
	"golang.org/x/crypto/ssh"
)

func send_cmd_InsecureWay(identify_file string, addr string, cmd string) (string, error) {
	key, err := os.ReadFile(identify_file)
	if err != nil {
		return "", err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return "", err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//connect
	peer_addr := addr + ":22"
	client, err := ssh.Dial("tcp", peer_addr, config)
	if err != nil {
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}

	result, err := session.Output(cmd)
	if err != nil {
		return "", err
	}

	return string(result), err

}

func send_file_InsecureWay(identify_file string, addr string, file_name string, work_dir string) (bool, error) {
	key, err := os.ReadFile(identify_file)
	if err != nil {
		return false, err
	}
	singer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return false, err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(singer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	peer_addr := addr + ":22"

	client, err := ssh.Dial("tcp", peer_addr, config)
	if err != nil {
		return false, err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		logger.Error.Fatal("can not create new session ")
	}
	defer session.Close()

	file_dir := global.Env_file["send_file_dir"]
	file_abs_name := file_dir + file_name
	fd, err := os.Open(file_abs_name)
	if err != nil {
		return false, err
	}
	stat, err := fd.Stat()
	if err != nil {
		return false, err
	}
	hostpipe, err := session.StdinPipe()
	defer hostpipe.Close()

	fmt.Fprintf(hostpipe, "C0664 %d %s \n", stat.Size(), file_name)
	io.Copy(hostpipe, fd)
	fmt.Fprintf(hostpipe, "\x00")

	err = session.Run("/usr/bin/scp -t " + work_dir)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func Send_Cmd_with_HostKey() {
	//todo
}

func (c *Send_Cmd_Struct) Send_Cmd(db *sql.DB) (map[string]string, map[string]string, error) {

	var succ map[string]string
	var fail map[string]string

	if len(c.Hosts) == 0 {
		//说明输入的被控主机为0
		return nil, nil, errors.New("Please input  hosts address")
	}

	for _, i := range c.Hosts {
		//if hosts exist
		exist, err := property.Check_Property_Exist_By_IP(db, i)
		if err != nil {
			return nil, nil, errors.New(err.Error())
		}
		//not exist
		if exist == false {
			return nil, nil, errors.New("Host not exist")
		}

		result, err := send_cmd_InsecureWay(global.Env_file["identify_file"], i, c.Cmd)
		if err != nil {
			logger.Error.Fatal(err)
		}

		if result == "" {
			//cmd run fail
			fail[i] = err.Error()
		}
		succ[i] = result
	}
	return succ, fail, nil

}

func (c *Send_File_Struct) Send_File(db *sql.DB) ([]string, map[string]string, error) {
	var succ []string
	var fail map[string]string

	if c.Hosts == nil {
		return nil, nil, errors.New("Please input  hosts address")
	}

	for _, i := range c.Hosts {
		//if hosts exist
		exist, err := property.Check_Property_Exist_By_IP(db, i)
		if err != nil {
			return nil, nil, errors.New(err.Error())
		}
		//not exist
		if exist == false {
			return nil, nil, errors.New("Host not exist")
		}

		if_succ, err := send_file_InsecureWay(global.Env_file["identify_file"], i, c.File_Name, c.Work_Dir)

		if if_succ != true {
			fail[i] = err.Error()
		} else {
			succ = append(succ, i)
		}
	}
	return succ, fail, nil
}
