package main

import (
	"GoHelper/QCloud"
	"GoHelper/Util"
	"path/filepath"
)

var preferences = Util.ParsePreferenceMap("./preference.config")

func main() {

	const QCloudKey = "QCloud"

	var appId = preferences[QCloudKey].Config["appId"]
	var secretId = preferences[QCloudKey].Config["secretId"]
	var secretKey = preferences[QCloudKey].Config["secretKey"]
	var QCloudClient = QCloud.Client{
		AppId: appId,
		SecretKey: secretId,
		SecretId: secretKey,
	}

	filePath, err := filepath.Abs("../../Desktop/test.png")
	if err != nil {

	}
	QCloud.UploadObject(&QCloudClient, filePath)
}