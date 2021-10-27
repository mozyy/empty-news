package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
)

func main() {
	port := flag.Int("port", 1080, "proxy port")
	flag.Parse()
	WSLIP := getWSLIP()
	proxyAddr := fmt.Sprintf("http://%s:%d", WSLIP, *port)
	writeDockerProxy(proxyAddr)

	cmd := exec.Command("setx", "proxy_addr", proxyAddr)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func writeDockerProxy(proxyAddr string) {
	home := os.Getenv("home")
	dockerConfigPath := path.Join(home, ".docker", "config.json")

	configByte, err := ioutil.ReadFile(dockerConfigPath)
	if err != nil || string(configByte) == "" {
		configByte = []byte("{}")
	}

	var config map[string]interface{}
	err = json.Unmarshal(configByte, &config)
	if err != nil {
		config = map[string]interface{}{}
	}
	proxies, proxiesOk := config["proxies"].(map[string]interface{})
	if !proxiesOk {
		proxies = map[string]interface{}{}
	}
	defaul, defaultOk := proxies["default"].(map[string]interface{})
	if !defaultOk {
		defaul = map[string]interface{}{"httpProxy": "", "httpsProxy": ""}
	}
	defaul["httpProxy"] = proxyAddr
	defaul["httpsProxy"] = proxyAddr

	proxies["default"] = defaul
	config["proxies"] = proxies
	fmt.Println(config)
	configByte, err = json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(dockerConfigPath, configByte, 0666)
	if err != nil {
		panic(err)
	}
}

func getWSLIP() string {
	cmd := exec.Command("ipconfig")
	ipconfig, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	// os.WriteFile("out", out, os.ModePerm)
	reg := regexp.MustCompile(`(?sU)WSL.*IPv4.*((?:\d+\.){3}\d)`)
	subs := reg.FindSubmatch(ipconfig)
	// fmt.Println("===============")
	if len(subs) != 2 {
		panic("not find wsl ip")
	}
	// fmt.Println(string(subs[1]))
	return string(subs[1])
}
