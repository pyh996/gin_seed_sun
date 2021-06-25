package initialize

import (
	"github.com/minio/minio-go"
	_ "github.com/minio/minio-go/pkg/encrypt"
	"go_gin/global"
	"go_gin/utils"
	"log"
)

func InitMinIO() {
	minioInfo := global.Settings.MinioInfo
	// 初使化 minio client对象。
	minioClient, err := minio.New(minioInfo.Endpoint, minioInfo.AccessKeyID, minioInfo.SecretAccessKey, false )
	if err != nil {
		log.Fatalln(err)
	}
	//
	global.MinioClient = minioClient
	//创建一个叫userheader的存储桶。
	utils.CreateMinoBuket("userheader")
}
