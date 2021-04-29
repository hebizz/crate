package handler

import (
  "bytes"
  "context"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "time"

  pb "Mcu-gin/interface"
  "Mcu-gin/serial"
  "Mcu-gin/ymodem"
  log "k8s.io/klog"
)

func PrintCommandExecuteLogs(cmd *exec.Cmd) error {
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    log.Error(err, stderr.String())
    return err
  }
  log.Info(out.String())
  return nil
}

func Senddlmodule(dlname string) error {
  Port := "/dev/ttyUSB3"
  // Set port
  connection, err := serial.OpenPort(&serial.Config{Name: Port, Baud: 115200})
  f, err := os.OpenFile(Port, os.O_RDWR, 0666)
  if err != nil {
    log.Error(err)
    return err
  }
  // Send ymodem signal
  _, err = f.WriteString("ymodem_rec uart1\n")
  if err != nil {
    log.Error(err)
    return err
  }

  time.Sleep(time.Duration(1000) * time.Millisecond)

  buffer := make([]byte, 100)
  _, err = f.Read(buffer)
  if err != nil {
    log.Error(err)
    return err
  }

  fmt.Printf("%s\n", string(buffer))
  f.Close()

  fIn, err := os.Open(dlname)
  if err != nil {
    log.Error(err)
    return err
  }
  data, err := ioutil.ReadAll(fIn)
  if err != nil {
    log.Error(err)
    return err
  }
  fIn.Close()

  // Send file
  if err = ymodem.ModemSend(connection, data, dlname); err != nil {
    log.Error(err)
    return err
  }

  // Flush Port
  f, _ = os.Open(Port)
  message := make([]byte, 100)
  if _, err = f.Read(message); err == nil {
    log.Info("send ok")
    log.Info(string(message))
  }
  f.Close()
  return nil
}

func (s *Server) CompileProject(ctx context.Context, request *pb.CompileRequest) (*pb.Reply, error) {
  dlname := request.Client + ".mo"
  log.Info(dlname)
  cmd := exec.Command("./compile.sh", request.Client, dlname)
  err := PrintCommandExecuteLogs(cmd)
  if err != nil {
    return &pb.Reply{Message: ""}, err
  }
  err = Senddlmodule(dlname)
  log.Info(err)
  if err != nil {
    return &pb.Reply{Message: ""}, err
  }
  return &pb.Reply{Message: "compile success"}, nil
}
