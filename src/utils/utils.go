package utils

import (
	"fmt"
	"reflect"
	"time"
)


func TimeFunc(f interface{}, args ...interface{}){
  start := time.Now()
  in := make([]reflect.Value,len(args))
  for i,arg := range args{
    in[i] = reflect.ValueOf(arg)
  }
  reflect.ValueOf(f).Call(in)
  elapsed := time.Since(start)
  fmt.Printf("Func %v took %v",reflect.TypeOf(f),elapsed)
}

func Contains(arr []string,val string) bool {
  for _,v := range arr{
    if v == val{
      return true
    }
  }
  return false
}
