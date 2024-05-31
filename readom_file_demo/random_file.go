package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func writeFile(directory string, filename string, content string, maxSize int64) error {
	// 确保目录存在
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}

	// 构造文件路径
	filePath := filepath.Join(directory, filename)

	// 创建文件并写入内容
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入内容直到达到最大文件大小
	currentSize := int64(0)
	for currentSize < maxSize {
		n, err := file.WriteString(content)
		if err != nil {
			return err
		}
		currentSize += int64(n)
	}

	return nil
}

func generateSQLContent(table string, createUserCode string, additionalParams []string) string {
	conditions := []string{fmt.Sprintf("create_user_code = '%s'", createUserCode)}

	// 生成随机的 create_time 范围
	startTime, endTime := generateRandomTimeRange()
	conditions = append(conditions, fmt.Sprintf("create_time BETWEEN '%s' AND '%s'", startTime, endTime))

	for _, param := range additionalParams {
		conditions = append(conditions, param)
	}

	return fmt.Sprintf("SELECT * FROM %s WHERE %s;\n", table, strings.Join(conditions, " AND "))
}

func generateRandomTimeRange() (string, string) {
	// 生成一个随机的过去一年中的日期范围
	now := time.Now()
	start := now.AddDate(-1, 0, 0).Unix() // 一年前的时间戳
	end := now.Unix()                     // 当前时间的时间戳

	// 生成随机的开始和结束时间
	startTime := time.Unix(rand.Int63n(end-start)+start, 0)
	endTime := time.Unix(rand.Int63n(end-start)+start, 0)

	// 确保 startTime 小于 endTime
	if startTime.After(endTime) {
		startTime, endTime = endTime, startTime
	}

	return startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")
}

func main() {
	app := &cli.App{
		Name:  "File Writer",
		Usage: "Write files to a specified directory at regular intervals",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dir",
				Aliases:  []string{"d"},
				Usage:    "Directory to write files to",
				Required: true,
			},
			&cli.Int64Flag{
				Name:     "maxsize",
				Aliases:  []string{"m"},
				Usage:    "Maximum file size in bytes",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "table",
				Usage:    "SQL table name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "create_user_code",
				Usage:    "User code for creating entries",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:  "param",
				Usage: "Additional parameters for the SQL query (e.g., 'column=value')",
			},
			&cli.DurationFlag{
				Name:     "interval",
				Aliases:  []string{"i"},
				Usage:    "Interval between file writes (e.g., '10m' for 10 minutes)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			directory := c.String("dir")
			maxSize := c.Int64("maxsize")
			table := c.String("table")
			createUserCode := c.String("create_user_code")
			additionalParams := c.StringSlice("param")
			interval := c.Duration("interval")

			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			//ticker.C 是一个
			for range ticker.C {
				filename := fmt.Sprintf("file_%d.sql", time.Now().Unix()) // 根据时间戳生成文件名
				content := generateSQLContent(table, createUserCode, additionalParams)
				err := writeFile(directory, filename, content, maxSize)
				if err != nil {
					fmt.Printf("Error writing file: %v\n", err)
				} else {
					fmt.Printf("File %s written successfully at %v\n", filename, time.Now())
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
