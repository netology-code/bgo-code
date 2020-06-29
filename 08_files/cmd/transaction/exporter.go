package main

import (
	"io"
	"lectionchannels/pkg/transaction"
	"log"
	"os"
)

func main() {
	if err := execute("export.csv"); err != nil {
		os.Exit(1)
	}
}

//func execute(filename string) (err error) {
//	file, err := os.Create(filename)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func(c io.Closer) {
//		if cerr := c.Close(); cerr != nil {
//			log.Println(cerr)
//			if err == nil {
//				err = cerr
//			}
//		}
//	}(file)
//
//	svc := transaction.NewService()
//
//	_, err = svc.Register("0001", "0002", 10_000_00)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	err = svc.Export(file)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}


//// TODO: show in debugger
//type errCloser struct {}
//func (e *errCloser) Close() error {
//	return errors.New("sample error")
//}
//
//func execute(filename string) (err error) {
//	file, err := os.Create(filename)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	defer func(c io.Closer) {
//		if cerr := c.Close(); cerr != nil {
//			log.Println(cerr)
//			if err == nil {
//				err = cerr
//			}
//		}
//	}(&errCloser{})
//
//	svc := transaction.NewService()
//
//	_, err = svc.Register("0001", "0002", 10_000_00)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	err = svc.Export(file)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}

func execute(filename string) error {
	var err error
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	svc := transaction.NewService()

	if _, err = svc.Register("0001", "0002", 10_000_00); err != nil {
		log.Println(err)
		return err
	}

	if err = svc.Export(file); err != nil {
		log.Println(err)
		return err
	}
	return err
}


