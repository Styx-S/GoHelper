package main

import (
	"GoHelper/QCloud"
	"GoHelper/Util"
)

var preferences = Util.ParsePreferenceMap("./preference.config")

func main() {

	const QCloudKey = "QCloud"

	var appId = preferences[QCloudKey].Config["appId"]
	var secretId = preferences[QCloudKey].Config["secretId"]
	var secretKey = preferences[QCloudKey].Config["secretKey"]
	var bucketName = preferences[QCloudKey].Config["bucket"]
	var QCloudClient = QCloud.Client{
		AppId: appId,
		SecretKey: secretId,
		SecretId: secretKey,
	}

	QCloud.UploadObject(&QCloudClient, bucketName, "test.png", "../../Desktop/test.png")
}