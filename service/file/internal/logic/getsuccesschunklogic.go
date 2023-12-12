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
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"path"
	"sort"
	"strconv"
	"strings"
)

type GetSuccessChunkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSuccessChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSuccessChunkLogic {
	return &GetSuccessChunkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type completedParts []minio.CompletePart

func (a completedParts) Len() int           { return len(a) }
func (a completedParts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a completedParts) Less(i, j int) bool { return a[i].PartNumber < a[j].PartNumber }

func (l *GetSuccessChunkLogic) GetSuccessChunk(in *file.GetSuccessChunkRequest) (*file.GetSuccessChunkResponse, error) {
	var uuid, uploadID, chunks string
	qFile, err := db.QueryFileByMD5(l.ctx, in.FileMD5)
	if err != nil {
		return nil, err
	}
	var uploaded = true
	if qFile.UUID == "" {
		uuid = uuid2.NewV4().String()
		bucketName := config.MINIO_BUCKET
		objectName := fmt.Sprintf("%s.%s", strings.TrimPrefix(path.Join(config.MINIO_BASE_PATH, uuid), "/"), in.Filetype)
		uploadID, err = l.svcCtx.MinIOCore.NewMultipartUpload(l.ctx, bucketName, objectName, minio.PutObjectOptions{})
		if err != nil {
			return nil, err
		}
		chunks = ""
		err = db.InsertFile(l.ctx, &db.File{
			Model:    gorm.Model{},
			Url:      "",
			Type:     in.Filetype,
			Md5:      in.FileMD5,
			UUID:     uuid,
			UploadID: uploadID,
		})
		if err != nil {
			return nil, err
		}
		uploaded = false
	} else if qFile.Uploaded == false {
		uuid = qFile.UUID
		uploadID = qFile.UploadID
		bucketName := config.MINIO_BUCKET
		objectName := fmt.Sprintf("%s.%s", strings.TrimPrefix(path.Join(config.MINIO_BASE_PATH, uuid), "/"), in.Filetype)
		partInfos, err := l.svcCtx.MinIOCore.ListObjectParts(l.ctx, bucketName, objectName, uploadID, 0, int(in.TotalChunkCount))
		if err != nil {
			return nil, err
		}

		for _, partInfo := range partInfos.ObjectParts {
			chunks += strconv.Itoa(partInfo.PartNumber) + "-" + partInfo.ETag + ","
		}
		uploaded = false
		// 上传完成
		if len(partInfos.ObjectParts) == int(in.TotalChunkCount) {
			// 文件合并
			var parts []minio.CompletePart
			for _, partInfo := range partInfos.ObjectParts {
				parts = append(parts, minio.CompletePart{
					PartNumber: partInfo.PartNumber,
					ETag:       partInfo.ETag,
				})
			}
			sort.Sort(completedParts(parts))
			_, err = l.svcCtx.MinIOCore.CompleteMultipartUpload(l.ctx, bucketName, objectName, uploadID, parts, minio.PutObjectOptions{})
			if err != nil {
				return nil, err
			}
			// 将url写入t_file中
			accessUrl := fmt.Sprintf("%s/%s/%s", l.svcCtx.Config.MinIO.Endpoint, bucketName, objectName)
			err = db.UpdateFile(l.ctx, &db.File{
				Model: gorm.Model{
					ID: qFile.ID,
				},
				Url:      accessUrl,
				Type:     qFile.Type,
				Md5:      qFile.Md5,
				UUID:     qFile.UUID,
				UploadID: qFile.UploadID,
				Uploaded: true,
			})
			if err != nil {
				return nil, err
			}
			uploaded = true
		}
	}

	return &file.GetSuccessChunkResponse{UuId: uuid, UploadID: uploadID, Chunks: chunks, Uploaded: uploaded}, nil
}
