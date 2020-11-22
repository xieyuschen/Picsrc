package main

//func getTimeStamp() (timestamp int64){
//	timestamp = time.Now().Unix()
//	return
//}
//
////User upload file should never be trusted
////so replace name with timestamp
//func ParseFileName(filename string) (newfilename string){
//	extension := filepath.Ext(filename)
//	timestamp_int := int(getTimeStamp())
//	newfilename = strconv.Itoa(timestamp_int)+extension
//	return
//}