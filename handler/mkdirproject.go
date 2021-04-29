package handler

import (
  "context"
  "io"
  "os"

  pb "Mcu-gin/interface"
  log "k8s.io/klog"
)

type Server struct{}

var Projectname string

func CopyFile(dstName, srcName string) (written int64, err error) {
  src, err := os.Open(dstName)
  if err != nil {
    return
  }
  defer src.Close()
  dst, err := os.OpenFile(srcName, os.O_WRONLY|os.O_CREATE, 0644)
  if err != nil {
    return
  }
  defer dst.Close()
  return io.Copy(dst, src)
}

func (s *Server) MkdirProject(ctx context.Context, in *pb.MkdirRequest) (*pb.Reply, error) {
  Projectname = in.Name
  log.Info("the name is ", Projectname)
  path := "RT_Thread_SDK/sdk/" + Projectname
  err := os.MkdirAll(path, os.ModePerm)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  _, err = CopyFile("RT_Thread_SDK/sdk/SConscript", path+"/SConscript")
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  } else {
    return &pb.Reply{Message:"mkdir success"}, nil
  }
}
