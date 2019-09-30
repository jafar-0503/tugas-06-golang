package library

import "fmt"


func Filepublic(){
  var x1 mahasiswa
  x1.nama = "Jafar Verhoeten"
  x1.umur = 27

  fmt.Println("Nama Saya adalah", x1.nama)
  fmt.Println("Umur Saya adalah", x1.umur,"Tahun")
  fmt.Println("Terima Kasih Bosqueee")
}

type mahasiswa struct{
  nama string
  umur int
}
