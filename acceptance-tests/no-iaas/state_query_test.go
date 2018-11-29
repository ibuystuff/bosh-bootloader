package acceptance_test

import (
	"io/ioutil"
	"path/filepath"

	acceptance "github.com/cloudfoundry/bosh-bootloader/acceptance-tests"
	"github.com/cloudfoundry/bosh-bootloader/acceptance-tests/actors"
	"github.com/cloudfoundry/bosh-bootloader/storage"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("state query against a bbl 5.1.0 state file", func() {
	var bbl actors.BBL

	BeforeEach(func() {
		stateDir, err := ioutil.TempDir("", "")
		ioutil.WriteFile(filepath.Join(stateDir, "bbl-state.json"), []byte(BBL_STATE_5_1_0), storage.StateMode)
		Expect(err).NotTo(HaveOccurred())
		bbl = actors.NewBBL(stateDir, pathToBBL, acceptance.Config{}, "no-env", false)
	})

	It("bbl lbs", func() {
		stdout := bbl.Lbs()
		Expect(stdout).To(Equal(`CF Router LB: 35.201.97.214
CF SSH Proxy LB: 104.196.181.208
CF TCP Router LB: 35.185.98.78
CF WebSocket LB: 104.196.197.242`))
	})

	It("bbl jumpbox-address", func() {
		stdout := bbl.JumpboxAddress()
		Expect(stdout).To(Equal("35.185.60.196"))
	})

	It("bbl director-address", func() {
		stdout := bbl.DirectorAddress()
		Expect(stdout).To(Equal("https://10.0.0.6:25555"))
	})

	It("bbl director-username", func() {
		stdout := bbl.DirectorUsername()
		Expect(stdout).To(Equal("admin"))
	})

	It("bbl director-password", func() {
		stdout := bbl.DirectorPassword()
		Expect(stdout).To(Equal("some-password"))
	})

	It("bbl director-ca-cert", func() {
		stdout := bbl.DirectorCACert()
		Expect(stdout).To(Equal("-----BEGIN CERTIFICATE-----\ndirector-ca-cert\n-----END CERTIFICATE-----"))
	})

	It("bbl env-id", func() {
		stdout := bbl.EnvID()
		Expect(stdout).To(Equal("some-env-bbl5"))
	})

	It("bbl latest-error", func() {
		stdout := bbl.LatestError()
		Expect(stdout).To(Equal("latest terraform error"))
	})

	It("bbl print-env", func() {
		stdout := bbl.PrintEnv()
		Expect(stdout).To(ContainSubstring("export BOSH_CLIENT=admin"))
		Expect(stdout).To(ContainSubstring("export BOSH_CLIENT_SECRET=some-password"))
		Expect(stdout).To(ContainSubstring("export BOSH_ENVIRONMENT=https://10.0.0.6:25555"))
		Expect(stdout).To(ContainSubstring("export BOSH_CA_CERT='-----BEGIN CERTIFICATE-----\ndirector-ca-cert\n-----END CERTIFICATE-----"))
		Expect(stdout).To(ContainSubstring("export JUMPBOX_PRIVATE_KEY="))
		Expect(stdout).To(ContainSubstring("export BOSH_ALL_PROXY=ssh+socks5://jumpbox@35.185.60.196:22?private-key="))
		Expect(stdout).To(ContainSubstring("bosh_jumpbox_private.key"))
	})

	It("bbl ssh-key", func() {
		stdout := bbl.SSHKey()
		Expect(stdout).To(Equal("-----BEGIN RSA PRIVATE KEY-----\nssh-key\n-----END RSA PRIVATE KEY-----"))
	})

	It("bbl director-ssh-key", func() {
		stdout := bbl.DirectorSSHKey()
		Expect(stdout).To(Equal("-----BEGIN RSA PRIVATE KEY-----\ndirector-ssh-key\n-----END RSA PRIVATE KEY-----"))
	})
})
