package httptest

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	userProto "protobuf-demo/proto/pb/users"
	"time"
)

func main() {
	var user = &userProto.User{
		Id:    0,
		Name:  "abc",
		Email: "aa@qq.com",
	}

	data, _ := json.Marshal(user)
	fmt.Println(string(data))

	msg, _ := proto.Marshal(user)
	fmt.Println(string(msg))

	//位置 + 具体的消息

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/protobuf")
		data, _ := proto.Marshal(user)
		w.Write(data)
	})

	go func() {

		c := time.Tick(1 * time.Second)
		for {
			select {
			case <-c:
				resp, err := http.Get("http://localhost:8080")
				if err != nil {
					return
				}
				defer resp.Body.Close()

				data, _ := ioutil.ReadAll(resp.Body)
				var respUser userProto.User
				proto.Unmarshal(data, &respUser)

				fmt.Println("respUser:", respUser)
			}
		}

	}()

	http.ListenAndServe(":8080", nil)

}
