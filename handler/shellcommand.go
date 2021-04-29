package handler

import (
  "context"
  "fmt"
  "os"
  "strings"
  "time"

  "Mcu-gin/serial"

  pb "Mcu-gin/interface"
  log "k8s.io/klog"
)

func (s *Server) ShellCommand(ctx context.Context, in *pb.ShellCommandRequest) (*pb.Reply, error) {
  _, err := serial.OpenPort(&serial.Config{Name: "/dev/ttyUSB3", Baud: 115200})
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
//  if in.Module == "dlmodule" {
//    switch in.Method {
//    case "start":
//      return Startdlmodule(in.Object)
//   case "stop":
//      return Stopdlmodule(in.Object)
//    case "check":
//      return Checkdlmodule(in.Object)
//    case "rm":
//      return Rmdlmodule(in.Object)
//    case "list":
//      return Listdlmodule()
//    case "auto":
//      return Autodlmodule(in.Object, in.Mode)
//    default:
//      return &pb.Reply{Message: "argv error"}, nil
//    }
//  } else {
  return OtherCommand(in.Command)
//  }
}

func OtherCommand(shell string) (*pb.Reply, error) {
  f, _ := os.OpenFile("/dev/ttyUSB3", os.O_RDWR, 0666)
  _, err := f.WriteString((shell + "\n"))
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(450) * time.Millisecond)
  var data string
  for n := 1024 ; n == 1024 ; {
    buffer := make([]byte, 1024)
    n, err = f.Read(buffer)
    if err != nil {
      log.Error(err)
      return &pb.Reply{Message: ""}, err
    }
    data += string(buffer)
    time.Sleep(time.Duration(1000) * time.Millisecond)
  }
  rawlen := len((shell + "\n"))
  reslen := len(data)
//  log.Info(data[rawlen+1:reslen])
  return &pb.Reply{Message: data[rawlen+1:reslen]}, nil
}

func Startdlmodule(object string) (*pb.Reply, error) {
  f, _ := os.OpenFile("/dev/ttyUSB3", os.O_RDWR, 0666)
  command := "dlmodule start " + object + "\n"
  _, err := f.WriteString(command)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(200) * time.Millisecond)
  buffer := make([]byte, 128)
  n, err := f.Read(buffer)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  length := len(command) + 1
  log.Info("the reply is\n", string(buffer), n)
  signal := string(buffer[length : n-6])
  log.Info("the signal is \n", len(signal), signal)
  if strings.Contains(signal, "start fail") {
    return &pb.Reply{Message: signal}, nil
  } else if len(signal) == 0 {
    return &pb.Reply{Message: fmt.Sprintf("dlmodule start %s success!", object)}, nil
  } else {
    return &pb.Reply{Message: signal}, nil
  }
}

func Stopdlmodule(object string) (*pb.Reply, error) {
  f, _ := os.OpenFile("/dev/ttyUSB3", os.O_RDWR, 0666)
  command := "dlmodule stop " + object + "\n"
  _, err := f.WriteString(command)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(200) * time.Millisecond)
  buffer := make([]byte, 128)
  n, _ := f.Read(buffer)
  // right result
  length := len(command) + 1
  log.Info("the reply is", string(buffer), n)
  signal := string(buffer[length : n-6])
  log.Info("the signal is \n", len(signal), signal)
  if strings.Contains(signal, "stop fail") {
    return &pb.Reply{Message: signal}, nil
  } else if len(signal) == 0 {
    return &pb.Reply{Message: fmt.Sprintf("dlmodule stop %s success!", object)}, nil
  } else {
    return &pb.Reply{Message: signal}, nil
  }
}

func Checkdlmodule(object string) (*pb.Reply, error) {
  f, _ := os.OpenFile("/dev/ttyUSB3", os.O_RDWR, 0666)
  command := "dlmodule check " + object + "\n"
  _, err := f.WriteString(command)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(200) * time.Millisecond)
  // flush message
  buffer := make([]byte, 128)
  n, _ := f.Read(buffer)
  length := len(command) + 1
  log.Info("the reply is", string(buffer), n)
  // recv /r/n msh->
  signal := string(buffer[length : n-8])
  log.Info("size is ", len(signal), signal)
  if (signal == "closed") || (signal == "running") {
    return &pb.Reply{Message: fmt.Sprintf("dlmodule %s  is  %s", string(object), signal)}, nil
  } else {
    return &pb.Reply{Message: fmt.Sprintf("dlmodule check %s failed!", object)}, nil
  }
}

func Rmdlmodule(object string) (*pb.Reply, error) {
  f, _ := os.OpenFile("/dev/ttyUSB3", os.O_RDWR, 0666)
  command := "dlmodule rm " + object + "\n"
  _, err := f.WriteString(command)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(200) * time.Millisecond)
  buffer := make([]byte, 128)
  n, _ := f.Read(buffer)
  length := len(command) + 1
  log.Info("the reply is::", string(buffer), n)
  signal := string(buffer[length : n-6])
  log.Info("size is::", len(signal), signal)
  if strings.Contains(signal, "rm fail") {
    return &pb.Reply{Message: signal}, nil
  } else if len(signal) == 0 {
    return &pb.Reply{Message: fmt.Sprintf("dlmodule rm %s success!", object)}, nil
  } else {
    return &pb.Reply{Message: signal}, nil
  }
}

func Listdlmodule() (*pb.Reply, error) {
  f, _ := os.OpenFile("/dev/ttyUSB3", os.O_RDWR, 0666)
  _, err := f.WriteString("dlmodule list\n")
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(200) * time.Millisecond)
  buffer := make([]byte, 500)
  n, _ := f.Read(buffer)
  log.Info("the size is", n, string(buffer))
  return &pb.Reply{Message: string(buffer)}, nil
}

func Autodlmodule(object string, mode string) (*pb.Reply, error) {
  f, _ := os.Create("/dev/ttyUSB3")
  command := "dlmodule auto " + object + " " + mode + "\n"
  _, err := f.WriteString(command)
  if err != nil {
    log.Error(err)
    return &pb.Reply{Message: ""}, err
  }
  defer f.Close()
  time.Sleep(time.Duration(200) * time.Millisecond)
  buffer := make([]byte, 128)
  n, _ := f.Read(buffer)
  length := len(command) + 1
  log.Info("the reply is::", string(buffer), n)
  signal := string(buffer[length : n-6])
  log.Info("the signal is::", signal, len(signal))
  if strings.Contains(signal, "auto fail") {
    return &pb.Reply{Message: signal}, nil
  } else if len(signal) == 0 {
    return &pb.Reply{Message: fmt.Sprintf("dlmodule auto %s %s success!", object, mode)}, nil
  } else {
    return &pb.Reply{Message: signal}, nil
  }
}
