package main

import (
  "context"
  "os"
  "time"
  "fmt"

  pb "Mcu-gin/interface"
  "google.golang.org/grpc"
  log "k8s.io/klog"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Error("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewGreeterClient(conn)
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
  defer cancel()

  if len(os.Args) == 3 && os.Args[1] == "create" {
    MkdirProject(c, ctx, os.Args[2])
    os.Exit(0)
  }

  if len(os.Args) == 3 && os.Args[1] == "upload" {
    UploadFile(c, ctx, os.Args[2])
    os.Exit(0)
  }

  if len(os.Args) == 3 && os.Args[1] == "compile" {
    CompileProject(c, ctx, os.Args[2])
    os.Exit(0)
  }

  if len(os.Args) == 1 {
    log.Info("Please input correct argv")
    os.Exit(-1)
  } else {
    var shell string
    for index, value  :=  range os.Args {
      if index == 0 {
        continue
      } else {
        shell += (value + " ")
      }
    }
    ShellCommand(c, ctx, shell)
  }
}

func MkdirProject(c pb.GreeterClient, ctx context.Context, ProjectName string) {
  r, err := c.MkdirProject(ctx, &pb.MkdirRequest{Name: ProjectName})
  if err != nil {
    log.Error(err)
  }
  fmt.Println(r.Message)
}

func UploadFile(c pb.GreeterClient, ctx context.Context, path string) {
  r, err := c.UploadFile(ctx, &pb.UploadRequest{File: path})
  if err != nil {
    log.Error(err)
  }
  fmt.Println(r.Message)
}

func CompileProject(c pb.GreeterClient, ctx context.Context, ProjectName string) {
  r, err := c.CompileProject(ctx, &pb.CompileRequest{Client: ProjectName})
  if err != nil {
    log.Error(err)
  }
  fmt.Println(r.Message)
}

func ShellCommand(c pb.GreeterClient, ctx context.Context, command string) {
  r, err := c.ShellCommand(ctx, &pb.ShellCommandRequest{Command:command})
  if err != nil {
    log.Error(err)
  }
  fmt.Println(r.Message)
}

