package main

import(
    "fmt"
    "net"
)


func Scan(ip string, porta int) (string) {
    connection, err := net.Dial("tcp", ip)
    if err != nil{ return "" }

    connection.Close()
    fmt.Printf("Porta %d aberta\n",porta)
    return ip
}

func ValidIps() (interface{}){
    var validIps []*net.IPNet
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Print(fmt.Errorf("Endereços Locais: %+v\n", err.Error()))
        return -1
    }

    for _,i := range interfaces {
      ip,err := i.Addrs()
      if err != nil {
        fmt.Print(fmt.Errorf("Endereços Locais: %+v\n", err.Error()))
        continue
      }
      for _,a := range ip{
        switch v:= a.(type){
          case *net.IPNet:
            fmt.Printf("%v : %s (%s)\n",i.Name,v,v.IP.DefaultMask())
            validIps = append(validIps, v)
         }
       }
     }
     return validIps
}

func main(){
    var ipValidos []string
    iporta:="127.0.0.1:3306"
    porta:=3306
    result := Scan(iporta,porta)
    if result == "" { fmt.Println("Fudeu") }
    ipValidos = append(ipValidos,result)
    ips := ValidIps()
    fmt.Println(ips)
}
