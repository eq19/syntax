// package launchertester runs end-to-end tests on Ubuntu's WSL launcher
package launchertester

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestBasicSetup runs a battery of assertions after installing with the distro launcher.
func TestBasicSetup(t *testing.T) {
	wslSetup(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	// TODO: try to inject user/password to stdin to avoid --root arg.
	out, err := launcherCommand(ctx, "install", "--root").CombinedOutput() // Installing as root to avoid Stdin
	require.NoErrorf(t, err, "Unexpected error installing: %s\n%v", out, err)
	// TODO: check with Carlos if this is necessary
	require.NotEmpty(t, out, "Failed to install the distro: No output produced.")

	testCases := map[string]func(t *testing.T){
		// TODO: Re-enable those tests once the latest wsl-setup with distro patching land.
		// "SystemdEnabled":          testSystemdEnabled,
		// "SystemdUnits":            testSystemdUnits,
		// "CorrectUpgradePolicy":    testCorrectUpgradePolicy,
		// "UpgradePolicyIdempotent": testUpgradePolicyIdempotent,
		"InteropIsEnabled": testInteropIsEnabled,
		"HelpFlag":         testHelpFlag,
	}

	for name, tc := range testCases {
		t.Run(name, tc)
	}
}
