package handler

import (
  "context"
  "os"
  "strings"

  pb "Mcu-gin/interface"
  log "k8s.io/klog"
)


func (s *Server) UploadFile(ctx context.Context, in *pb.UploadRequest) (*pb.Reply, error) {

  data := strings.Split(in.File, "/")
  log.Info(data[len(data)-1])
  path := "RT_Thread_SDK/sdk/" + Projectname + "/" + data[len(data)-1]
  log.Info(path)
  CopyFile(in.File, path)
  if _, err := os.Stat(path); err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  } else {
    return &pb.Reply{Message: "upload file success"}, nil
  }
}
