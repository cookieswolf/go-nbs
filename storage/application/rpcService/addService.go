package rpcService

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/NBSChain/go-nbs/console/pb"
	"github.com/NBSChain/go-nbs/storage/application/rpcServiceImpl"
	"github.com/NBSChain/go-nbs/utils/crypto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"io"
	"sync"
)

const DefaultStreamTaskNo = 30
const StreamSessionIDKey = "sessionid"

type addService struct {
	taskLock    sync.RWMutex
	fileAddTask map[string]*pb.AddRequest
}

/*****************************************************************
*
*		service callback function.
*
*****************************************************************/
func (service *addService) AddFile(ctx context.Context, request *pb.AddRequest) (*pb.AddResponse, error) {

	switch request.FileType {

	case pb.FileType_LARGEFILE:
		{
			if len(service.fileAddTask) >= DefaultStreamTaskNo {
				return nil, errors.New("too many add task running now. ")
			}

			sessionId := crypto.MD5SS(request.FullPath + string(request.FileSize))

			service.appendTask(sessionId, request)

			return &pb.AddResponse{Message: "continue", SessionId: sessionId}, nil
		}

	case pb.FileType_FILE:
		{
			if cs := request.SplitterSize; cs == 0 {
				request.SplitterSize = SplitterSize
			}

			reader := &fileReader{
				reader: bytes.NewReader(request.FileData),
			}

			importer := &RpcFileImporter{
				splitterSize: request.SplitterSize,
				reader:       reader,
				fileName:     request.FileName,
				fullPath:     request.FullPath,
				isDirectory:  false,
				Out:          make(chan *pb.AddResponse),
			}

			err := rpcServiceImpl.ImportFile(importer)

			if err != nil {
				return nil, err
			} else {
				result := <-importer.Out

				return result, nil
			}
		}
	case pb.FileType_DIRECTORY:
		{
			return &pb.AddResponse{Message: "success"}, nil
		}

	case pb.FileType_SYSTEMLINK:
		{
			return &pb.AddResponse{Message: "success"}, nil
		}
	default:
		return nil, errors.New("Unsupported file type! ")
	}

	return &pb.AddResponse{Message: "I want to add " + request.FileName}, nil
}

func (service *addService) TransLargeFile(stream pb.AddTask_TransLargeFileServer) error {
	header, ok := metadata.FromIncomingContext(stream.Context())
	if !ok || len(header[StreamSessionIDKey]) == 0 {
		return errors.New("unknown stream without session info")
	}

	sessionId := header[StreamSessionIDKey][0]

	request := service.fileAddTask[sessionId]
	if request == nil {
		return errors.New(fmt.Sprintf("can't find the "+
			"request of this session: %s", sessionId))
	}

	if cs := request.SplitterSize; cs == 0 {
		request.SplitterSize = SplitterSize
	}

	reader := &streamReader{
		stream:    stream,
		service:   service,
		sessionId: sessionId,
	}

	importer := &RpcFileImporter{
		splitterSize: request.SplitterSize,
		reader:       reader,
		fileName:     request.FileName,
		fullPath:     request.FullPath,
		isDirectory:  false,
		Out:          make(chan *pb.AddResponse),
	}

	err := rpcServiceImpl.ImportFile(importer)

	return err
}
func (service *addService) appendTask(sessionId string, request *pb.AddRequest) {

	service.taskLock.Lock()
	defer service.taskLock.Unlock()
	service.fileAddTask[sessionId] = request
}

func (service *addService) removeTask(sessionId string) {

	service.taskLock.Lock()
	defer service.taskLock.Unlock()
	delete(service.fileAddTask, sessionId)
}

/*****************************************************************
*
*		simple file reader implement
*
*****************************************************************/
type fileReader struct {
	reader io.Reader
}

func (r *fileReader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
func (r *fileReader) Close(Out chan *pb.AddResponse) error {
	return nil
}

/*****************************************************************
*
*		stream reader implement
*
*****************************************************************/
type streamReader struct {
	stream    pb.AddTask_TransLargeFileServer
	sessionId string
	service   *addService
	reminder  []byte
	err       error
}

func (s *streamReader) Read(p []byte) (int, error) {

	pLen := len(p)
	if pLen == 0 {
		return 0, nil
	}

	if s.err != nil {
		return 0, s.err
	}

	var copied = 0
	for copied < pLen {

		if s.reminder != nil && len(s.reminder) > 0 {

			dataSize := copy(p[copied:], s.reminder)
			copied += dataSize

			if dataSize == len(s.reminder) {
				s.reminder = nil
			} else {
				s.reminder = s.reminder[dataSize:]
			}
		} else {
			fileChunk, err := s.stream.Recv()

			if err != nil {
				if err != io.EOF {
					return 0, err
				} else {
					s.err = io.EOF
					return copied, nil
				}
			}

			dataSize := copy(p[copied:], fileChunk.Content)
			copied += dataSize

			if dataSize < len(fileChunk.Content) {
				s.reminder = make([]byte, len(fileChunk.Content)-dataSize)
				copy(s.reminder, fileChunk.Content[dataSize:])
			}
		}
	}

	return copied, nil
}

func (s *streamReader) Close(Out chan *pb.AddResponse) error {

	result := <-Out

	s.stream.SendAndClose(result)

	s.service.removeTask(s.sessionId)

	return nil
}

/*****************************************************************
*
*		rpc file importer tool.
*
*****************************************************************/

type RpcReadCloser interface {
	io.Reader
	Close(chan *pb.AddResponse) error
}

type RpcFileImporter struct {
	reader       RpcReadCloser
	fileName     string
	fullPath     string
	isDirectory  bool
	splitterSize int32
	Out          chan *pb.AddResponse
}

func (importer *RpcFileImporter) NextChunk() ([]byte, error) {

	chunk := make([]byte, importer.splitterSize)
	realSize, err := importer.reader.Read(chunk)

	if realSize > 0 {
		return chunk[:realSize], nil
	} else {
		return nil, err
	}
}

func (importer *RpcFileImporter) FileName() string {
	return importer.fileName
}

func (importer *RpcFileImporter) FullPath() string {
	return importer.fullPath
}

func (importer *RpcFileImporter) IsDirectory() bool {
	return importer.isDirectory
}

func (importer *RpcFileImporter) NextFile() (rpcServiceImpl.FileImporter, error) {
	return nil, io.EOF
}

func (importer *RpcFileImporter) Close() error {
	return importer.reader.Close(importer.Out)
}

func (importer *RpcFileImporter) ResultCh() chan *pb.AddResponse {
	return importer.Out
}
