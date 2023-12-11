package logic

import (
	"bblili/service/file/file"
	"bblili/service/file/internal/config"
	"bblili/service/file/internal/db"
	"bblili/service/file/internal/svc"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	uuid2 "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileSlicesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileSlicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileSlicesLogic {
	return &UploadFileSlicesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileSlicesLogic) UploadFileSlices(in *file.UpLoadFileBySlicesRequest) (*file.UploadFileBySliceResponse, error) {
	qFile, err := db.QueryFileByMD5(l.ctx, in.FileMD5)
	if err != nil {
		return nil, err
	}
	var uuid, uploadID string
	if qFile.UUID == "" {
		uuid = uuid2.NewV4().String()
		bucketName := config.MINIO_BUCKET
		objectName := fmt.Sprintf("%s.%s", strings.TrimPrefix(path.Join(config.MINIO_BASE_PATH, uuid), "/"), in.FileType)
		uploadID, err = l.svcCtx.MinIOCore.NewMultipartUpload(l.ctx, bucketName, objectName, minio.PutObjectOptions{})
		if err != nil {
			return nil, err
		}
		err = db.InsertFile(l.ctx, &db.File{
			Model:    gorm.Model{},
			Url:      "",
			Type:     in.FileType,
			Md5:      in.FileMD5,
			UUID:     uuid,
			UploadID: uploadID,
		})
		if err != nil {
			return nil, err
		}
	} else {
		uuid = qFile.UUID
		uploadID = qFile.UploadID
	}

	bucketName := config.MINIO_BUCKET
	objectName := fmt.Sprintf("%s.%s", strings.TrimPrefix(path.Join(config.MINIO_BASE_PATH, uuid), "/"), in.FileType)
	urlValues := make(url.Values)
	// Set part number.
	urlValues.Set("partNumber", strconv.Itoa(int(in.SliceNumber)))
	// Set upload id.
	urlValues.Set("uploadId", uploadID)
	signedUrl, err := l.svcCtx.MinIOCore.Presign(l.ctx, http.MethodPut, bucketName, objectName, time.Hour, urlValues)
	if err != nil {
		return nil, err
	}

	return &file.UploadFileBySliceResponse{Url: signedUrl.String()}, nil
}
