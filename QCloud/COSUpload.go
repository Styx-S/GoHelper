package QCloud

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func UploadObject(client *Client, filePath string) {
	u, _ := url.Parse("https://conf-pic-1252384928.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  client.SecretId,
			SecretKey: client.SecretKey,
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "test/test.png"
	//// 1.通过字符串上传对象
	//f := strings.NewReader("test")
	//
	//_, err := c.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	panic(err)
	//}
	// 2.通过本地文件上传对象
	_, err := c.Object.PutFromFile(context.Background(), name, filePath, nil)
	if err != nil {
		panic(err)
	}
	// 3.通过文件流上传对象
	//fd, err := os.Open("./test")
	//if err != nil {
	//	panic(err)
	//}
	////defer fd.Close()
	////_, err = c.Object.Put(context.Background(), name, fd, nil)
	////if err != nil {
	////	panic(err)
	////}
}