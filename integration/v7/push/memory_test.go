package push

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"path/filepath"
)

var _ = FDescribe("memory field in manifest", func() {
	var (
		appName  string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()
	})

	When("the manifest is in the current directory", func() {

		It("uses the manifest for memory settings", func() {
			helpers.WithHelloWorldApp(func(dir string) {
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name":   appName,
							"memory": "70M",
						},
					},
				})

				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say(`Creating app with these attributes\.\.\.`))
				Eventually(session).Should(Say(`\+\s+name:\s+%s`, appName))
				Eventually(session).Should(Say(`\s+memory:\s+70M`))
				Eventually(session).Should(Say(`Waiting for app to start\.\.\.`))
				Eventually(session).Should(Say(`requested state:\s+started`))
				Eventually(session).Should(Exit(0))
			})
		})
	})
})
