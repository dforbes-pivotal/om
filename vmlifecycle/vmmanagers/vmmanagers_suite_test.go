package vmmanagers_test

import (
	"log"
	"regexp"
	"testing"

	"os"
	"path/filepath"

	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/fatih/color"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/om/vmlifecycle/vmmanagers"
)

func TestVMManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VMManagers Suite")
}

var _ = BeforeSuite(func() {
	log.SetOutput(GinkgoWriter)
	pathToStub, err := gexec.Build("github.com/pivotal-cf/om/vmlifecycle/vmmanagers/stub")
	Expect(err).ToNot(HaveOccurred())

	tmpDir := filepath.Dir(pathToStub)
	os.Setenv("PATH", tmpDir+":"+os.Getenv("PATH"))

	govcPath := tmpDir + "/govc"
	gcloudPath := tmpDir + "/gcloud"
	omPath := tmpDir + "/om"
	err = os.Link(pathToStub, govcPath)
	Expect(err).ToNot(HaveOccurred())

	err = os.Link(pathToStub, omPath)
	Expect(err).ToNot(HaveOccurred())

	err = os.Link(pathToStub, gcloudPath)
	Expect(err).ToNot(HaveOccurred())

	color.NoColor = true
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func testIAASForPropertiesInExampleFile(iaas string) {
	It("has an example file the represents all the correct fields", func() {
		exampleFile, err := ioutil.ReadFile(fmt.Sprintf("../../../docs-platform-automation/docs/examples/opsman-config/%s.yml", strings.ToLower(iaas)))
		Expect(err).ToNot(HaveOccurred())

		isolateCommentedParamRegex := regexp.MustCompile(`(?m)^(\s+)# ([\w-]+: )`)
		exampleFile = isolateCommentedParamRegex.ReplaceAll(exampleFile, []byte("$1$2"))

		config := vmmanagers.OpsmanConfigFilePayload{}
		err = yaml.UnmarshalStrict(exampleFile, &config)
		Expect(err).ToNot(HaveOccurred())

		configStruct := reflect.ValueOf(config.OpsmanConfig)
		iaasPtrStruct := configStruct.FieldByName(iaas)
		iaasStruct := iaasPtrStruct.Elem()

		Expect(iaasStruct.NumField()).To(BeNumerically(">", 0))

		testPropertiesExist(iaasStruct)
	})
}

func testPropertiesExist(vst reflect.Value) {
	tst := vst.Type()
	for i := 0; i < vst.NumField(); i++ {
		errorMsg := fmt.Sprintf("field %s does not exist or is an empty value in the iaas example config", tst.Field(i).Name)
		field := vst.Field(i)
		switch field.Kind() {
		case reflect.Struct:
			testPropertiesExist(vst.Field(i))
		case reflect.Bool:
			if tst.Field(i).Name != "UseUnmanagedDiskDEPRECATED" && tst.Field(i).Name != "UseInstanceProfileDEPRECATED" {
				Expect(field.Bool()).ToNot(Equal(false), errorMsg)
			}
		case reflect.String:
			Expect(field.String()).ToNot(Equal(""), errorMsg)
		case reflect.Int:
			Expect(field.Int()).ToNot(Equal(0), errorMsg)
		case reflect.Slice:
			Expect(field.Slice(0, 0)).ToNot(Equal(""), errorMsg)
		case reflect.Map:
			Expect(field.MapKeys()).ToNot(HaveLen(0), errorMsg)
		default:
			Fail(fmt.Sprintf("unexpected type: '%s' in the iaas config", field.Kind()))
		}
	}
}

func writePDFFile(contents string) string {
	tempfile, err := ioutil.TempFile("", "some*.pdf")
	Expect(err).ToNot(HaveOccurred())
	_, err = tempfile.WriteString(contents)
	Expect(err).ToNot(HaveOccurred())
	err = tempfile.Close()
	Expect(err).ToNot(HaveOccurred())

	return tempfile.Name()
}
