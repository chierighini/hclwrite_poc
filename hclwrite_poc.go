package main

import (
  "fmt"
  "os"
  "github.com/hashicorp/hcl/v2"
  "github.com/hashicorp/hcl/v2/hclwrite"
  //"reflect"
  "github.com/zclconf/go-cty/cty"
)

func main() {
  tfFile, err := os.ReadFile("example.tf")
  if err != nil{
    fmt.Println(err)
    return
  }
  f, status := hclwrite.ParseConfig(tfFile,"",hcl.Pos{Line:1, Column:1})
  fmt.Println(status)
  rootBody := f.Body()
  resourceBody := rootBody.FirstMatchingBlock("resource",[]string{"aws_instance","app_server"}).Body()
  enclaveBody := resourceBody.FirstMatchingBlock("enclave_options", nil).Body()
  enclaveBody.SetAttributeValue("enabled", cty.BoolVal(true))
  newFile, err := os.Create("new_example.tf")
  if err != nil {
    return
  }
  newFile.Write(f.Bytes())
}
