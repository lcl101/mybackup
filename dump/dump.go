package dump

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/lcl101/mybackup/option"
)

func Dump(opt *option.Options) {
	//https://www.cnblogs.com/qq78292959/p/3637135.html
	t := time.Now()
	backFile := fmt.Sprintf("%s/backdata%d", opt.BackupPath, t.Weekday())
	bfsql := backFile + ".sql"
	os.Remove(bfsql)
	var args []string
	args = append(args, fmt.Sprintf("-u%s", opt.UserName))
	args = append(args, fmt.Sprintf("-p%s", opt.Password))
	args = append(args, fmt.Sprintf("-h%s", opt.HostName))
	args = append(args, fmt.Sprintf("-P%s", opt.Port))
	// --single-transaction --master-data=2
	args = append(args, "--single-transaction")
	args = append(args, "--master-data=2")
	args = append(args, opt.Databases)
	args = append(args, fmt.Sprintf("-r%s", bfsql))

	cmd := exec.Command(opt.MySQLDumpPath, args...)
	cmdOut, _ := cmd.StdoutPipe()
	cmdErr, _ := cmd.StderrPipe()

	cmd.Start()

	output, _ := ioutil.ReadAll(cmdOut)
	err, _ := ioutil.ReadAll(cmdErr)
	cmd.Wait()

	fmt.Println("mysqldump output is : " + string(output))
	if err != nil {
		fmt.Println("mysqldump error is: " + string(err))
		// os.Exit(4)
	}

	bfgz := backFile + ".tar.gz"
	os.Remove(bfgz)

	file, errcreate := os.Create(bfgz)

	if errcreate != nil {
		fmt.Println("error to create a compressed file: " + backFile)
		os.Exit(4)
	}

	defer file.Close()
	// set up the gzip writer
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	Compress(tw, backFile)
	file.Sync()
	gw.Flush()
	tw.Flush()
}

func Compress(tw *tar.Writer, backFile string) error {

	bfsql := backFile + ".sql"

	file, err := os.Open(bfsql)
	if err != nil {
		return err
	}
	defer file.Close()
	// fmt.Println(file)
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		// header := new(tar.Header)
		// header.Name = bfsql
		// header.Size = stat.Size()
		// header.Mode = int64(stat.Mode())
		// header.ModTime = stat.ModTime()

		header, err := tar.FileInfoHeader(stat, "")
		if err != nil {
			return err
		}
		header.Name = bfsql

		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}

		// Removing the original file after zipping it
		err = os.Remove(bfsql)

		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
