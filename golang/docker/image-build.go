package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	fmt.Println(cli.ClientVersion())
	fmt.Println(cli.ServerVersion(ctx))

	tarFile := "/Users/jemy/Temp/cloudctl-lab-bundle/cloudctl-lab/app/echo-go/echo.tar.gz"
	tarFp, openErr := os.Open(tarFile)
	if openErr != nil {
		return
	}
	defer tarFp.Close()
	resp, err := cli.ImageBuild(ctx, tarFp, types.ImageBuildOptions{
		Tags:       []string{"echo-go:1.2"},
		Dockerfile: "./Dockerfile",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.OSType)
	io.Copy(os.Stdout, resp.Body)
}
