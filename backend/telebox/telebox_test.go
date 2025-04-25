// Test TeleBox filesystem interface
package telebox_test

import (
	"testing"

	"github.com/ky1vstar/rclone-telebox-plugin/backend/telebox"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestTeleBox:",
		NilObject:  (*telebox.Object)(nil),
		// Linkbox doesn't support leading dots for files
		SkipLeadingDot: true,
	})
}
