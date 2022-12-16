package cmd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"sigs.k8s.io/cluster-api-provider-azure/azure/scope"
)

const (
	KubeloginDefault     = "kubelogin"
	KubeloginEnvVarName  = "KUBELOGIN_PATH"
	KubeConfigEnvVarName = "KUBECONFIG"
)

// GetKubeloginExecutablePath returns the location of the kubelogin binary.
func GetKubeloginExecutablePath() string {
	kubeloginPath, set := os.LookupEnv(KubeloginEnvVarName)
	if set {
		return kubeloginPath
	}
	return KubeloginDefault
}

func ExecCommand(cmdName string, cmdArgs ...string) (err error) {
	path, err := exec.LookPath(cmdName)
	if err != nil {
		return
	}
	cmd := exec.Command(path, cmdArgs...)
	var errb bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &errb
	cmd.Env = []string{fmt.Sprintf("HOME=%s", os.Getenv("HOME")), fmt.Sprintf("PATH=%s", os.Getenv("PATH"))}
	err = cmd.Run()
	if err != nil {
		return errors.New(errb.String())
	}
	return nil
}

// ConvertKubeConfig converts kube-config from interactive loging to non-interactive login using kubelogin.
// Expects kubelogin binary to be available at /bin/kubelogin
func ConvertKubeConfig(ctx context.Context, clusterName string, configData []byte, credentialsProvider *scope.ManagedControlPlaneCredentialsProvider) (config []byte, err error) {
	if credentialsProvider == nil {
		return nil, errors.New("cannot convert kubeconfig without credential provider")
	}
	fileName := fmt.Sprintf("%s/kubeconfig-%s", os.TempDir(), clusterName)
	klog.V(2).Infof("Temp file name: %s", fileName)
	kubeConfigToRestore, set := os.LookupEnv(KubeConfigEnvVarName)
	defer func() {
		if set {
			os.Setenv(KubeConfigEnvVarName, kubeConfigToRestore)
		} else {
			os.Unsetenv(KubeConfigEnvVarName)
		}
		os.Remove(fileName)
	}()
	os.Setenv(KubeConfigEnvVarName, fileName)
	err = os.WriteFile(fileName, configData, 0644)
	klog.V(2).Infof("Wrote config to file: %s", fileName)
	if err != nil {
		return nil, err
	}
	clientSecret, err := credentialsProvider.GetClientSecret(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get client secret")
	}
	kubeLogin := GetKubeloginExecutablePath()
	err = ExecCommand(kubeLogin,
		"convert-kubeconfig", "-l", "spn", "--client-id", credentialsProvider.GetClientID(), "--client-secret",
		clientSecret, "--tenant-id", credentialsProvider.GetTenantID(), "--kubeconfig", os.Getenv(KubeConfigEnvVarName))
	if err != nil {
		return nil, errors.Wrap(err, "could not convert kubeconfig to non interactive format")
	}
	klog.V(2).Infof("Converted kubeconfig: %s", fileName)
	return os.ReadFile(fileName)
}
