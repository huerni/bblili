package constant

const (
	// key prefix:videoId
	BARRAGE_SORTEDSET string = "bblili-video-service:barrage_sortedset:%s"

	LOCK_BARRAGE_SORTEDEST string = "bblili-video-service:lock:barrage_sortedset"

	// key prefix:videoId_rootID
	VIDEOCOMMENT_SORTEDSET string = "bblili-video-service:video-comment_sortedset:%s_%s"

	LOCK_VIDEOCOMMENT_SORTEDSET string = "bblili-video-service:lock:video-comment_sortedset:%s"
)
