package gonvm

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/clozed2u/gonvm/utils"
)

// exist check given path exist or not
func exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// Init create gonvm in user home dir if not exist
func Init() error {
	gonvmPath, err := getGoNVMPath()
	if err != nil {
		return err
	}
	if !exist(gonvmPath) {
		err = os.MkdirAll(gonvmPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	gonvmSourcesPath, err := getGoNVMSourcesPath()
	if err != nil {
		return err
	}
	if !exist(gonvmSourcesPath) {
		err = os.MkdirAll(gonvmSourcesPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	gonvmVersionsPath, err := getGoNVMVersionsPath()
	if err != nil {
		return err
	}
	if !exist(gonvmVersionsPath) {
		err = os.MkdirAll(gonvmVersionsPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func getDirName(version string) string {
	nodeOS := runtime.GOOS
	goArch := runtime.GOARCH
	nodeArch := ""
	switch goArch {
	case "amd64":
		nodeArch = "x64"
	case "386":
		nodeArch = "x86"
	default:
		nodeArch = goArch
	}
	return fmt.Sprintf("node-v%s-%s-%s", version, nodeOS, nodeArch)
}

func getFileName(version string) string {
	dirName := getDirName(version)
	return fmt.Sprintf("%s.tar.gz", dirName)
}

func getURL(version string) string {
	return fmt.Sprintf("https://nodejs.org/dist/v%s/%s", version, getFileName(version))
}

func getGoNVMPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gonvm"), nil
}

func getGoNVMBinPath() (string, error) {
	goNVMPath, err := getGoNVMPath()
	if err != nil {
		return "", nil
	}
	return filepath.Join(goNVMPath, "bin"), nil
}

func getGoNVMSourcesPath() (string, error) {
	goNVMPath, err := getGoNVMPath()
	if err != nil {
		return "", nil
	}
	return filepath.Join(goNVMPath, "sources"), nil
}

func getGoNVMVersionsPath() (string, error) {
	goNVMPath, err := getGoNVMPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(goNVMPath, "versions"), nil
}

func getNodeSourcePath(version string) (string, error) {
	srcPath, err := getGoNVMSourcesPath()
	if err != nil {
		return "", nil
	}
	fileName := getFileName(version)
	return filepath.Join(srcPath, fileName), nil
}

func getNodeVersionPath(version string) (string, error) {
	goNVMVersionPath, err := getGoNVMVersionsPath()
	if err != nil {
		return "", err
	}
	nodeDirName := getDirName(version)
	return filepath.Join(goNVMVersionPath, nodeDirName), nil
}

func getNodeVersionBinPath(version string) (string, error) {
	nodeVersionPath, err := getNodeVersionPath(version)
	if err != nil {
		return "", nil
	}
	return filepath.Join(nodeVersionPath, "bin"), nil
}

// Download specify node version into gonvm sources path
func Download(version string) error {
	dst, err := getNodeSourcePath(version)
	if err != nil {
		return fmt.Errorf("Failed to download node version %s with error: %s", version, err.Error())
	}
	url := getURL(version)
	err = utils.Download(url, dst)
	if err != nil {
		return fmt.Errorf("Failed to download node version %s with error: %s", version, err.Error())
	}
	return nil
}

// Unarchive specify node version into gonvm versions path
func Unarchive(version string) error {
	src, err := getNodeSourcePath(version)
	if err != nil {
		return fmt.Errorf("Failed to unarchive node version %s with error: %s", version, err.Error())
	}
	dst, err := getGoNVMVersionsPath()
	if err != nil {
		return fmt.Errorf("Failed to unarchive node version %s with error: %s", version, err.Error())
	}
	err = utils.Unarchive(src, dst)
	if err != nil {
		return fmt.Errorf("Failed to unarchive node version %s with error: %s", version, err.Error())
	}
	return nil
}

// Link specify node version to gonvm bin path
func Link(version string) error {
	goNVMBinPath, err := getGoNVMBinPath()
	if err != nil {
		return fmt.Errorf("Failed to link node version %s with error: %s", version, err.Error())
	}
	err = os.RemoveAll(goNVMBinPath)
	if err != nil {
		return fmt.Errorf("Failed to link node version %s with error: %s", version, err.Error())
	}
	nodeVersionBinPath, err := getNodeVersionBinPath(version)
	if err != nil {
		return fmt.Errorf("Failed to link node version %s with error: %s", version, err.Error())
	}
	err = os.Symlink(nodeVersionBinPath, goNVMBinPath)
	if err != nil {
		return fmt.Errorf("Failed to link node version %s with error: %s", version, err.Error())
	}
	return nil
}

// Use specify node version by doing necessary steps
func Use(version string) error {
	err := Init()
	if err != nil {
		return err
	}
	downloadDst, err := getNodeSourcePath(version)
	if err != nil {
		return fmt.Errorf("Failed to check downloaded node version %s", version)
	}
	if !exist(downloadDst) {
		err = Download(version)
		if err != nil {
			return err
		}
	}
	nodeVersionPath, err := getNodeVersionPath(version)
	if err != nil {
		return fmt.Errorf("Failed to check extracted node version %s", version)
	}
	if !exist(nodeVersionPath) {
		err = Unarchive(version)
		if err != nil {
			return err
		}
	}
	err = Link(version)
	if err != nil {
		return err
	}
	return nil
}

// List all installed versions of node
func List() ([]string, error) {
	versions := []string{}
	versionsPath, err := getGoNVMVersionsPath()
	if err != nil {
		return versions, err
	}
	versionsPathWithNodePrefix := fmt.Sprintf("%s/node-*", versionsPath)
	nodes, err := filepath.Glob(versionsPathWithNodePrefix)
	if err != nil {
		return versions, err
	}
	for _, node := range nodes {
		nodeWithoutVersionsPath := strings.ReplaceAll(node, versionsPath+"/", "")
		versionWithVPrefix := strings.Split(nodeWithoutVersionsPath, "-")[1]
		versionWithoutVPrefix := strings.ReplaceAll(versionWithVPrefix, "v", "")
		versions = append(versions, versionWithoutVPrefix)
	}
	return versions, nil
}
