package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	_ "kusionstack.io/kcl-plugin"
	_ "kusionstack.io/kclvm-go"
	"kusionstack.io/kusion/pkg/log"
	"kusionstack.io/kusion/pkg/version"
)

var (
	flagOutFile = flag.String("o", "z_update_version.go", "set output faile")
)

func main() {
	flag.Parse()
	if *flagOutFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	versionInfo, err := version.NewInfo()
	if err != nil {
		log.Fatal(err)
	}

	data := makeUpdateVersionGoFile(versionInfo)

	err = ioutil.WriteFile(*flagOutFile, []byte(data), 0o666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}

	fmt.Println(versionInfo.String())
}

func makeUpdateVersionGoFile(v *version.Info) string {
	return fmt.Sprintf(`// Auto generated by 'go run gen.go', DO NOT EDIT.

package version

func init() {
	info = &Info{
		ReleaseVersion: %q,
		GitInfo: &GitInfo{
			LatestTag:   %q,
			Commit:      %q,
			TreeState:   %q,
		},
		BuildInfo: &BuildInfo{
			GoVersion: %q,
			GOOS:      %q,
			GOARCH:    %q,
			NumCPU:    %d,
			Compiler:  %q,
			BuildTime: %q,
		},
		Dependency: &DependencyVersion{
			KclvmgoVersion:   %q,
			KclPluginVersion: %q,
		},
	}
}
`,
		v.ReleaseVersion,
		v.GitInfo.LatestTag,
		v.GitInfo.Commit,
		v.GitInfo.TreeState,
		v.BuildInfo.GoVersion,
		v.BuildInfo.GOOS,
		v.BuildInfo.GOARCH,
		v.BuildInfo.NumCPU,
		v.BuildInfo.Compiler,
		v.BuildInfo.BuildTime,
		v.Dependency.KclvmgoVersion,
		v.Dependency.KclPluginVersion,
	)
}
