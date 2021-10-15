package sha11

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestSha11(t *testing.T) {

	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	ss := fmt.Sprintf("%x\n", bs)
	fmt.Println(ss[:32])
}

func TestHmac(t *testing.T) {
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))

	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", mac.Sum(nil))
}

func TestEncodeUtf8(t *testing.T) {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
}

// ReqCC request to CC info
type ReqCC struct {
	AK     string
	SK     string
	Now    string
	Nonce  string
	Method string
	Path   string
	Body   []byte
	Sign   string
}

// genSign gen sign
func (r *ReqCC) updateSign() error {
	bMd5 := md5.Sum(r.Body)
	md5Body := fmt.Sprintf("%x", bMd5)
	method := strings.ToUpper(r.Method)
	sha1Str := strings.Join([]string{r.Now, r.Nonce, method, r.Path, md5Body}, "&")
	hmacSha1 := hmac.New(sha1.New, []byte(r.SK))
	hmacSha1.Write([]byte(sha1Str))
	sign := hex.EncodeToString(hmacSha1.Sum(nil))
	r.Sign = sign
	fmt.Println(r.Sign)
	return nil
}

type usageStats struct {
	Rum    int64 `json:"rum_session"`
	Log    int64 `json:"log"`
	Trace  int64 `json:"trace"`
	BLog   int64 `json:"backup_log"`
	DK     int64 `json:"datakit"`
	Job    int64 `json:"job"`
	Series int64 `json:"series"`
}

// 每个工作空间的使用统计量
type workSpaceUsage struct {
	WN    string      `json:"workspace_name"`
	WID   string      `json:"workspace_uuid"`
	Date  string      `json:"date"`
	Stats *usageStats `json:"stats"`
}

func genBody() []byte {
	us := usageStats{}
	us.Rum = 300
	us.Log = 300
	us.Trace = 200
	us.BLog = 200
	us.DK = 50
	us.Job = 30
	us.Series = 300
	ws := workSpaceUsage{}
	ws.Stats = &us
	ws.WN = "测试空间1"
	ws.WID = "wksp_xxx"
	ws.Date = "2021-05-25"
	wss := []workSpaceUsage{
		ws,
	}
	res, err := json.Marshal(wss)
	if err != nil {
		return nil
	}
	return res
}

// PostUsageInfo post usage info
func TestPostUsageInfo(t *testing.T) {
	reqCC := ReqCC{}
	reqCC.AK = "ak-e8KHNqsPvs5hsPGi"
	reqCC.SK = "zUUGA25ITxzRP8PqLHbQ2S0kd0GZCIRU"
	reqCC.Now = "1622438749"
	reqCC.Nonce = "9285e594f69242e9810d5186bfa92d5a"
	// nowStr := fmt.Sprintln(time.Now().Unix())
	// nonce := genUUID(nowStr)
	reqCC.Body = genBody()
	reqCC.updateSign()
	fmt.Println("---sign--", reqCC.Sign)

}
