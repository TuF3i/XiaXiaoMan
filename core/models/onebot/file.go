package onebot

type UploadGroupFileRequest struct {
	GroupID  int64  `json:"group_id"`
	File     string `json:"file"`
	Name     string `json:"name"`
	FolderID string `json:"folder_id,omitempty"`
}

type SetGroupFileForeverRequest struct {
	GroupID int64  `json:"group_id"`
	FileID  string `json:"file_id"`
	BusID   int64  `json:"bus_id"`
}

type DeleteGroupFileRequest struct {
	GroupID int64  `json:"group_id"`
	FileID  string `json:"file_id"`
	BusID   int64  `json:"bus_id"`
}

type MoveGroupFileRequest struct {
	GroupID      int64  `json:"group_id"`
	FileID       string `json:"file_id"`
	BusID        int64  `json:"bus_id"`
	TargetFolder string `json:"target_folder"`
}

type CreateGroupFileFolderRequest struct {
	GroupID  int64  `json:"group_id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id,omitempty"`
}

type DeleteGroupFolderRequest struct {
	GroupID  int64  `json:"group_id"`
	FolderID string `json:"folder_id"`
}

type GetGroupFileSystemInfoRequest struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupRootFilesRequest struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupFilesByFolderRequest struct {
	GroupID  int64  `json:"group_id"`
	FolderID string `json:"folder_id"`
}

type RenameGroupFileFolderRequest struct {
	GroupID  int64  `json:"group_id"`
	FolderID string `json:"folder_id"`
	Name     string `json:"name"`
}

type GetGroupFileURLRequest struct {
	GroupID int64  `json:"group_id"`
	FileID  string `json:"file_id"`
	BusID   int64  `json:"bus_id"`
}

type GetPrivateFileURLRequest struct {
	FileID string `json:"file_id"`
}

type UploadPrivateFileRequest struct {
	UserID int64  `json:"user_id"`
	File   string `json:"file"`
	Name   string `json:"name"`
}

type UploadFlashFileRequest struct {
	File     string `json:"file"`
	FileName string `json:"file_name,omitempty"`
}

type DownloadFlashFileRequest struct {
	FileMD5  string `json:"file_md5"`
	FileName string `json:"file_name,omitempty"`
}

type GetFlashFileInfoRequest struct {
	MsgID int64 `json:"msg_id"`
}

type DownloadFileRequest struct {
	URL         string `json:"url"`
	ThreadCount int    `json:"thread_count,omitempty"`
	Headers     string `json:"headers,omitempty"`
}

type GroupFileSystemInfo struct {
	FileCount  int64 `json:"file_count"`
	LimitCount int64 `json:"limit_count"`
	UsedSpace  int64 `json:"used_space"`
	TotalSpace int64 `json:"total_space"`
}

type GroupFiles struct {
	Files   []GroupFile   `json:"files"`
	Folders []GroupFolder `json:"folders"`
}

type GroupFile struct {
	FileID       string `json:"file_id"`
	FileName     string `json:"file_name"`
	FileSize     int64  `json:"file_size"`
	BusID        int64  `json:"bus_id"`
	UploadTime   int64  `json:"upload_time"`
	DeadTime     int64  `json:"dead_time"`
	ModifyTime   int64  `json:"modify_time"`
	DownloadTime int64  `json:"download_time"`
	Uploader     int64  `json:"uploader"`
	UploaderName string `json:"uploader_name"`
	MD5          string `json:"md5"`
	SHA1         string `json:"sha1"`
	SHA3         string `json:"sha3"`
}

type GroupFolder struct {
	FolderID       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int64  `json:"create_time"`
	Creator        int64  `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int    `json:"total_file_count"`
}

type GroupFileURL struct {
	URL string `json:"url"`
}

type PrivateFileURL struct {
	URL string `json:"url"`
}

type FlashFileInfo struct {
	FileMD5  string `json:"file_md5"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	ImageURL string `json:"image_url"`
}

type DownloadFileInfo struct {
	File string `json:"file"`
}
