package main

// v1
//func main() {
//	//log.Println(os.Getwd())
//	//err := os.Chdir("cmd")
//	//if err != nil {
//	//	log.Println(err)
//	//	return
//	//}
//
//	file, err := os.Create("write.txt")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	_, err = file.WriteString("transactions history")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	err = file.Close()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//}

// v2
//func main() {
//	file, err := os.Create("write.txt")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer file.Close()
//	defer func() {
//		if err := file.Close(); err != nil {
//			log.Println(err)
//		}
//	}()
//
//	_, err = file.WriteString("transactions history")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//}

// v3
//func main() {
//	file, err := os.Create("write.txt")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer func() {
//		if err := file.Close(); err != nil {
//			log.Println(err)
//		}
//	}()
//
//	_, err = file.WriteString("transactions history")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//}

//func main() {
//	file, err := os.Create("write.txt")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer func(c io.Closer) {
//		if err := c.Close(); err != nil {
//			log.Println(err)
//		}
//	}(file)
//
//	_, err = file.WriteString("transactions history")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//}
