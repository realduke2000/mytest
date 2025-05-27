package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	r := TreeNode{
		0,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			nil,
		},
	}
	ret := preorderTraversal(&r)
	for _, v := range ret {
		fmt.Println(v)
	}
}

func myapp(ret []int) {
	ret = append(ret, 5)
}

func TestArray(t *testing.T) {
	ret := make([]int, 5)
	ret[0] = 0
	ret[1] = 1
	ret[2] = 2
	myapp(ret[:])
	fmt.Printf("%v\n", ret)
}

func TestSize(t *testing.T) {
	var m1 map[int]int
	m2 := make(map[int]int, 5)
	fmt.Printf("m1=%d\n", len(m1))
	fmt.Printf("m2=%d\n", len(m2))
}

func TestPath(t *testing.T) {
	fmt.Println("URLEncoding")
	s, _ := base64.URLEncoding.DecodeString("L3N5c3RlbS9kaWFnbm9zdGljL2RhdGEtY29sbGVjdA==")
	fmt.Println(string(s))
	s, _ = base64.URLEncoding.DecodeString("L3N5c3RlbS9kaWFnbm9zdGljL2RhdGEtY29sbGVjdA")
	fmt.Println(string(s))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("/system/diagnostic/data-collect")))

	fmt.Println("RawURLEncoding")
	s, _ = base64.RawURLEncoding.DecodeString("L3N5c3RlbS9kaWFnbm9zdGljL2RhdGEtY29sbGVjdA==")
	fmt.Println(string(s))
	s, _ = base64.RawURLEncoding.DecodeString("L3N5c3RlbS9kaWFnbm9zdGljL2RhdGEtY29sbGVjdA")
	fmt.Println(string(s))
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("/system/diagnostic/data-collect")))
}

var (
	fExists = func(path string) bool {
		fi, err := os.Stat(path)
		fmt.Printf("fileinfo: %v\n", fi)
		if err == nil {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return true
	}
)

func TestStat(t *testing.T) {
	if fExists("/var/run/appliance/xxx.xxx") {
		fmt.Println("exists!")
	} else {
		fmt.Println("NOT exists!")
	}
}

type workerResponse struct {
	StatusCode   int         `json:"status-code"`
	ResponseBody interface{} `json:"response-body"`
}


func TestMarshal(t *testing.T) {
	_, err := ioutil.ReadFile("/tmp")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	
} 

// APIAccessToken: Input received from user
type APIAccessToken struct {
	//Token name of API access token
	ID string `json:"id"`

	//Role assign to API access token
	Role string `json:"role"`

	//Period after which token will expire in d (days), w (weeks), M (months), y (years)
	//e.g- 1d, 2M, 3w, 4y (without dots)
	TTL string `json:"time-to-live"`

	//Fingerprints[] is used to identify the client identity.
	//client's fingerprint should be set when creating this token.
	Fingerprints []string `json:"-"`

	//subject-id of a token, can only be set for internal roles
	Subject string `json:"-"`
}
func TestJson(t *testing.T){
	token := APIAccessToken {
		ID: "123456",
		Role: "observer role",
		TTL: "60s",
		Fingerprints: []string{"fp1"},
		Subject: "subject",
	}

	bs, _ := json.Marshal(token)
	fmt.Printf("%s\n", string(bs))
}