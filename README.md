# mapkha-cli
mapkha-cli is a command line tool for Mapkha - Thai word segmentation (wordcut; word boundary identification; ตัดคำ) program in Go (golang)

## Example

     echo กากากากา | go run mapkha_cli 

     echo กากากากา | go run mapkha_cli -dix /mypath/mydix.txt
