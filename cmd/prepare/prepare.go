package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/liennie/AdventOfCode/pkg/log"
)

//go:embed .template
var template embed.FS

var now time.Time

func init() {
	loc, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}
	now = time.Now().In(loc)
}

func logCreate(path string) {
	log.Printf("🟢 Creating %s", path)
}

func logSkip(path string, reason string) {
	log.Printf("🔵 Skipping %s: %s", path, reason)
}

func prepDay() (int, int, error) {
	args := os.Args[1:]
	if len(args) > 2 {
		return 0, 0, fmt.Errorf("invalid number of arguments")
	}

	year, month, day := now.Date()

	if month != time.December || day > 25 {
		day = 0
	}

	var err error
	if len(args) == 2 {
		year, err = strconv.Atoi(args[0])
		if err != nil {
			return 0, 0, fmt.Errorf("invalid year argument: %w", err)
		}
		args = args[1:]
	}

	if len(args) == 1 {
		day, err = strconv.Atoi(args[0])
		if err != nil {
			return 0, 0, fmt.Errorf("invalid day argument: %w", err)
		}
		args = args[1:]
	}

	return year, day, nil
}

func mkDir(year, day int) error {
	dirPath := filepath.Join(strconv.Itoa(year), fmt.Sprintf("%02d", day))

	logCreate(dirPath)
	err := os.MkdirAll(dirPath, 0775)
	if err != nil {
		return fmt.Errorf("mkpath %s: %w", dirPath, err)
	}

	sub, err := fs.Sub(template, ".template")
	if err != nil {
		panic(fmt.Errorf("template sub: %w", err))
	}
	err = fs.WalkDir(sub, ".", func(srcPath string, entry fs.DirEntry, err error) error {
		if srcPath == "." {
			return nil
		}

		dstPath := srcPath
		if srcPath == "00.go" {
			dstPath = fmt.Sprintf("%02d.go", day)
		}

		localPath, err := filepath.Localize(dstPath)
		if err != nil {
			return fmt.Errorf("localize %q: %w", dstPath, err)
		}

		dstPath = filepath.Join(dirPath, localPath)

		if _, err := os.Stat(dstPath); err == nil {
			logSkip(dstPath, "already exists")
			return nil
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("stat %s: %w", dstPath, err)
		}

		logCreate(dstPath)
		if entry.IsDir() {
			err := os.MkdirAll(dstPath, 0775)
			if err != nil {
				return fmt.Errorf("mkpath %s: %w", dstPath, err)
			}
		} else {
			src, err := sub.Open(srcPath)
			if err != nil {
				return fmt.Errorf("open src %s: %w", srcPath, err)
			}
			defer src.Close()

			info, err := src.Stat()
			if err != nil {
				return fmt.Errorf("stat src %s: %w", srcPath, err)
			}

			dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666|info.Mode()&0777)
			if err != nil {
				return fmt.Errorf("create %s: %w", dstPath, err)
			}
			defer dst.Close()

			if _, err := io.Copy(dst, src); err != nil {
				return fmt.Errorf("copy %s: %w", dstPath, err)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func downloadInput(year, day int) error {
	dstPath := filepath.Join(strconv.Itoa(year), fmt.Sprintf("%02d", day), "input.txt")

	if _, err := os.Stat(dstPath); err == nil {
		logSkip(dstPath, "already exists")
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("stat %s: %w", dstPath, err)
	}

	nYear, _, nDay := now.Date()
	if year > nYear || (year == nYear && day > nDay) {
		logSkip(dstPath, "not ready")
		return nil
	}

	token := os.Getenv("SESSION_TOKEN")
	if token == "" {
		logSkip(dstPath, "SESSION_TOKEN not set")
		return nil
	}

	// https://www.reddit.com/r/adventofcode/wiki/faqs/automation#wiki_put_your_contact_info_in_your_script.27s_user-agent_header
	ua := os.Getenv("USER_AGENT")
	if ua == "" {
		logSkip(dstPath, "USER_AGENT not set")
		return nil
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("User-Agent", ua)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: token,
	})

	cli := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Printf("🟣 Downloading %s", url)
	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			msg = []byte(err.Error())
		}
		return fmt.Errorf("http status %d: %s: %s", resp.StatusCode, resp.Status, string(msg))
	}

	logCreate(dstPath)
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create %s: %w", dstPath, err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, resp.Body); err != nil {
		return fmt.Errorf("copy %s: %w", dstPath, err)
	}

	return nil
}

func main() {
	year, day, err := prepDay()
	if err != nil {
		log.Printf("🔴 Error parsing aruments: %v", err)
		os.Exit(1)
		return
	}
	if year < 2015 || year > 3000 {
		log.Printf("🔴 Invalid year %d", year)
		os.Exit(1)
		return
	}
	if day == 0 {
		log.Println("🔴 Please specify which day to preprare")
		os.Exit(1)
		return
	}
	if day < 0 || day > 25 {
		log.Printf("🔴 Invalid day %d", day)
		os.Exit(1)
		return
	}

	err = mkDir(year, day)
	if err != nil {
		log.Printf("🔴 Error preparing directory: %v", err)
		os.Exit(2)
		return
	}

	err = downloadInput(year, day)
	if err != nil {
		log.Printf("🔴 Error downloading input: %v", err)
		os.Exit(3)
		return
	}
}
