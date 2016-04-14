package acceptance_test

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
	"github.com/cloudfoundry-incubator/volman/certification"
)

var (
	volmanPath         string
	volmanServerPort   int
	debugServerAddress string

	driverPath           string
	driverServerPort     int
	driverServerPortJson int

	mountDir       string
	tmpDriversPath string

	socketPath string
	fixturesFuncMapFunc func() (map[string]func()(certification.VolmanFixture, certification.DriverFixture))
)

func TestVolman(t *testing.T) {
	// these integration tests can take a bit, especially under load;
	// 1 second is too harsh
	SetDefaultEventuallyTimeout(10 * time.Second)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Fakedriver Certification Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	vPath, err := gexec.Build("github.com/cloudfoundry-incubator/volman/cmd/volman", "-race")
	Expect(err).NotTo(HaveOccurred())

	dPath, err := gexec.Build("github.com/cloudfoundry-incubator/volman/fakedriver/cmd/fakedriver", "-race")
	Expect(err).NotTo(HaveOccurred())

	return []byte(strings.Join([]string{vPath, dPath}, ","))
	}, func(pathsByte []byte) {
		path := string(pathsByte)
		volmanPath = strings.Split(path, ",")[0]
		driverPath = strings.Split(path, ",")[1]

		fmt.Printf("===>volmanPath: %s, driverPath: %s\n\n", volmanPath, driverPath)
})

var _ = BeforeEach(func() {
	var err error

	tmpDriversPath, err = ioutil.TempDir(os.TempDir(), "volman-cert-test")
	Expect(err).ShouldNot(HaveOccurred())

	mountDir, err = ioutil.TempDir("", "mountDir")
	Expect(err).NotTo(HaveOccurred())

	driverServerPort = 9750 + GinkgoParallelNode()
	driverServerPortJson = 9750 + 100 + GinkgoParallelNode()

	socketPath = path.Join(tmpDriversPath, "fakedriver.sock")

	volmanServerPort = 8750 + GinkgoParallelNode()
	debugServerAddress = fmt.Sprintf("0.0.0.0:%d", 8850+GinkgoParallelNode())

	fixturesPath, err := GetOrCreateFixturesPath()
	Expect(err).NotTo(HaveOccurred())

	fmt.Printf("===>Save fixturesPath: %#v\n\n", fixturesPath)

	transports := []string{"tcp"} //, "tcp-json", "unix"}
	driverListenAddresses := []string{
		fmt.Sprintf("0.0.0.0:%d", driverServerPort),
		fmt.Sprintf("0.0.0.0:%d", driverServerPortJson),
		socketPath,
	}
	ports := []int {driverServerPort, driverServerPortJson, 0}
	for i, transport := range transports {
		certificationFixture := certification.CertificationFixture{
			PathToVolman: volmanPath,
			VolmanFixture: certification.VolmanFixture{
				Config: certification.VolmanConfig{
					ServerPort: volmanServerPort,
					DebugServerAddress: debugServerAddress,
					DriversPath: tmpDriversPath,
					ListenAddress: fmt.Sprintf("0.0.0.0:%d", volmanServerPort),
				},
			},
			PathToDriver: driverPath,
			MountDir: mountDir,
			Transport: transport,
			DriverFixture: certification.DriverFixture{
				Config: certification.DriverConfig{
					Name: "fakedriver",
					Path: tmpDriversPath,
					ServerPort: ports[i],
					ListenAddress: driverListenAddresses[i],
				},
				VolumeData: certification.VolumeData{
					Name: "fake-volume-name",
					Config: map[string]interface{}{
						"volume_id": "fake-volume-name",
					},
				},
			},
		}

		fmt.Printf("===>certificationFixture: %#v\n\n", certificationFixture)
		fileName := filepath.Join(fixturesPath, fmt.Sprintf("certification-%s.json", transport))
		err := certification.SaveCertificationFixture(certificationFixture, fileName)
		Expect(err).NotTo(HaveOccurred())
	}
})

//var _ = SynchronizedAfterSuite(func() {
//}, func() {
//	gexec.CleanupBuildArtifacts()
//})

func GetOrCreateFixturesPath() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fixturesPath := filepath.Join(workingDir, "..", "..", "fixtures", "auto")

	err = os.MkdirAll(fixturesPath, 0700)
	if err != nil {
		return "", err
	}

	return fixturesPath, nil
}
