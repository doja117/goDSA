package main

func insertIntoarr(p []int,q int)[]int{
  if len(p)==0{
    p=append(p,q)
    return p
  }
  for i:=0;i<len(p);i++ {
    if q>p[i]{
      var arr []int
      arr=append(arr,p[:i-1]...)
      arr=append(arr,q)
      arr=append(arr,p[i-1:]...)
      return arr
    }
  }
  var arr []int
  return arr
}
func main(){
  var ans []int
   var arr []int={5,1,4,3,2}
    for i:=0;i<len(arr);i++ {
      ans=insetIntoarr(arr,q)
    }
  for i:=0;i<len(ans);i++ {
    fmt.Printf("%v\t",ans[i])
  }
}
